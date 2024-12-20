package queries

import (
	repository "github.com/Julia-Marcal/reusable-api/internal/database"
	database "github.com/Julia-Marcal/reusable-api/internal/user"
	_ "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(user_info *database.User) error {
	db := repository.NewPostgres()
	result := db.Create(user_info)
	return result.Error
}

func FindUser(email string) (*database.User, error) {
	db := repository.NewPostgres()
	user := &database.User{}
	result := db.First(user, "email = ?", email)
	return user, result.Error
}

func FindUsers() (int64, error) {
	db := repository.NewPostgres()
	var users []database.User
	result := db.Limit(10).Find(&users)
	return result.RowsAffected, result.Error
}

func DeleteOne(id string) *gorm.DB {
	db := repository.NewPostgres()
	user := &database.User{}
	result := db.Delete(user, "id = ?", id)
	return result
}

func CheckPassword(email string) (string, error) {
	db := repository.NewPostgres()
	user := &database.User{}
	result := db.First(user, "email = ?", email)
	return user.Password, result.Error
}
