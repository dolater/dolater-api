package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dolater/dolater-api/db"
	api "github.com/dolater/dolater-api/generated"
	"github.com/dolater/dolater-api/model"
	"github.com/dolater/dolater-api/server/utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) ApproveFollowRequest(c *gin.Context, uid string) {
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

	approveAt := time.Now()

	fromFollowStatus := model.FollowStatus{
		FromId:     uid,
		ToId:       token.UID,
		ApprovedAt: &approveAt,
	}

	toFollowStatus := model.FollowStatus{
		FromId:     token.UID,
		ToId:       uid,
		ApprovedAt: &approveAt,
	}

	followStatuses := []model.FollowStatus{fromFollowStatus, toFollowStatus}

	if err := db.Updates(&followStatuses).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	if err := db.First(&fromFollowStatus).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	fromUser := model.User{
		Id: fromFollowStatus.FromId,
	}
	toUser := model.User{
		Id: fromFollowStatus.ToId,
	}
	if err := db.First(&fromUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}
	if err := db.First(&toUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
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
		RequestedAt: fromFollowStatus.RequestedAt,
		ApprovedAt:  fromFollowStatus.ApprovedAt,
	}

	c.JSON(http.StatusOK, response)
}
