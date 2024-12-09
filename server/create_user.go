package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kantacky/p2hacks2024-test-api/db"
	api "github.com/kantacky/p2hacks2024-test-api/generated"
	"github.com/kantacky/p2hacks2024-test-api/server/utility"
)

func (s *Server) CreateUser(c *gin.Context) {
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

	user := api.User{}

	c.JSON(http.StatusOK, user)
}
