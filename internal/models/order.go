package models

import "time"

type Order struct {
	ID        uint64    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	PatientID uint64    `json:"patient_id" db:"patient_id"`
	ServiceID []uint64  `json:"service_id" db:"service_id"`
	IsActive  int8      `json:"is_active" db:"is_active"`
}
