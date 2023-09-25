package entity

import (
	"context"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindAll(ctx context.Context) []Product
	Save(ctx context.Context, product *Product)
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func New(name, description string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
	}
}
