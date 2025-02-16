package WalletAcoes

type WalletAcoes struct {
	WalletId string `gorm:"primaryKey"`
	AcoesId  string `gorm:"primaryKey"`
	Quantity int    `gorm:"not null"`
}

func (WalletAcoes) TableName() string {
	return "wallet_acoes"
}
