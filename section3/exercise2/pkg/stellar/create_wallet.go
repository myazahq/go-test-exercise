package stellar

import (
	"fmt"

	"github.com/stellar/go/keypair"
)

type StellarPair struct {
	PublicKey string
	SecretKey string
}

func CreateStellarPair() (*StellarPair, error) {
	pair, err := keypair.Random()
	if err != nil {
		return nil, fmt.Errorf("error creating stellar pair: %v", err)
	}

	return &StellarPair{PublicKey: pair.Address(), SecretKey: pair.Seed()}, nil
}

// CreateStellarAccount adds the created wallet address to the stellar network
func CreateStellarAccount(addr string) {}
