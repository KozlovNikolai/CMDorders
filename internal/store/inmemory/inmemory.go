package inmemory

import (
	"context"
	"errors"
	"fmt"
	"time"

	"sync"

	"github.com/KozlovNikolai/CMDorders/internal/client"
	"github.com/KozlovNikolai/CMDorders/internal/client/restclient"
	"github.com/KozlovNikolai/CMDorders/internal/models"
	"go.uber.org/zap"
)

type InMemoryOrderRepository struct {
	logger      *zap.Logger
	orders      map[uint64]models.Order
	nextID      uint64
	cliPatients client.IRemoteStore
	cliServices client.IRemoteStore
	mutex       sync.Mutex
}

func NewInMemoryOrderRepository(logger *zap.Logger) *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		logger:      logger,
		orders:      make(map[uint64]models.Order),
		nextID:      1,
		cliPatients: restclient.NewRestClient("http://localhost:8080", "/patients/", models.NewPatient(), logger),
		cliServices: restclient.NewRestClient("http://localhost:8081", "/services/", models.NewService(), logger),
	}
}

func (repo *InMemoryOrderRepository) CreateOrder(ctx context.Context, order models.Order) (uint64, error) {
	patient, err := repo.cliPatients.GetByID(ctx, uint64(order.Patient.ID))
	if err != nil {
		return 0, err
	}
	for _, service := range order.Services {
		svc, err := repo.cliServices.GetByID(ctx, uint64(service.ID))
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

func (repo *InMemoryOrderRepository) GetOrderByID(ctx context.Context, order_id uint64) (*models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order, exists := repo.orders[order_id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

func (repo *InMemoryOrderRepository) GetAllOrdersList(ctx context.Context, is_active int8) ([]models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []models.Order
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

func (repo *InMemoryOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]models.Order, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var orders []models.Order
	msg := fmt.Sprintf("InMemory p-id=%d, is-a=%d\n", patient_id, is_active)
	repo.logger.Debug("GetOrdersByPatientID",
		zap.String("info", msg),
	)
	for _, order := range repo.orders {
		if uint64(order.Patient.ID) != patient_id {
			continue
		}
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

func (repo *InMemoryOrderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.orders[order.ID]; !exists {
		return errors.New("order not found")
	}
	repo.orders[order.ID] = order
	return nil
}
func (repo *InMemoryOrderRepository) AddServicesToOrder(ctx context.Context, order_id uint64, patient_id uint64, services []models.Service) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	order, exists := repo.orders[order_id]
	if !exists {
		return errors.New("order not found")
	}
	if patient_id != uint64(order.Patient.ID) {
		return errors.New("this patient is not represented in this order")
	}
	svcs := []models.Service{}
	var servicesCasted []models.Service
	var i interface{} = services
	if s, ok := i.([]models.Service); ok {
		servicesCasted = s
	} else {
		return errors.New("input parametr not equal []model.Service")
	}
	for _, svc := range servicesCasted {
		_, err := repo.cliServices.GetByID(ctx, svc.ID)
		if err != nil {
			return err
		}
		svcs = append(svcs, svc)
	}
	order.Services = append(order.Services, svcs...)
	repo.orders[order_id] = order
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
