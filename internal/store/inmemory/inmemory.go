package inmemory

import (
	"context"
	"errors"
	"time"

	"sync"

	"github.com/KozlovNikolai/CMDorders/internal/model"
)

type InMemoryOrderRepository struct {
	orders map[uint64]model.Order
	nextID uint64
	mutex  sync.Mutex
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[uint64]model.Order),
		nextID: 1,
	}
}

func (repo *InMemoryOrderRepository) CreateOrder(ctx context.Context, order model.Order) (uint64, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order.ID = repo.nextID
	order.CreatedAt = time.Now().UTC()
	repo.nextID++
	repo.orders[order.ID] = order
	return order.ID, nil
}

func (repo *InMemoryOrderRepository) GetOrderByID(ctx context.Context, order_id uint64) (*model.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order, exists := repo.orders[order_id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

func (repo *InMemoryOrderRepository) GetAllOrdersList(ctx context.Context, is_active int8) ([]model.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []model.Order
	for _, order := range repo.orders {
		if is_active == 1 {
			if order.IsActive == 1 {
				orders = append(orders, order)
			}
		} else {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (repo *InMemoryOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]*model.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []*model.Order
	for _, order := range repo.orders {
		if order.PatientID != patient_id {
			continue
		}
		if is_active == 1 {
			if order.IsActive == 1 {
				orders = append(orders, &order)
			}
		} else {
			orders = append(orders, &order)
		}
	}
	return orders, nil
}

func (repo *InMemoryOrderRepository) UpdateOrder(ctx context.Context, order model.Order) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	repo.orders[order.ID] = order
	return nil
}

func (repo *InMemoryOrderRepository) DeleteOrder(ctx context.Context, id uint64) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.orders[id]; !exists {
		return errors.New("order not found")
	}
	delete(repo.orders, id)
	return nil
}
