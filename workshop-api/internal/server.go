package internal

import (
	"github.com/weslenteche/workshop-api/internal/product/controller"
	"github.com/weslenteche/workshop-api/internal/product/entity"
	"github.com/weslenteche/workshop-api/internal/product/infra/repository"
	"github.com/weslenteche/workshop-api/internal/product/usecase"
)

type Server struct {
	ProductRepository  entity.ProductRepository
	FindAllProductUc   usecase.FindAllProductUseCase
	CreateProductUc    usecase.CreateProductUseCase
	FindAllProductCtrl controller.FindAllProductController
	CreatProductCtrl   controller.CreateProductController
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	s.initRepositories()
	s.initUseCases()
	s.initControllers()
}

func (s *Server) initUseCases() {
	s.FindAllProductUc = *usecase.NewFindAllProductUseCase(s.ProductRepository)
	s.CreateProductUc = *usecase.NewCreateProductUseCase(s.ProductRepository)
}

func (s *Server) initRepositories() {
	//s.ProductRepository = repository.NewProductMemoryRepository()
	s.ProductRepository = repository.NewProductMysqlRepository()
}

func (s *Server) initControllers() {
	s.FindAllProductCtrl = *controller.NewFindAllProductController(s.FindAllProductUc)
	s.CreatProductCtrl = *controller.NewCreateProductController(s.CreateProductUc)
}
