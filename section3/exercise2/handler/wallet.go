package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myazahq/go-test-exercise/section3/exercise2/service/wallet"
)

type WalletHandler interface {
	CreateWallet(ctx *gin.Context)
}

type walletHandler struct {
	walletService wallet.WalletService
}

func NewWalletHandler(ws wallet.WalletService) *walletHandler {
	return &walletHandler{
		walletService: ws,
	}
}

func (vh *walletHandler) CreateWallet(ctx *gin.Context) {
	walletResponse, err := vh.walletService.CreateWallet()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	ctx.JSON(http.StatusOK, walletResponse)
}
