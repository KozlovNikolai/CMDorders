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
func (m *MongoOrderRepository) AddServicesToOrder(ctx context.Context, order_id uint64, patient_id uint64, services []models.Service) error {
	panic("unimplemented")
}

// CreateOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) CreateOrder(ctx context.Context, order models.Order) (uint64, error) {
	panic("unimplemented")
}

// DeleteOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) DeleteOrder(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// GetAllOrdersList implements store.IOrderRepository.
func (m *MongoOrderRepository) GetAllOrdersList(ctx context.Context, is_active int8) ([]models.Order, error) {
	panic("unimplemented")
}

// GetOrderByID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrderByID(ctx context.Context, order_id uint64) (*models.Order, error) {
	panic("unimplemented")
}

// GetOrdersByPatientID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]models.Order, error) {
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

// func (repo *MongoOrderRepository) CreateEmployer(ctx context.Context, order models.Order) (int, error) {
// 	result, err := repo.collection.InsertOne(ctx, order)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result.InsertedID.(int), nil
// }

// func (repo *MongoEmployerRepository) GetEmployerByID(ctx context.Context, id int) (modelss.Employer, error) {
// 	var employer modelss.Employer
// 	filter := bson.D{{Key: "id", Value: id}}
// 	err := repo.collection.FindOne(ctx, filter).Decode(&employer)
// 	if err == mongo.ErrNoDocuments {
// 		return employer, errors.New("employer not found")
// 	}
// 	return employer, err
// }

// func (repo *MongoEmployerRepository) GetAllEmployers(ctx context.Context) ([]modelss.Employer, error) {
// 	var employers []modelss.Employer
// 	cursor, err := repo.collection.Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var employer modelss.Employer
// 		err := cursor.Decode(&employer)
// 		if err != nil {
// 			return nil, err
// 		}
// 		employers = append(employers, employer)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}
// 	return employers, nil
// }

// func (repo *MongoEmployerRepository) UpdateEmployer(ctx context.Context, id int, employer modelss.Employer) error {
// 	filter := bson.D{{Key: "id", Value: id}}
// 	update := bson.D{{Key: "$set", Value: employer}}
// 	_, err := repo.collection.UpdateOne(ctx, filter, update)
// 	return err
// }

// func (repo *MongoEmployerRepository) DeleteEmployer(ctx context.Context, id int) error {
// 	filter := bson.D{{Key: "id", Value: id}}
// 	_, err := repo.collection.DeleteOne(ctx, filter)
// 	return err
// }
