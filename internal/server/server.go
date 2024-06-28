package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KozlovNikolai/CMDorders/internal/config"
	"github.com/KozlovNikolai/CMDorders/internal/handlers"
	"github.com/KozlovNikolai/CMDorders/internal/store"
	"github.com/KozlovNikolai/CMDorders/internal/store/inmemory"
	"github.com/KozlovNikolai/CMDorders/internal/store/mongostore"
	"github.com/KozlovNikolai/CMDorders/internal/store/pgstore"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	logger *zap.Logger
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	// Инициализация логгера Zap
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	var repo store.IOrderRepository

	// Выбор репозитория
	switch cfg.RepoType {
	case "memory":
		repo = inmemory.NewInMemoryOrderRepository()
	case "postgres":
		pool, err := pgxpool.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
		if err != nil {
			logger.Fatal("Unable to connect to database", zap.Error(err))
		}
		repo = pgstore.NewPostgresOrderRepository(pool)
	case "mongo":
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			logger.Fatal("Unable to connect to MongoDB", zap.Error(err))
		}
		repo = mongostore.NewMongoOrderRepository(client, "mydatabase", "employers")
	default:
		logger.Fatal("Invalid repository type")
	}

	// Создание сервера
	server := &Server{
		router: gin.Default(),
		logger: logger,
		config: cfg,
	}

	// Инициализация обработчиков
	orderHandler := handlers.NewOrderHandler(logger, repo)

	// CRUD маршруты для Employers
	server.router.POST("/orders", orderHandler.CreateOrder)
	server.router.GET("/orders/:id", orderHandler.GetOrderByID)
	server.router.GET("/orders/list", orderHandler.GetAllOrdersList)
	server.router.GET("/orders/bypatient/:patient_id/:is_active", orderHandler.GetOrdersByPatientID)
	server.router.PUT("/orders/:id", orderHandler.UpdateOrder)
	server.router.DELETE("/orders/:id", orderHandler.DeleteOrder)

	return server
}

func (s *Server) Run() {
	defer s.logger.Sync() // flushes buffer, if any

	// Настройка сервера с таймаутами
	server := &http.Server{
		Addr:         s.config.Address,
		Handler:      s.router,
		ReadTimeout:  s.config.Timeout,
		WriteTimeout: s.config.Timeout,
		IdleTimeout:  s.config.IdleTimout,
	}

	// Запуск сервера
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Fatal(fmt.Sprintf("Could not listen on %s", s.config.Address), zap.Error(err))
	}
}
