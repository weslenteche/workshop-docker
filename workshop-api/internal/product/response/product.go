package response

import "github.com/weslenteche/workshop-api/internal/product/entity"

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ToProductResponse(product entity.Product) ProductResponse {
	productResponse := ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	return productResponse
}
