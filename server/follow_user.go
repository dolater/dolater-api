package server

import (
	"errors"
	"log"
	"net/http"

	"firebase.google.com/go/v4/messaging"
	"github.com/dolater/dolater-api/db"
	api "github.com/dolater/dolater-api/generated"
	"github.com/dolater/dolater-api/model"
	"github.com/dolater/dolater-api/server/utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Server) FollowUser(c *gin.Context, uid string) {
	token := utility.GetToken(c)
	if token == nil {
		message := "Unauthorized"
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{
			Message: &message,
		})
		return
	}

	if uid == token.UID {
		message := "Cannot follow yourself"
		c.JSON(http.StatusBadRequest, api.Error{
			Message: &message,
		})
		return
	}

	db, err := db.GormDB("public")
	if err != nil {
		message := err.Error()
		c.JSON(http.StatusInternalServerError, api.Error{
			Message: &message,
		})
		return
	}
	defer func() {
		sqldb, err := db.DB()
		if err != nil {
			log.Println("Failed to close database connection")
		}
		sqldb.Close()
	}()

	followStatus := model.FollowStatus{
		FromId: token.UID,
		ToId:   uid,
	}

	if err := db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&followStatus).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	fromUser := model.User{
		Id: followStatus.FromId,
	}

	toUser := model.User{
		Id: followStatus.ToId,
	}

	if err := db.First(&fromUser).Error; err != nil {
		message := err.Error()
		c.JSON(http.StatusInternalServerError, api.Error{
			Message: &message,
		})
		return
	}

	if err := db.First(&toUser).Error; err != nil {
		message := err.Error()
		c.JSON(http.StatusInternalServerError, api.Error{
			Message: &message,
		})
		return
	}

	response := api.FollowStatus{
		From: api.User{
			Id: fromUser.Id,
			DisplayName: func() string {
				if fromUser.DisplayName == nil {
					return ""
				}
				return *fromUser.DisplayName
			}(),
			PhotoURL: func() string {
				if fromUser.PhotoURL == nil {
					return ""
				}
				return *fromUser.PhotoURL
			}(),
		},
		To: api.User{
			Id: toUser.Id,
			DisplayName: func() string {
				if toUser.DisplayName == nil {
					return ""
				}
				return *toUser.DisplayName
			}(),
			PhotoURL: func() string {
				if toUser.PhotoURL == nil {
					return ""
				}
				return *toUser.PhotoURL
			}(),
		},
		Timestamp: followStatus.CreatedAt,
	}

	fcmToken := model.FCMToken{
		UserId: followStatus.ToId,
	}

	if err := db.First(&fcmToken).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	// 通知メッセージを作成
	messages := []*messaging.Message{}
	messages = append(messages, &messaging.Message{
		Notification: &messaging.Notification{
			Title: "You have a new follower!",
			Body:  response.From.DisplayName + " is now following you.",
		},
		Token: fcmToken.RegistrationToken,
	})

	if len(messages) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	// 通知を送信
	client, err := s.FirebaseApp.Messaging(c)
	if err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{
			Message: &message,
		})
		return
	}

	_, err = client.SendEach(c, messages)
	if err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{
			Message: &message,
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
