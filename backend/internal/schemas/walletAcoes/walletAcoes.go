package WalletAcoes

type WalletAcoes struct {
	WalletId string `gorm:"primaryKey"`
	AcoesId  string `gorm:"primaryKey"`
}

func (WalletAcoes) TableName() string {
	return "wallet_acoes"
}
