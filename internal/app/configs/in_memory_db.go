package configs

import (
	models "crud/internal/app/models"
	"fmt"
	"sync"
)

type InMemoryDB[T models.Brand | models.Model] struct {
	data map[string]T
	mu   sync.RWMutex
}

func NewInMemoryDB[T models.Brand | models.Model]() *InMemoryDB[T] {
	return &InMemoryDB[T]{
		data: make(map[string]T),
	}
}

func (db *InMemoryDB[T]) Save(key string, value T) (*T, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
	return &value, nil
}

func (db *InMemoryDB[T]) FindById(key string) (*T, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	val, ok := db.data[key]
	return &val, ok
}

func (db *InMemoryDB[T]) FindAll() ([]*T, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var result []*T
	for _, value := range db.data {
		v := value // Create a copy to
		result = append(result, &v)
	}
	return result, nil
}

func (db *InMemoryDB[T]) DeleteById(key string) error {
	_, exist := db.FindById(key)
	if !exist {
		return fmt.Errorf("id not found `%s`", key)
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, key)
	return nil
}
