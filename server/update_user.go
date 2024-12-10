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

func (s *Server) UpdateUser(c *gin.Context, id string) {
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

	var displayName string
	var photoURL string
	var ok bool

	displayName, ok = token.Claims["name"].(string)
	if !ok {
		displayName = ""
	}
	photoURL, ok = token.Claims["picture"].(string)
	if !ok {
		photoURL = ""
	}

	user := model.User{
		Id:          id,
		DisplayName: displayName,
		PhotoURL:    photoURL,
	}

	db.Save(&user)

	c.JSON(http.StatusOK, user)
}
