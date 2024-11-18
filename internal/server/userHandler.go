package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kakaya-dosada/auth-backend/internal/models"
)

// @BasePath     /api/v1
// @Summary Create new user
// @Description This will create a new user in the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserDto true "Create user"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router       /user/new [POST]
func (s *Server) SaveUser(c *gin.Context) {
	// user, err := s.db.Save(c.Query("role"), c.Query("name"), c.Query("password"), c.Query("email"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong http data"})
		c.Abort()
		return
	}
	//cache
	if err := s.db.Save(user); err != nil {
		c.JSON(500, map[string]error{"Error": err}) //todo UTIL erroring
		c.Abort()
		return
	}
	c.JSON(201, map[string]string{"Created User": user.Username})

}
