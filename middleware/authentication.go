package middleware

import (
	"asssignment2/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			errorMessage := fmt.Sprintf("%s", err)
			response := helpers.APIResponse("Unauthorized", http.StatusUnauthorized, "Error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
