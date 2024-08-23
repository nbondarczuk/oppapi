package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"oppapi/internal/cache"
	"oppapi/internal/config"
	"oppapi/internal/handler"
	"oppapi/internal/handler/payment"
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
	cache.InitWithRedis(config.CacheRedisAddress(), config.CacheRedisPassword(), config.CacheRedisDB())
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

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	s.router.GET("/health", handler.HealthHandler)
	s.router.POST("/payment", payment.CreateHandler)
	s.router.GET("/payment/:id", payment.ReadOneHandler)
}
