package models

type CreateVirtualCardRequest struct {
	Currency    string `json:"currency"`
	Amount      int32  `json:"amount"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"` //expected format YYYY/MM/DD
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Title       string `json:"title"`
	Gender      string `json:"gender"`
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
