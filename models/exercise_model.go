package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseHistory struct {
	Date             time.Time          `bson:"date,omitempty"`
	ExerciseType     string             `bson:"exerciseType,omitempty" validate:"required"`
	ExerciseSetName  string             `bson:"exerciseSetName,omitempty" validate:"required"`
	ExerciseRecorded []ExerciseRecorded `bson:"exerciseRecorded,omitempty" validate:"required"`
}

type ExerciseRecorded struct {
	Exercise   primitive.ObjectID `bson:"exercise" json:"exercise,omitempty"`
	TimePeriod int                `bson:"timePeriod,omitempty" validate:"required"`
	Reps       int                `bson:"reps,omitempty" validate:"required"`
	VideoLink  string             `bson:"videoLink,omitempty" validate:"required"`
}

type Exercise struct {
	Id                  primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name                string             `bson:"name,omitempty" validate:"required"`
	Description         string             `bson:"description,omitempty" validate:"required"`
	MusculoskeltalTypes []string           `bson:"musculoskeltalTypes" validate:"required"`
	Injury              []Steps            `bson:"injury"`
}

type Steps struct {
	Name        string `bson:"name,omitempty" validate:"required"`
	Image       string `bson:"image,omitempty" validate:"required"`
	Description string `bson:"description,omitempty" validate:"required"`
	ModelClass  string `bson:"modelClass,omitempty" validate:"required"`
	Model       string `bson:"model,omitempty" validate:"required"`
}
