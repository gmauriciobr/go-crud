package repositories

import (
	config "crud/internal/app/configs"
	models "crud/internal/app/models"
)

type CrudRepository[T models.Types] interface {
	Save(id string, entity *T) (*T, error)
	FindById(id string) (*T, error)
	FindAll() ([]*T, error)
	DeleteById(id string) error
}

type crudRepositoryImpl[T models.Types] struct {
	database config.Database[T]
}

func NewCrudRepository[T models.Types](db config.Database[T]) CrudRepository[T] {
	return &crudRepositoryImpl[T]{
		database: db,
	}
}

func (r *crudRepositoryImpl[T]) Save(id string, entity *T) (*T, error) {
	return r.database.Save(id, *entity)
}

func (r *crudRepositoryImpl[T]) FindAll() ([]*T, error) {
	return r.database.FindAll()
}

func (r *crudRepositoryImpl[T]) FindById(id string) (*T, error) {
	resp, _ := r.database.FindById(id)
	return resp, nil
}

func (r *crudRepositoryImpl[T]) DeleteById(id string) error {
	return r.database.DeleteById(id)
}
