package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Physiotherapist struct {
	Id       primitive.ObjectID    `bson:"_id" json:"_id,omitempty"`
	Details  PTLanguageDescription `bson:"details,omitempty" validate:"required"`
	Email    string                `json:"email,omitempty" validate:"required"`
	Password string                `json:"password,omitempty" validate:"required"`
	Phone    string                `json:"phone,omitempty" validate:"required"`
	Photo    string                `json:"photo,omitempty" validate:"required"`
}

type PTDescription struct {
	Name           string `bson:"name,omitempty" validate:"required"`
	Specialization string `bson:"specialization,omitempty" validate:"required"`
	Background     string `bson:"background,omitempty" validate:"required"`
	Hospital       string `bson:"hospital,omitempty" validate:"required"`
}

type PTLanguageDescription struct {
	En_Description PTDescription `bson:"en_description,omitempty" validate:"required"`
	Th_Description PTDescription `bson:"th_description,omitempty" validate:"required"`
}
