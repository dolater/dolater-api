package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kantacky/p2hacks2024-test-api/db"
	api "github.com/kantacky/p2hacks2024-test-api/generated"
	"github.com/kantacky/p2hacks2024-test-api/model"
	"github.com/kantacky/p2hacks2024-test-api/server/utility"
)

func (s *Server) UpsertFCMToken(c *gin.Context) {
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

	var requestBody api.UpsertFCMTokenJSONRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		message := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, api.Error{
			Message: &message,
		})
		return
	}

	fcmToken := model.FCMToken{
		RegistrationToken: requestBody.Token,
		UserId:            token.UID,
		Timestamp:         requestBody.Timestamp,
	}

	if err := db.Save(&fcmToken).Error; err != nil {
		if err.Error() != "record not found" {
			message := err.Error()
			c.JSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	c.Status(http.StatusNoContent)
}
