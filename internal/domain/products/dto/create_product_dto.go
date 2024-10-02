package dto

type CreateProductDto struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Brand       string   `json:"brand" binding:"required"`
	Categories  []string `json:"categories,omitempty" validate:"dive"`
	Price       float64  `json:"price" binding:"required"`
	Image       string   `json:"image,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}
