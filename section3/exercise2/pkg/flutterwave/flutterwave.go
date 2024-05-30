package flutterwave

import (
	"net/http"
	"time"
)

const (
	SuccessStatus        = "success"
	flutterBaseURL       = "https://api.flutterwave.com/v3/"
	createVirtualCardURL = "virtual-cards"
)

type Flutterwave struct {
	secretKey string
	client    *http.Client
}

// NewFlutterwave sets up the flutter wave client with the secret key
func NewFlutterwave(secret string) *Flutterwave {
	return &Flutterwave{
		secretKey: secret,
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
	}
}
