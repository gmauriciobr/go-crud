package services

import (
	dtos "crud/internal/app/dtos"
	models "crud/internal/app/models"
	repositories "crud/internal/app/repositories"

	"github.com/google/uuid"
)

type ModelService interface {
	Create(dto *dtos.CreateModelDto) (*models.Model, error)
	FindAll() ([]*models.Model, error)
	FindById(id string) (*models.Model, error)
	DeleteById(id string) error
}

type ModelServiceImpl struct {
	modelRepository repositories.CrudRepository[models.Model]
	brandService    BrandService
}

func NewModelService(repo repositories.CrudRepository[models.Model], serv BrandService) ModelService {
	return &ModelServiceImpl{
		modelRepository: repo,
		brandService:    serv,
	}
}

func (s *ModelServiceImpl) Create(dto *dtos.CreateModelDto) (*models.Model, error) {
	brand, _ := s.brandService.FindById(dto.BrandId)
	model := models.Model{
		ID:    uuid.NewString(),
		Brand: *brand,
		Name:  dto.Name,
		Year:  dto.Year,
		Price: dto.Price,
	}
	return s.modelRepository.Save(model.ID, &model)
}

func (s *ModelServiceImpl) FindAll() ([]*models.Model, error) {
	return s.modelRepository.FindAll()
}

func (s *ModelServiceImpl) FindById(id string) (*models.Model, error) {
	return s.modelRepository.FindById(id)
}

func (s *ModelServiceImpl) DeleteById(id string) error {
	return s.modelRepository.DeleteById(id)
}
