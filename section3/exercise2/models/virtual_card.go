package models

type CreateVirtualCardRequest struct {
	Currency    string `json:"currency" binding:"required"`
	Amount      int32  `json:"amount" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"` //expected format YYYY/MM/DD
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}

type CreateVirtualCardResponse struct {
	ID         string `json:"id"`
	AmountID   int    `json:"amount_id"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CardPan    string `json:"card_pan"`
	MaskedPan  string `json:"masked_pan"`
	City       string `json:"city"`
	State      string `json:"state"`
	Address1   string `json:"address_1"`
	ZipCode    string `json:"zip_code"`
	CVV        string `json:"cvv"`
	Expiration string `json:"expiration"`
	CardType   string `json:"card_type"`
	NameOnCard string `json:"name_on_card"`
	CreatedAt  string `json:"created_at"`
	IsActive   bool   `json:"is_active"`
}
