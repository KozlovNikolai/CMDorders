package client

import (
	"context"

	"github.com/KozlovNikolai/CMDorders/internal/models"
)

type IPatientsStore interface {
	GetByID(ctx context.Context, patient_id uint64) (*models.Patient, error)
	GetList(ctx context.Context) ([]models.Patient, error)
}
