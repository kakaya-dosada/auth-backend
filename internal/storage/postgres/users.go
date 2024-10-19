package postgres

import (
	"github.com/google/uuid"
	"github.com/kakaya-dosada/auth-backend/internal/models"
)

// // Save user details
func (service *service) Save(RoleID, username, password, email string) (*models.User, error) {
	user := models.User{
		ID:       uuid.NewString(),
		RoleID:   RoleID,
		Username: username,
		// Role:      service.GetRoleByID(RoleID),
		Email:    email,
		Password: password,
	}
	user.BeforeSave()
	// user.Role, err := service.GetRoleByID(RoleID)
	// if err != nil {
	// return nil, err
	// }
	_, err := service.db.Exec("INSERT INTO users (id,username, email, password, role_id) values ($1,$2,$3,$4,$5)", user.ID, user.Username, user.Email, user.Password, user.RoleID)
	if err != nil {
		return nil, err
	}
	return &user, nil

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
