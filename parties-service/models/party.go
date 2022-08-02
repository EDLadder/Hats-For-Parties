package models

import (
	"time"
)

type Party struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name" validate:"omitempty"`
	Status    string    `json:"status" bson:"status"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Hats      int       `json:"hats" bson:"hats" validate:"required,numeric"`
}
