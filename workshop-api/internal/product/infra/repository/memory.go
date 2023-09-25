package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/weslenteche/workshop-api/internal/product/entity"
)

type ProductMemoryRepository struct {
	products []entity.Product
}

func NewProductMemoryRepository() entity.ProductRepository {
	return &ProductMemoryRepository{
		products: []entity.Product{
			{
				ID:          uuid.New().String(),
				Name:        "TV",
				Description: "Televisão Samsung 43 polegadas",
				Price:       1689.99,
			},
			{
				ID:          uuid.New().String(),
				Name:        "iPhone 13",
				Description: "IPhone 13 128GB",
				Price:       4449.90,
			},
		},
	}
}

func (r *ProductMemoryRepository) FindAll(ctx context.Context) []entity.Product {
	return r.products
}

func (r *ProductMemoryRepository) Save(ctx context.Context, product *entity.Product) {
	products := append(r.products, *product)
	r.products = products
}
