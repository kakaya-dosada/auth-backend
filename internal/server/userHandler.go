package server

import "github.com/gin-gonic/gin"

// @BasePath     /api/v1

// SaveUser   godoc
// @Summary      SaveUser
// @Description  Creates new user
// @Tags         util
// @Accept       json
// @Produce      json
// @Param 		 role query string false "role"
// @Param        name    query     string  false  "name"
// @Param        email    query     string  false  "email"  Format(email)
// @Param        password    query     string  false  "password"
// @Router       /user/new [POST]
func (s *Server) SaveUser(c *gin.Context) {
	user, err := s.db.Save(c.Query("role"), c.Query("name"), c.Query("password"), c.Query("email"))
	if err != nil {
		c.JSON(500, err.Error()) //todo UTIL erroring
	}
	c.JSON(200, user)
}
