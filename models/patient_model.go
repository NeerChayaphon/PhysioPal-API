package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name              PatientName        `bson:"name,omitempty" validate:"required"`
	Email             string             `bson:"email,omitempty" validate:"required,email"`
	Password          string             `bson:"password,omitempty" validate:"required"`
	Phone             string             `bson:"phone,omitempty" validate:"required"`
	Photo             string             `bson:"photo,omitempty" validate:"required"`
	Address           string             `bson:"specialization,omitempty" `
	CongenitalDisease []string           `bson:"congenitalDisease"`
	ExerciseHistory   []ExerciseHistory  `bson:"exerciseHistory"`
}

type PatientName struct {
	En_Name string `bson:"en_name,omitempty" validate:"required"`
	Th_Name string `bson:"th_name,omitempty" validate:"required"`
}
