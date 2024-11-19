package postgres

import (
	"github.com/kakaya-dosada/auth-backend/internal/models"
	"github.com/kakaya-dosada/auth-backend/internal/storage/postgres/queries"
)

// // Save user details
func (service *service) Save(user models.User) error {

	_, err := service.db.Exec(queries.INSERT_USER, user.ID, user.Username, user.Email, user.Password, user.RoleID)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) AllUsers() ([]models.User, error) {
	rows, err := service.db.Query(queries.SELECT_ALL_USERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.RoleID, &user.Username, &user.Email, &user.Password); err != nil {
			//skip
			_ = "a"
		}
		users = append(users, user)
	}
	return users, nil
}

func (service *service) GetUserByID(id string) (*models.User, error) {
	return nil, nil
}

// 	err := database.Db.Create(&user).Error
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	return user, nil
// }

// // Generate encrypted password
// func (user *User) BeforeSave(*gorm.DB) error {
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(passwordHash)
// 	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
// 	return nil
// }

// // Get all users
// func GetUsers(User *[]User) (err error) {
// 	err = database.Db.Find(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Get user by username
// func GetUserByUsername(username string) (User, error) {
// 	var user User
// 	err := database.Db.Where("username=?", username).Find(&user).Error
// 	if err != nil {
// 		return User{}, err
// 	}
// 	return user, nil
// }

// // Validate user password
// func (user *User) ValidateUserPassword(password string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
// }

// // Get user by id
// func GetUserById(id uint) (User, error) {
// 	var user User
// 	err := database.Db.Where("id=?", id).Find(&user).Error
// 	if err != nil {
// 		return User{}, err
// 	}
// 	return user, nil
// }

// // Get user by id
// func GetUser(User *User, id int) (err error) {
// 	err = database.Db.Where("id = ?", id).First(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Update user
// func UpdateUser(User *User) (err error) {
// 	err = database.Db.Omit("password").Updates(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
