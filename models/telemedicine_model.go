package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Telemedicine struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	RoomID          string             `bson:"roomID,omitempty" validate:"required"`
	Patient         primitive.ObjectID `bson:"patient,omitempty"`
	Physiotherapist primitive.ObjectID `bson:"physiotherapist,omitempty"`
	Date            time.Time          `bson:"date,omitempty"`
}
