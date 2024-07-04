package models

import "time"

type Order struct {
	ID           int       `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	PatientID    int       `json:"patient_id" db:"patient_id"`
	ServiceIDs   []int     `json:"service_id" db:"service_id"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	InProcessing bool      `json:"in_processing" db:"in_processing"`
}

// type Order struct {
// 	ID        uint64    `json:"id" db:"id"`
// 	CreatedAt time.Time `json:"created_at" db:"created_at"`
// 	Patient   Patient   `json:"patient" db:"patient"`
// 	Services  []Service `json:"services" db:"services"`
// 	IsActive  int8      `json:"is_active" db:"is_active"`
// }
