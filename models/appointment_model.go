package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Patient         primitive.ObjectID `bson:"patient,omitempty"`
	Physiotherapist primitive.ObjectID `bson:"physiotherapist,omitempty"`
	Date            time.Time          `bson:"date,omitempty"`
	Injury          string             `bson:"injury,omitempty"`
	Treatment       string             `bson:"treatment,omitempty"`
	Therapeutic     Therapeutic        `bson:"therapeutic,omitempty"`
}
