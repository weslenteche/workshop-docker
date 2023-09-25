package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/weslenteche/workshop-api/internal/product/response"
	"github.com/weslenteche/workshop-api/internal/product/usecase"
	"net/http"
)

type CreateProductController struct {
	usecase usecase.CreateProductUseCase
}

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCreateProductController(
	usecase usecase.CreateProductUseCase,
) *CreateProductController {
	return &CreateProductController{
		usecase: usecase,
	}
}

func (cp *CreateProductController) Handle(c *gin.Context) {
	var req CreateProductRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	product := cp.usecase.Handle(c.Request.Context(), usecase.CreateProductInput{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})

	c.JSON(http.StatusCreated, response.ToProductResponse(product))
	return
}
