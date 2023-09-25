package usecase

import (
	"context"
	"github.com/weslenteche/workshop-api/internal/product/entity"
)

type CreateProductUseCase struct {
	repository entity.ProductRepository
}

type CreateProductInput struct {
	Name        string
	Description string
	Price       float64
}

func NewCreateProductUseCase(
	repository entity.ProductRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		repository: repository,
	}
}

func (u *CreateProductUseCase) Handle(ctx context.Context, input CreateProductInput) entity.Product {
	product := entity.New(input.Name, input.Description, input.Price)
	u.repository.Save(ctx, product)
	return *product
}
