package models

import (
	"github.com/hainguyen27798/open-search-engine/pkg/common"
)

type Product struct {
	common.DefaultModel `bson:",inline"`
	Name                string   `bson:"name" json:"name"`
	Description         string   `bson:"description" json:"description"`
	Brand               string   `bson:"brand" json:"brand"`
	Categories          []string `bson:"categories,omitempty" json:"categories,omitempty"`
	Price               float64  `bson:"price" json:"price"`
	Image               string   `bson:"image" json:"image"`
	Popularity          int      `bson:"popularity" json:"popularity"`
}
