package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JwtSecret = os.Getenv("JWT_SECRET")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Fprintf(gin.DefaultWriter, `{""error":"Missing Authorization header"}`+"\n")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			fmt.Fprintf(gin.DefaultWriter, `{"error":"Invalid Authorization header format"}`+"\n")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(JwtSecret), nil
		})

		if err != nil || !token.Valid {
			fmt.Fprintf(gin.DefaultWriter, `{"error":"Invalid JWT token","details":"%s"}`+"\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Optionally you can extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("user_id", claims["user_id"])
		}

		c.Next()
	}
}
