package mongostore

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

// CreateOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) CreateOrder(ctx context.Context, order model.Order) (uint64, error) {
	panic("unimplemented")
}

// DeleteOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) DeleteOrder(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// GetAllOrdersList implements store.IOrderRepository.
func (m *MongoOrderRepository) GetAllOrdersList(ctx context.Context, is_active int8) ([]*model.Order, error) {
	panic("unimplemented")
}

// GetOrderByID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrderByID(ctx context.Context, order_id uint64) (*model.Order, error) {
	panic("unimplemented")
}

// GetOrdersByPatientID implements store.IOrderRepository.
func (m *MongoOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id uint64, is_active int8) ([]*model.Order, error) {
	panic("unimplemented")
}

// UpdateOrder implements store.IOrderRepository.
func (m *MongoOrderRepository) UpdateOrder(ctx context.Context, order model.Order) error {
	panic("unimplemented")
}

func NewMongoOrderRepository(client *mongo.Client, dbName, collName string) *MongoOrderRepository {
	collection := client.Database(dbName).Collection(collName)
	return &MongoOrderRepository{collection: collection}
}

// func (repo *MongoOrderRepository) CreateEmployer(ctx context.Context, order model.Order) (int, error) {
// 	result, err := repo.collection.InsertOne(ctx, order)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result.InsertedID.(int), nil
// }

// func (repo *MongoEmployerRepository) GetEmployerByID(ctx context.Context, id int) (models.Employer, error) {
// 	var employer models.Employer
// 	filter := bson.D{{Key: "id", Value: id}}
// 	err := repo.collection.FindOne(ctx, filter).Decode(&employer)
// 	if err == mongo.ErrNoDocuments {
// 		return employer, errors.New("employer not found")
// 	}
// 	return employer, err
// }

// func (repo *MongoEmployerRepository) GetAllEmployers(ctx context.Context) ([]models.Employer, error) {
// 	var employers []models.Employer
// 	cursor, err := repo.collection.Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var employer models.Employer
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

// func (repo *MongoEmployerRepository) UpdateEmployer(ctx context.Context, id int, employer models.Employer) error {
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
