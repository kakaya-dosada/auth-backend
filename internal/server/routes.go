package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/kakaya-dosada/auth-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           AuthServiceAPI
// @version         1.0
// @description     AuthServer API
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/util")
		{
			eg.GET("/health", s.healthHandler)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/login", s.AuthHandler)
		}
		admin := v1.Use(s.AdminMiddleware)
		{
			admin.POST("/user/new", s.SaveUser)
		}

		protected := v1.Use(s.AuthMiddleware)
		{
			protected.POST("/user/popa", func(c *gin.Context) {
				username := c.MustGet("username").(string)
				user, err := s.findUserByName(username)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
					return
				}

				c.JSON(http.StatusOK, user)
			})

		}

	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// @BasePath     /api/v1

// healthcheck   godoc
// @Summary      healthcheck
// @Description  healthcheck
// @Tags         util
// @Accept       json
// @Produce      json
// @Router       /util/health [get]
func (s *Server) healthHandler(c *gin.Context) {
	checks := append(make([]map[string]string, 0), s.db.Health(), s.cache.Health())
	c.JSON(http.StatusOK, checks)
}
