package services

import (
	dtos "crud/internal/app/dtos"
	models "crud/internal/app/models"
	repositories "crud/internal/app/repositories"

	"github.com/google/uuid"
)

type BrandService interface {
	Create(d *dtos.CreateBrandDto) (*models.Brand, error)
	FindById(id string) (*models.Brand, error)
	FindAll() ([]*models.Brand, error)
	UpdateById(id string, d *dtos.UpdateBrandDto) (*models.Brand, error)
	DeleteById(id string) error
}

type BrandServiceImpl struct {
	repository repositories.CrudRepository[models.Brand]
}

func NewBrandService(repo repositories.CrudRepository[models.Brand]) BrandService {
	return &BrandServiceImpl{
		repository: repo,
	}
}

func (s *BrandServiceImpl) Create(d *dtos.CreateBrandDto) (*models.Brand, error) {
	brand := models.Brand{
		ID:   uuid.NewString(),
		Name: d.Name,
	}

	return s.repository.Save(brand.ID, &brand)
}

func (s *BrandServiceImpl) FindById(id string) (*models.Brand, error) {
	return s.repository.FindById(id)
}

func (s *BrandServiceImpl) FindAll() ([]*models.Brand, error) {
	return s.repository.FindAll()
}

func (s *BrandServiceImpl) UpdateById(id string, d *dtos.UpdateBrandDto) (*models.Brand, error) {
	brand, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	brand.Name = d.Name
	return s.repository.Save(brand.ID, brand)
}

func (s *BrandServiceImpl) DeleteById(id string) error {
	return s.repository.DeleteById(id)
}
