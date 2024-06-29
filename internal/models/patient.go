package models

import "time"

type Patient struct {
	ID         int       `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Surname    string    `json:"surname" db:"surname"`
	Name       string    `json:"name" db:"name"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
	Gender     uint8     `json:"gender" db:"gender"`
	Birthday   time.Time `json:"birthday" db:"birthday"`
}

func NewPatient() Patient {
	return Patient{}
}
