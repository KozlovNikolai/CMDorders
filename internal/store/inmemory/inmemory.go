package inmemory

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sync"

	"github.com/KozlovNikolai/CMDorders/internal/models"
	"github.com/KozlovNikolai/CMDorders/internal/pkg/client"
	"github.com/KozlovNikolai/CMDorders/internal/pkg/client/restclient"
	"go.uber.org/zap"
)

type InMemoryOrderRepository struct {
	logger      *zap.Logger
	orders      map[int]models.Order
	nextID      int
	cliPatients client.IRemoteStore
	cliServices client.IRemoteStore
	mutex       sync.Mutex
}

func NewInMemoryOrderRepository(logger *zap.Logger) *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		logger:      logger,
		orders:      make(map[int]models.Order),
		nextID:      1,
		cliPatients: restclient.NewRestClient("http://localhost:8080", "/patients/", models.NewPatient(), logger),
		cliServices: restclient.NewRestClient("http://localhost:8081", "/services/", models.NewService(), logger),
	}
}

func (repo *InMemoryOrderRepository) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	patient, err := repo.cliPatients.GetByID(ctx, order.PatientID)
	if err != nil {
		return 0, err
	}
	for _, serviceID := range order.ServiceIDs {
		svc, err := repo.cliServices.GetByID(ctx, serviceID)
		if err != nil {
			return 0, err
		}
		msg := fmt.Sprintf("%v", svc)
		repo.logger.Debug("AddServiceToOrder",
			zap.String("info", msg),
		)
	}

	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order.ID = repo.nextID
	order.CreatedAt = time.Now().UTC()
	repo.nextID++
	repo.orders[order.ID] = order
	msg := fmt.Sprintf("Для пациента:\n%v\nсоздан заказ номер: %d\n", patient, order.ID)
	repo.logger.Debug("CreateOrder",
		zap.String("info", msg),
	)
	return order.ID, nil
}

func (repo *InMemoryOrderRepository) GetOrderByID(ctx context.Context, orderID int) (*models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order, exists := repo.orders[orderID]
	if !exists {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

func (repo *InMemoryOrderRepository) GetAllOrdersList(ctx context.Context, isActive bool) ([]models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []models.Order
	for _, order := range repo.orders {
		switch isActive {
		case true:
			if order.IsActive {
				orders = append(orders, order)
			}
		case false:
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (repo *InMemoryOrderRepository) GetOrdersByPatientID(ctx context.Context, patientID int, isActive bool) ([]models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []models.Order
	msg := fmt.Sprintf("InMemory p-id=%d, is-a=%v\n", patientID, isActive)
	repo.logger.Debug("GetOrdersByPatientID",
		zap.String("info", msg),
	)
	for _, order := range repo.orders {
		if order.PatientID != patientID {
			continue
		}
		if isActive {
			if order.IsActive {
				orders = append(orders, order)
			}
		} else {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

func (repo *InMemoryOrderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	repo.orders[order.ID] = order
	return nil
}
func (repo *InMemoryOrderRepository) AddServicesToOrder(ctx context.Context, orderID int, patientID int, services []int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order, exists := repo.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}
	if patientID != order.PatientID {
		return errors.New("this patient is not represented in this order")
	}
	svcs := []int{}
	for _, svc := range services {
		_, err := repo.cliServices.GetByID(ctx, svc)
		if err != nil {
			return err
		}
		svcs = append(svcs, svc)
	}
	order.ServiceIDs = append(order.ServiceIDs, svcs...)
	repo.orders[orderID] = order
	return nil
}

func (repo *InMemoryOrderRepository) DeleteOrder(ctx context.Context, id int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.orders[id]; !exists {
		return errors.New("order not found")
	}
	delete(repo.orders, id)
	return nil
}
