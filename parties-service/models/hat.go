package models

import (
	"time"
)

type Hat struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string    `json:"name" bson:"name" validate:"omitempty"`
	FirstUseAt    time.Time `json:"firstUse" bson:"firstUse"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	CanBeUseAfter time.Time `json:"canBeUseAfter" bson:"canBeUseAfter"`
	ActivePartyID string    `json:"partyId" bson:"partyId"`
}
