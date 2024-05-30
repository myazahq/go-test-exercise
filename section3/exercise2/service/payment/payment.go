package payment

import (
	"github.com/myazahq/go-test-exercise/section3/exercise2/models"
	"github.com/myazahq/go-test-exercise/section3/exercise2/pkg/flutterwave"
)

type paymentService struct {
	flutterwave *flutterwave.Flutterwave
}

type PaymentService interface {
	CreateVirtualCard(*models.CreateVirtualCardRequest) (*models.CreateVirtualCardResponse, error)
}

func NewPaymentService(fw *flutterwave.Flutterwave) *paymentService {
	return &paymentService{
		flutterwave: fw,
	}
}

func (ps *paymentService) CreateVirtualCard(req *models.CreateVirtualCardRequest) (*models.CreateVirtualCardResponse, error) {

	virtualCardReq := flutterwave.CreateVirtualCardReq{
		Currency:    req.Currency,
		Amount:      req.Amount,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Email:       req.Email,
		Phone:       req.Phone,
		Title:       req.Title,
		Gender:      req.Gender,
	}

	res, err := ps.flutterwave.CreateVirtualCard(&virtualCardReq)
	if err != nil {
		return nil, err
	}

	return &models.CreateVirtualCardResponse{
		ID:         res.Data.ID,
		AmountID:   res.Data.AmountID,
		Amount:     res.Data.Amount,
		Currency:   res.Data.Currency,
		CardPan:    res.Data.CardPan,
		MaskedPan:  res.Data.MaskedPan,
		City:       res.Data.City,
		State:      res.Data.State,
		Address1:   res.Data.Address1,
		ZipCode:    res.Data.ZipCode,
		CVV:        res.Data.CVV,
		Expiration: res.Data.Expiration,
		CardType:   res.Data.CardType,
		NameOnCard: res.Data.NameOnCard,
		CreatedAt:  res.Data.CreatedAt,
		IsActive:   res.Data.IsActive,
	}, nil
}
