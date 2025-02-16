package acoes

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Acoes struct {
	Id        string     `gorm:"primaryKey"`
	Name      string     `gorm:"not null"`
	Type      string     `gorm:"not null"`
	Price     float64    `gorm:"not null"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp"`
}

func (Acoes) TableName() string {
	return "acoes"
}

func (Acoes *Acoes) BeforeCreate(tx *gorm.DB) (err error) {
	Acoes.Id = uuid.NewString()

	return nil
}
