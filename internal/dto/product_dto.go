package dto

type ProductDto struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Categories  []string `json:"categories,omitempty"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Popularity  int      `json:"popularity"`
}

type CreateProductDto struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Brand       string   `json:"brand" binding:"required"`
	Categories  []string `json:"categories,omitempty"`
	Price       float64  `json:"price" binding:"required"`
	Image       string   `json:"image,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}
