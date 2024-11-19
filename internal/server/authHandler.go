package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kakaya-dosada/auth-backend/internal/models"
)

func (server *Server) AuthHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	u, err := server.findUserByName(user.Username)
	if err != nil || u == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "User with name " + user.Username + " not found"})
		return
	}
	if u.Password != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (server *Server) findUserByID(id string) (*models.User, error) {

	user, err := server.cache.GetUserByID(id)
	if err != nil {
		user, err = server.db.GetUserByID(id)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (server *Server) findUserByName(name string) (*models.User, error) {

	user, err := server.cache.GetUserByName(name)
	if err != nil {
		user, err = server.db.GetUserByID(name)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
