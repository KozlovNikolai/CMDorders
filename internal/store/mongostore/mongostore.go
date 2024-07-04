package mongostore

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

// AddServicesToOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) AddServicesToOrder(ctx context.Context, order_id int, patient_id int, services []int) error {
	panic("unimplemented")
}

// CreateOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	panic("unimplemented")
}

// DeleteOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) DeleteOrder(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAllOrdersList implements store.IOrderRepository.
func (m *MongoOrderRepository) GetAllOrdersList(ctx context.Context, is_active bool) ([]models.Order, error) {
	panic("unimplemented")
}

// GetOrderByID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrderByID(ctx context.Context, order_id int) (*models.Order, error) {
	panic("unimplemented")
}

// GetOrdersByPatientID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id int, is_active bool) ([]models.Order, error) {
	panic("unimplemented")
}

// UpdateOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	panic("unimplemented")
}

func NewMongoOrderRepository(client *mongo.Client, dbName, collName string, logger *zap.Logger) *MongoOrderRepository {
	collection := client.Database(dbName).Collection(collName)
	return &MongoOrderRepository{collection: collection}
}
