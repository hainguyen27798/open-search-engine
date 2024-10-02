package dto

import (
	"github.com/hainguyen27798/open-search-engine/pkg/common"
)

type ProductDto struct {
	common.DefaultDto `json:",inline"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Brand             string   `json:"brand"`
	Categories        []string `json:"categories,omitempty"`
	Price             float64  `json:"price"`
	Image             string   `json:"image"`
	Popularity        int      `json:"popularity"`
}
