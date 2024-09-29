package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type IDField struct {
	ID bson.ObjectID `json:"id" bson:"_id,omitempty"`
}

type DateFields struct {
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

type DefaultModel struct {
	IDField    `bson:",inline"`
	DateFields `bson:",inline"`
}

func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	f.UpdatedAt = f.CreatedAt
	return nil
}

func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
