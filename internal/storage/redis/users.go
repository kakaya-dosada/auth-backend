package redis

import (
	"context"
	"encoding/json"

	"github.com/kakaya-dosada/auth-backend/internal/models"
	"github.com/kakaya-dosada/auth-backend/pkg/logger"
)

func (service *service) Save(user models.User) error {
	userJson, _ := json.Marshal(user)

	err := service.db.Set(context.Background(), user.ID, userJson, 0)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (service *service) RestoreUsers(users []models.User) error {

	for _, user := range users {
		err := service.Save(user)
		if err != nil {
			logger.Warnf("Cannot restore some data")
		}
	}
	return nil
}
