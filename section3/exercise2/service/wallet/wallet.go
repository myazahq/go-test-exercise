package wallet

import (
	"fmt"

	"github.com/myazahq/go-test-exercise/section3/exercise2/models"
	"github.com/myazahq/go-test-exercise/section3/exercise2/pkg/stellar"
)

type walletService struct{}

type WalletService interface {
	CreateWallet() (*models.WalletResponse, error)
}

func NewWalletService() *walletService {
	return &walletService{}
}

func (ws *walletService) CreateWallet() (*models.WalletResponse, error) {
	stellarPair, err := stellar.CreateStellarPair()
	if err != nil {
		return nil, fmt.Errorf("wallet service error: %v", err)
	}

	return &models.WalletResponse{
		PublicKey: stellarPair.PublicKey,
		SecretKey: stellarPair.SecretKey,
	}, nil
}
