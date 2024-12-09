package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	api "github.com/kantacky/p2hacks2024-test-api/generated"
)

func (m Middleware) GetFirebaseAuthIDToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(bearerToken, "Bearer ", "", 1)

		client, err := m.FirebaseApp.Auth(c)
		if err != nil {
			message := err.Error()
			c.AbortWithStatusJSON(http.StatusInternalServerError, api.Error{
				Message: &message,
			})
			return
		}

		token, err := client.VerifyIDToken(c, idToken)
		if err != nil {
			message := err.Error()
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{
				Message: &message,
			})
			return
		}

		c.Set("X-Firebase-Authentication-ID-Token", token)

		c.Next()
	}
}
