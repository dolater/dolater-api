package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/dolater/dolater-api/db"
	api "github.com/dolater/dolater-api/generated"
	"github.com/dolater/dolater-api/model"
	"github.com/dolater/dolater-api/server/utility"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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
		Id:          token.UID,
		DisplayName: &displayName,
		PhotoURL:    &photoURL,
	}

	if err := db.Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) || pgErr.Code != "23505" {
			message := err.Error()
			c.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}
	}

	c.JSON(http.StatusCreated, user)
}
