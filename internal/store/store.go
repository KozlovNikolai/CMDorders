package store

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/models"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order models.Order) (int, error)
	GetOrderByID(ctx context.Context, order_id int) (*models.Order, error)
	GetOrdersByPatientID(ctx context.Context, patient_id int, is_active bool) ([]models.Order, error)
	GetAllOrdersList(ctx context.Context, is_active bool) ([]models.Order, error)
	UpdateOrder(ctx context.Context, order models.Order) error
	DeleteOrder(ctx context.Context, id int) error
	AddServicesToOrder(ctx context.Context, order_id int, patient_id int, services []int) error
}
