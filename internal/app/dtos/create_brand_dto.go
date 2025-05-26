package dtos

type CreateBrandDto struct {
	Name string `json:"name" validate:"required"`
}
