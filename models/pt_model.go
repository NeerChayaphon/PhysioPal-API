package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Physiotherapist struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name,omitempty" validate:"required"`
	Email          string             `json:"email,omitempty" validate:"required,email"`
	Password       string             `json:"password,omitempty" validate:"required"`
	Phone          string             `json:"phone,omitempty" validate:"required"`
	Photo          string             `json:"photo,omitempty" validate:"required"`
	Specialization string             `json:"specialization,omitempty" validate:"required"`
	Background     string             `json:"background,omitempty" validate:"required"`
	Hospital       string             `json:"hospital,omitempty" validate:"required"`
}
