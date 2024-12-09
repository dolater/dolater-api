package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/kantacky/p2hacks2024-test-api/generated"
)

func (m Middleware) RequireAppCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Request.Header[http.CanonicalHeaderKey("X-Firebase-AppCheck")]
		if !ok {
			message := "X-Firebase-AppCheck header is required"
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{
				Message: &message,
			})
			return
		}

		client, err := m.FirebaseApp.AppCheck(c)
		if err != nil {
			log.Printf("Error initializing app: %v\n", err)
		}

		_, err = client.VerifyToken(token[0])
		if err != nil {
			message := "X-Firebase-AppCheck header value is invalid"
			c.AbortWithStatusJSON(http.StatusUnauthorized, api.Error{
				Message: &message,
			})
			return
		}

		c.Next()
	}
}
