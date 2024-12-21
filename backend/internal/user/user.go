package database

import (
	"time"

	env "github.com/Julia-Marcal/reusable-api/config/env"
	security "github.com/Julia-Marcal/reusable-api/helpers/security"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user document in MongoDB.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UUID      string             `bson:"uuid"`
	Name      string             `bson:"name"`
	LastName  string             `bson:"last_name"`
	Age       int32              `bson:"age"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (user *User) BeforeInsert() error {
	user.UUID = uuid.NewString()

	salt := env.SetSalt()
	_, hashedPassword, err := security.DeriveKey(user.Password, salt)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return nil
}
