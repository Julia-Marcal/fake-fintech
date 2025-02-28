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

func FindUserByEmail(email string) (*database.User, error) {
	db := repository.NewPostgres()
	user := &database.User{}
	result := db.First(user, "email = ?", email)
	return user, result.Error
}

func FindUserById(id string) (*database.User, error) {
	db := repository.NewPostgres()
	user := &database.User{}
	result := db.First(user, "id = ?", id)
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

func TotalAmountByUser(userId string) (float64, error) {
	db := repository.NewPostgres()
	var totalAmount float64

	result := db.Table("acoes").
		Select("COALESCE(SUM(acoes.price * acoes.quantity), 0) as total").
		Joins("JOIN wallet_acoes ON wallet_acoes.acoes_id = acoes.id").
		Joins("JOIN wallets ON wallets.id = wallet_acoes.wallet_id AND wallets.user_id = ?", userId).
		Scan(&totalAmount)

	if result.Error != nil {
		return 0, result.Error
	}

	return totalAmount, nil
}

func Update(user_info *database.User) error {
	db := repository.NewPostgres()
	result := db.Save(user_info)
	return result.Error
}
