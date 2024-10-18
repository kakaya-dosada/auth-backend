package models

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	RoleID    uint   `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Password  string `gorm:"size:255;not null" json:"-"`
	Role      Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
