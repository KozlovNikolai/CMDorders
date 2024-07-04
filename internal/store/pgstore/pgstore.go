package pgstore

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type PostgresOrderRepository struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

// AddServicesToOrder implements store.IOrderRepository.
func (repo *PostgresOrderRepository) AddServicesToOrder(ctx context.Context, order_id int, patient_id int, services []int) error {
	panic("unimplemented")
}

func NewPostgresOrderRepository(db *pgxpool.Pool, logger *zap.Logger) *PostgresOrderRepository {
	return &PostgresOrderRepository{
		logger: logger,
		db:     db,
	}
}

// DeleteOrder(ctx context.Context, id int) error

func (repo *PostgresOrderRepository) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	var id int
	query := `
		INSERT INTO orders (created_at, patient_id, service_id, is_active) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`
	err := repo.db.QueryRow(ctx, query, order.CreatedAt, order.PatientID, order.ServiceIDs, order.IsActive).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *PostgresOrderRepository) GetOrderByID(ctx context.Context, order_id int) (*models.Order, error) {
	var order models.Order
	query := `
		SELECT id, created_at, patient_id, service_id, is_active 
		FROM orders
		WHERE id=$1`
	row := repo.db.QueryRow(ctx, query, order_id)
	err := row.Scan(&order.ID, &order.CreatedAt, &order.PatientID, &order.ServiceIDs, &order.IsActive)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *PostgresOrderRepository) GetOrdersByPatientID(ctx context.Context, patient_id int, is_active bool) ([]models.Order, error) {
	var orders []models.Order
	var query string
	if is_active {
		query = `
		SELECT id, created_at, patient_id, service_id, is_active 
		FROM orders
		WHERE patient_id=$1 AND is_active=$2`
	} else {
		query = `
		SELECT id, created_at, patient_id, service_id, is_active 
		FROM orders
		WHERE patient_id=$1`
	}
	rows, err := repo.db.Query(ctx, query, patient_id, is_active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.CreatedAt, &order.PatientID, &order.ServiceIDs, &order.IsActive)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
func (repo *PostgresOrderRepository) GetAllOrdersList(ctx context.Context, is_active bool) ([]models.Order, error) {
	var orders []models.Order
	var query string
	if is_active {
		query = `
		SELECT id, created_at, patient_id, service_id, is_active 
		FROM orders
		WHERE is_active=$1`
	} else {
		query = `
		SELECT id, created_at, patient_id, service_id, is_active 
		FROM orders`
	}
	rows, err := repo.db.Query(ctx, query, is_active)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.CreatedAt, &order.PatientID, &order.ServiceIDs, &order.IsActive)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *PostgresOrderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	query := `
		UPDATE orders 
		SET created_at=$1, patient_id=$2, service_id=$3, is_active=$4 
		WHERE id=$5`
	_, err := repo.db.Exec(ctx, query, order.CreatedAt, order.PatientID, order.ServiceIDs, order.IsActive, order.ID)
	return err
}

func (repo *PostgresOrderRepository) DeleteOrder(ctx context.Context, id int) error {
	query := "DELETE FROM orders WHERE id=$1"
	_, err := repo.db.Exec(ctx, query, id)
	return err
}
