# Web Service
The web service creates a virtual card using flutterwave and blockchain wallet using stellar blockchain

### Setup Locally
- Use the .env_sample to set up your .env file
- Run `go mod tidy` on your terminal to download service dependencies
- Run `go run main.go` to start up app.

## API ENDPOINTS DOCUMENTATION

### Create Virtual Card (POST)
creates a virtual card using flutterwave

Endpoint
```
{HTTP_SERVER_ADDRESS}/api/v1/create-virtual-card
```
Payload
```
{
	"currency": "USD",
	"amount": 100,
	"first_name": "Tobi",
	"last_name": "Olusola",
	"date_of_birth": "2024/05/28",
	"email": "email@email.com",
	"phone": "09000000000",
	"title": "Mrs",
	"gender": "Unknown"
}
```
Response
```
{
    "id": "199a344f-1dbe-4b00-ba4d-beb014345fae",
    "account_id": 2061620,
    "amount": "5.00",
    "currency": "USD",
    "card_pan": "5319938155020288",
    "masked_pan": "531993*******0288",
    "city": "San Francisco",
    "state": "CA",
    "address_1": "333 Fremont Street",
    "zip_code": "94105",
    "cvv": "905",
    "expiration": "2025-09",
    "card_type": "mastercard",
    "name_on_card": "Tobi Olusola",
    "created_at": "2022-09-21T16:54:53.3851427+00:00",
    "is_active": true,
}
```

### Create Wallet (Get)
create wallet using stellar blockchain

Endpoint
```
{HTTP_SERVER_ADDRESS}/api/v1/wallet
```
Response
```
{
    "address": "GBJZO2QHGF5PM7S75H3LUODWUQ6NLKE5KTQNF26XVKNODSY52DTMPLHQ",
    "password": "SCJUKDQYOBCT2E6FJ2PHRD6B44DGCWOVHTCAL7NN43EYUPQJVCV2VVND"
}
```
