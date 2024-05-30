package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/myazahq/go-test-exercise/section3/exercise2/handler"
	"github.com/myazahq/go-test-exercise/section3/exercise2/pkg/flutterwave"
	"github.com/myazahq/go-test-exercise/section3/exercise2/service/payment"
	"github.com/myazahq/go-test-exercise/section3/exercise2/service/wallet"
)

func main() {
	// env := os.Getenv("ENV")
	if os.Getenv("ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	fluttwaveSecretKey := os.Getenv("FLUTTER_WAVE_SECRET")

	router := gin.New()
	v1 := router.Group("/api/v1")

	virtualCardHandler := handler.NewVirtualCardHandler(payment.NewPaymentService(flutterwave.NewFlutterwave(fluttwaveSecretKey)))

	v1.POST("/create-virtual-card", virtualCardHandler.CreateVirtualCard)

	walletHandler := handler.NewWalletHandler(wallet.NewWalletService())

	v1.GET("/wallet", walletHandler.CreateWallet)

	httpServer := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: router,
	}

	log.Println("Starting server: ", os.Getenv("SERVER_ADDRESS"))

	log.Fatal(httpServer.ListenAndServe())
}
