package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kakaya-dosada/auth-backend/internal/models"
	"github.com/kakaya-dosada/auth-backend/pkg/logger"
)

func (service *service) Save(user models.User) error {
	userJson, _ := json.Marshal(user)

	err := service.db.Set(context.Background(), user.Username, userJson, 0)
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
func (service *service) GetUserByID(id string) (*models.User, error) {
	key := id
	data, err := service.db.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("error getting user from Redis: %w", err)
	}

	var user models.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("error unmarshalling user data: %w", err)
	}

	return &user, nil
}
func (service *service) GetUserByName(name string) (*models.User, error) {
	key := name
	data, err := service.db.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("error getting user from Redis: %w", err)
	}

	var user models.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("error unmarshalling user data: %w", err)
	}

	return &user, nil
}
