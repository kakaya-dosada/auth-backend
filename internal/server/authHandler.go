package server

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) AuthHandler(c *gin.Context) {
	// var user models.User
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	// 	return
	// }

	// u, ok := findUserByID(user.Username)
	// if !ok || u.Password != user.Password {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	// 	return
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"sub": u.Username,
	// 	"exp": time.Now().Add(time.Hour * 24).Unix(),
	// })

	// tokenString, err := token.SignedString([]byte(secretKey))
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (server *Server) findUserByID()
