package server

import (
	"log"
	"net/http"

	"github.com/dolater/dolater-api/db"
	api "github.com/dolater/dolater-api/generated"
	"github.com/dolater/dolater-api/model"
	"github.com/dolater/dolater-api/server/utility"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetFollowRequest(c *gin.Context, uid string) {
	token := utility.GetToken(c)
	if token == nil {
		message := "Unauthorized"
		c.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{
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
		FromId: uid,
		ToId:   token.UID,
	}

	if err := db.First(&followStatus).Error; err != nil {
		message := err.Error()
		c.JSON(http.StatusNotFound, api.Error{
			Message: &message,
		})
		return
	}

	fromUser := model.User{
		Id: followStatus.FromId,
	}
	toUser := model.User{
		Id: followStatus.ToId,
	}
	if err := db.First(&fromUser).Error; err != nil {
		message := err.Error()
		c.JSON(http.StatusNotFound, api.Error{
			Message: &message,
		})
		return
	}
	if err := db.First(&toUser).Error; err != nil {
		message := err.Error()
		c.JSON(http.StatusNotFound, api.Error{
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
		RequestedAt: followStatus.RequestedAt,
		ApprovedAt:  followStatus.ApprovedAt,
	}
	c.JSON(http.StatusOK, response)
}
