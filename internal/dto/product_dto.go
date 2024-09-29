package dto

type ProductDto struct {
	DefaultDto  `json:",inline"`
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
	Categories  []string `json:"categories,omitempty" validate:"dive"`
	Price       float64  `json:"price" binding:"required"`
	Image       string   `json:"image,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}

type UpdateProductDto struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Brand       string   `json:"brand,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Price       float64  `json:"price,omitempty"`
	Image       string   `json:"image,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}
