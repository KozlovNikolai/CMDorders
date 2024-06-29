package store

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/models"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) (uint64, error)
	GetOrderByID(ctx context.Context, order_id uint64) (*models.Order, error)
	GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]models.Order, error)
	GetAllOrdersList(ctx context.Context, is_active int8) ([]models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) error
	DeleteOrder(ctx context.Context, id uint64) error
	AddServicesToOrder(ctx context.Context, order_id uint64, patient_id uint64, services []models.Service) error
}
