package models

type WalletResponse struct {
	PublicKey string `json:"address"`
	SecretKey string `json:"password"`
}
