package dto

type UpdateProductDto struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Brand       string   `json:"brand,omitempty"`
	Categories  []string `json:"categories,omitempty"`
	Price       float64  `json:"price,omitempty"`
	Image       string   `json:"image,omitempty"`
	Popularity  int      `json:"popularity,omitempty"`
}
