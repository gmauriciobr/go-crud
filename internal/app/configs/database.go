package configs

import (
	models "crud/internal/app/models"
)

type Database[T models.Types] interface {
	Save(key string, value T) (*T, error)
	FindById(key string) (*T, bool)
	FindAll() ([]*T, error)
	DeleteById(key string) error
}
