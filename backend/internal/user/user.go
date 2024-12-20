package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	env "github.com/Julia-Marcal/reusable-api/config/env"
	security "github.com/Julia-Marcal/reusable-api/helpers/security"
)

type User struct {
	Id        string     `gorm:"primaryKey"`
	Name      string     `gorm:"not null;size:50" sql:"index"`
	LastName  string     `gorm:"not null;size:50"`
	Age       int32      `gorm:"not null"`
	Email     string     `gorm:"uniqueIndex"`
	Password  string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.NewString()

	salt := env.SetSalt()

	_, password, err := security.DeriveKey(user.Password, salt)
	if err != nil {
		return err
	}

	password_str := string(password)
	user.Password = password_str
	return nil
}
