package models

type Service struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
}

func NewService() Service {
	return Service{}
}
