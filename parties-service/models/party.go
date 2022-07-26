package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Party struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name" validate:"omitempty"`
	Status    string             `json:"status" bson:"status"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Hats      int                `json:"hats" bson:"hats" validate:"required,numeric"`
}
