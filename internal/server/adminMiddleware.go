package server

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (server *Server) AdminMiddleware(c *gin.Context) {
	//check xapi key
	authHeader := c.GetHeader("x-api-key")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Missing x-api-key header"})
		return
	}
	if authHeader != os.Getenv("API_KEY") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	c.Next()
}
