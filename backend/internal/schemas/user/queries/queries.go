package queries_users

import (
	repository "github.com/Julia-Marcal/fake-fintech/internal/database"
	database "github.com/Julia-Marcal/fake-fintech/internal/schemas/user"
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

func FindUsers(limit int) ([]database.User, error) {
	db := repository.NewPostgres()
	var users []database.User

	result := db.Limit(limit).Find(&users) 

	return users, result.Error
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
