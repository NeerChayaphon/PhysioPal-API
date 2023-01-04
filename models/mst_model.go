package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MusculoskeltalTypes struct {
	Id   primitive.ObjectID  `bson:"_id" json:"_id,omitempty"`
	Type LanguageDescription `bson:"type,omitempty" validate:"required"`
}
