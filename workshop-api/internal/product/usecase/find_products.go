package usecase

import (
	"context"
	"github.com/weslenteche/workshop-api/internal/product/entity"
)

type FindAllProductUseCase struct {
	repository entity.ProductRepository
}

func NewFindAllProductUseCase(
	repository entity.ProductRepository,
) *FindAllProductUseCase {
	return &FindAllProductUseCase{
		repository: repository,
	}
}

func (u *FindAllProductUseCase) Handle(ctx context.Context) []entity.Product {
	return u.repository.FindAll(ctx)
}
