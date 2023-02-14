package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name              string             `bson:"name,omitempty" validate:"required"` // ชื่อ-นามสกุล
	Email             string             `bson:"email,omitempty" validate:"required,email"`
	Password          string             `bson:"password,omitempty" validate:"required"`
	Phone             string             `bson:"phone,omitempty"`
	Photo             string             `bson:"photo,omitempty"`
	Address           string             `bson:"specialization,omitempty" `
	CongenitalDisease []string           `bson:"congenitalDisease,omitempty"`
	ExerciseHistory   []ExerciseHistory  `bson:"exerciseHistory"`
}

type PatientName struct {
	En_Name string `bson:"en_name,omitempty"`
	Th_Name string `bson:"th_name,omitempty"`
}

