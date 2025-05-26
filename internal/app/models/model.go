package models

type Model struct {
	ID    string  `json:"id"`
	Brand Brand   `json:"brand"`
	Name  string  `json:"name"`
	Year  int     `json:"year"`
	Price float64 `json:"price"`
}

func (b *Model) getId() string {
	return b.ID
}
