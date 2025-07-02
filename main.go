package main

import (
	"awesome-bank/configs"
	"awesome-bank/handlers"
	"awesome-bank/middleware"
	"awesome-bank/models"
	"awesome-bank/repositories"
	"awesome-bank/services"
	"awesome-bank/services/utils"

	"github.com/gorilla/mux"

	"log"
	"net/http"
)

func main() {
	utils.InitLogger()
	utils.Logger.Info("Starting server...")

	cfg := configs.LoadConfig()

	if err := configs.Connect(cfg); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	err := configs.DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Card{}, &models.Transaction{}, &models.Credit{}, &models.PaymentSchedule{})
	if err != nil {
		return
	}

	r := mux.NewRouter()

	userRepo := repositories.NewUserRepository(configs.DB)
	cardRepo := repositories.NewCardRepository(configs.DB)
	creditRepo := repositories.NewCreditRepository(configs.DB)
	accountRepo := repositories.NewAccountRepository(configs.DB)
	transactionRepo := repositories.NewTransactionRepository(configs.DB)
	paymentScheduleRepo := repositories.NewPaymentScheduleRepository(configs.DB)

	authService := services.NewAuthService(userRepo)
	accountService := services.NewAccountService(accountRepo, transactionRepo)
	analyticsService := services.NewAnalyticsService(creditRepo, accountRepo, transactionRepo, paymentScheduleRepo)
	cardService := services.NewCardService(cardRepo)
	creditService := services.NewCreditService(paymentScheduleRepo, accountRepo, creditRepo)
	transactionService := services.NewTransactionService(configs.DB, accountRepo, userRepo)

	userHandler := handlers.NewUserHandler(authService, utils.Logger)
	cardHandler := handlers.NewCardHandler(cardService, utils.Logger)
	accountHandler := handlers.NewAccountHandler(accountService, utils.Logger)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService, utils.Logger)
	cbrHandler := handlers.NewCbrHandler(creditService, utils.Logger)
	creditHandler := handlers.NewCreditHandler(creditService, utils.Logger)
	transactionHandler := handlers.NewTransactionHandler(transactionService, utils.Logger)

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/cbr/keyrate", cbrHandler.GetKeyRate).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/accounts", accountHandler.GetAccounts).Methods("GET")
	protected.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	protected.HandleFunc("/accounts/{id}", accountHandler.UpdateAccount).Methods("PUT")
	protected.HandleFunc("/cards", cardHandler.GetCards).Methods("GET")
	protected.HandleFunc("/cards", cardHandler.IssueCard).Methods("POST")
	protected.HandleFunc("/transfer", transactionHandler.TransferFunds).Methods("POST")
	protected.HandleFunc("/analytics", analyticsHandler.GetAnalytics).Methods("GET")
	protected.HandleFunc("/credits", creditHandler.ApplyForCredit).Methods("POST")

	finalRouter := middleware.LoggingMiddleware(r)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", finalRouter))
}
