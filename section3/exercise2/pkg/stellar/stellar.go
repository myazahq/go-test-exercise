package stellar

import (
	"net/http"
	"time"
)

const (
	SuccessStatus        = "success"
	flutterBaseURL       = "https://api.flutterwave.com/v3/"
	createVirtualCardURL = "virtual-cards"
)

type stellar struct {
	client *http.Client
}

// NewFlutterwave sets up the flutter wave client with the secret key
func NewStellar(secret string) *stellar {
	return &stellar{
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   30 * time.Second,
		},
	}
}
