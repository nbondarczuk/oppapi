package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"oppapi/internal/config"
	"oppapi/internal/handler"
	"oppapi/internal/handler/payment"
	"oppapi/internal/handler/refund"
	"oppapi/internal/handler/transaction"
	"oppapi/internal/logging"
	"oppapi/internal/middleware"
	"oppapi/internal/repository"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New(version string) (*Server, error) {
	repository.InitWithMongo(config.RepositoryDBName(), config.RepositoryURL())
	gin.SetMode(gin.ReleaseMode)

	// Make new server with gin router.
	s := &Server{
		router: gin.New(),
	}

	// Apply required middleware. The order matters as request log should
	// go before response log.
	s.router.Use(middleware.ResponseLogger())
	s.router.Use(middleware.RequestLogger())

	// Register defined handlers.
	s.RegisterHandlers()

	return s, nil
}

// Run the gin server on routes.
func (s *Server) Run() error {
	port := config.ServerHTTPPort()
	logging.Logger.Info("Starting HTTP server", slog.String("port", port))
	return s.router.Run(":" + port)
}

// RegisterHandlers links handlers to api points.
func (s *Server) RegisterHandlers() {
	// system api point for health check
	s.router.GET("/health", handler.HealthHandler)

	// api points for payments & refunds
	s.router.POST("/payment", payment.CreatePaymentHandler)
	s.router.GET("/payment/:id", payment.ReadOnePaymentHandler)
	s.router.POST("/refund/:id", refund.CreateRefundHandler)
	s.router.GET("/refund/:id", refund.ReadOneRefundHandler)

	// bank mock for payment processing
	s.router.POST("/bankmock/transaction", transaction.CreateHandler)
}
