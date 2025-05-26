package dtos

type CreateModelDto struct {
	BrandId string  `json:"brand_id"`
	Name    string  `json:"name"`
	Year    int     `json:"year"`
	Price   float64 `json:"price"`
}
