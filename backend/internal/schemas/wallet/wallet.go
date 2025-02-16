package wallet

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	Id        string     `gorm:"primaryKey"`
	UserId    string     `gorm:"not null"`
	Name      string     `gorm:"not null;size:50" sql:"index"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (w *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	w.Id = uuid.NewString()
	return nil
}
