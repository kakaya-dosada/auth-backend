package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/kakaya-dosada/auth-backend/internal/storage/postgres"
	"github.com/kakaya-dosada/auth-backend/internal/storage/redis"
	"github.com/kakaya-dosada/auth-backend/pkg/logger"
)

type Server struct {
	port  int
	cache redis.Service
	db    postgres.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port:  port,
		cache: redis.New(),
		db:    postgres.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if os.Getenv("restoreCacheOnRestart") == "true" {
		logger.Infof("Restoring Cache from DB ...")
		go func() {
			users, err := NewServer.db.AllUsers()
			if err != nil {
				logger.Warnf(err.Error())
			}
			err = NewServer.cache.RestoreUsers(users)
			if err != nil {
				logger.Warnf(err.Error())
			}
		}()
		logger.Infof("Restored cache!")

	}

	return server
}
