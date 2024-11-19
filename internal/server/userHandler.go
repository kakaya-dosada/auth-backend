package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kakaya-dosada/auth-backend/internal/models"
	"github.com/kakaya-dosada/auth-backend/internal/util"
	"github.com/kakaya-dosada/auth-backend/pkg/logger"
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
	var user models.User
	user.BeforeSave()
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong http data"})
		c.Abort()
		return
	}
	//cache
	if err := s.db.Save(user); err != nil {
		c.JSON(500, map[string]string{"error": util.DBErrors["MustBeUnique"]})
		c.Abort()
		return
	}

	if err := s.cache.Save(user); err != nil {
		logger.Warnf("Cache will be ignored cause redis unavialable %v", err)
	}
	c.JSON(201, map[string]string{"Created User": user.Username})

}
