package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name              string             `bson:"name,omitempty" validate:"required"`
	Email             string             `bson:"email,omitempty" validate:"required,email"`
	Password          string             `bson:"password,omitempty" validate:"required"`
	Phone             string             `bson:"phone,omitempty" validate:"required"`
	Photo             string             `bson:"photo,omitempty" validate:"required"`
	Address           string             `bson:"specialization,omitempty" `
	CongenitalDisease []string           `bson:"congenitalDisease"`
	ExerciseHistory   []ExerciseHistory  `bson:"exerciseHistory" validate:"required"`
}
