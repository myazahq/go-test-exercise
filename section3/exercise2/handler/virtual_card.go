package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myazahq/go-test-exercise/section3/exercise2/models"
	"github.com/myazahq/go-test-exercise/section3/exercise2/service/payment"
)

type VirtualCardHandler interface {
	CreateVirtualCard(ctx *gin.Context)
}

type virtualCardHandler struct {
	paymentService payment.PaymentService
}

func NewVirtualCardHandler(ps payment.PaymentService) *virtualCardHandler{
	return &virtualCardHandler{
		paymentService: ps,
	}
}

func (vh *virtualCardHandler) CreateVirtualCard(ctx *gin.Context) {

	var virtualCardReq models.CreateVirtualCardRequest

	if err := ctx.ShouldBindJSON(&virtualCardReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	virtualCardResponse, err := vh.paymentService.CreateVirtualCard(&virtualCardReq)
	if err != nil {
		log.Println(err)
		
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":http.StatusText(http.StatusInternalServerError)})
        return
	}

	ctx.JSON(http.StatusOK, virtualCardResponse)
}
