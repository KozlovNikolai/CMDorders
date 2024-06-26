package store

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/model"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, order model.Order) (uint64, error)
	GetOrderByID(ctx context.Context, order_id uint64) (*model.Order, error)
	GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]*model.Order, error)
	GetAllOrdersList(ctx context.Context, is_active int8) ([]*model.Order, error)
	UpdateOrder(ctx context.Context, order model.Order) error
	DeleteOrder(ctx context.Context, id uint64) error
}
