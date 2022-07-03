package models

type User struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Amount     int    `json:"amount"`
	Nonce      uint64 `json:"nonce"`
	PublicKey  string `json:"public_key"`
	PrimaryKey string `json:"primary_key"`
}
