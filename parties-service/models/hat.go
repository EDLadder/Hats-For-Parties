package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hat struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name" validate:"omitempty"`
	FirstUseAt    primitive.DateTime `json:"firstUse" bson:"firstUse"`
	CreatedAt     primitive.DateTime `json:"createdAt" bson:"createdAt"`
	CanBeUseAfter primitive.DateTime `json:"canBeUseAfter" bson:"canBeUseAfter"`
	ActivePartyID primitive.ObjectID `json:"partyId" bson:"partyId"`
}
