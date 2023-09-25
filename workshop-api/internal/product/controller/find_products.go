package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weslenteche/workshop-api/internal/product/response"
	"github.com/weslenteche/workshop-api/internal/product/usecase"
	"net/http"
)

type FindAllProductController struct {
	usecase usecase.FindAllProductUseCase
}

func NewFindAllProductController(
	usecase usecase.FindAllProductUseCase,
) *FindAllProductController {
	return &FindAllProductController{
		usecase: usecase,
	}
}

func (fp *FindAllProductController) Handle(c *gin.Context) {
	products := fp.usecase.Handle(c.Request.Context())
	productsResponse := make([]response.ProductResponse, len(products))
	for i, product := range products {
		productsResponse[i] = response.ToProductResponse(product)
	}
	c.JSON(http.StatusOK, productsResponse)
	return
}
