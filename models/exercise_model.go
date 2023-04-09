package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseHistory struct {
	Date             time.Time          `bson:"date,omitempty"`
	ExerciseType     string             `bson:"exerciseType,omitempty" validate:"required"`
	ExerciseSetId    primitive.ObjectID `bson:"exerciseSetId,omitempty" validate:"required"`
	ExerciseRecorded []string           `bson:"exerciseRecorded,omitempty"`
	IsComplete       bool               `bson:"isComplete,omitempty" validate:"required"`
	ExerciseStatus   []ExerciseStatus   `bson:"exerciseStatus,omitempty"`
}

type ExerciseStatus struct {
	ExerciseId primitive.ObjectID `bson:"exerciseId,omitempty" validate:"required"`
	Status     string             `bson:"status,omitempty" validate:"required"`
}

// type ExerciseHistory struct {
// 	Date             time.Time          `bson:"date,omitempty"`
// 	ExerciseType     string             `bson:"exerciseType,omitempty" validate:"required"`
// 	ExerciseSetName  string             `bson:"exerciseSetName,omitempty" validate:"required"`
// 	ExerciseRecorded []ExerciseRecorded `bson:"exerciseRecorded,omitempty" validate:"required"`
// }

type ExerciseRecorded struct {
	Exercise   primitive.ObjectID `bson:"exercise" json:"exercise,omitempty"`
	TimePeriod int                `bson:"timePeriod,omitempty" validate:"required"`
	Reps       int                `bson:"reps,omitempty" validate:"required"`
	VideoLink  string             `bson:"videoLink,omitempty" validate:"required"`
}

type Exercise struct {
	Id                  primitive.ObjectID   `bson:"_id" json:"_id,omitempty"`
	Details             LanguageDescription  `bson:"details,omitempty" validate:"required"`
	MusculoskeltalTypes []primitive.ObjectID `bson:"musculoskeltalTypes" validate:"required"`
	Steps               []Steps              `bson:"steps" validate:"required"`
	Accuracy            float64              `bson:"accuracy,omitempty" validate:"required"`
}

type Steps struct {
	Image      string              `bson:"image,omitempty" validate:"required"`
	Details    LanguageDescription `bson:"details,omitempty" validate:"required"`
	ModelClass string              `bson:"modelClass,omitempty" validate:"required"`
	Model      string              `bson:"model,omitempty" validate:"required"`
	ModelIndex int                 `bson:"modelIndex,omitempty" validate:"required"`
	Timer      bool                `bson:"timer,omitempty" validate:"required"`
}

type GeneralExercise struct {
	Id                  primitive.ObjectID    `bson:"_id" json:"_id,omitempty"`
	Details             LanguageDescription   `bson:"details,omitempty" validate:"required"`
	MusculoskeltalTypes primitive.ObjectID    `bson:"musculoskeltalTypes" validate:"required"`
	Functional          *bool                 `bson:"functional" validate:"required"`
	ExerciseSet         []ExerciseSet         `bson:"exerciseSet" validate:"required"`
	Injury              []LanguageDescription `bson:"injury"`
}

// Fix error
type ExerciseSet struct {
	Exercise   primitive.ObjectID `bson:"exercise" json:"exercise,omitempty"`
	TimePeriod int                `bson:"TimePeriod,omitempty" validate:"required"`
	Reps       int                `bson:"Reps,omitempty" validate:"required"`
}

type TherapeuticExercise struct {
	Id          primitive.ObjectID  `bson:"_id" json:"_id,omitempty"`
	Details     LanguageDescription `bson:"details,omitempty" validate:"required"`
	StartDate   time.Time           `bson:"startDate" validate:"required"`
	EndDate     time.Time           `bson:"endDate" validate:"required"`
	ExerciseSet []ExerciseSet       `bson:"exerciseSet" validate:"required"`
}

type GeneralExerciseJoined struct {
	Id                  primitive.ObjectID    `bson:"_id" json:"_id,omitempty"`
	Details             LanguageDescription   `bson:"details,omitempty" validate:"required"`
	MusculoskeltalTypes primitive.ObjectID    `bson:"musculoskeltalTypes" validate:"required"`
	Functional          *bool                 `bson:"functional" validate:"required"`
	ExerciseSet         []ExerciseSetJoined   `bson:"exerciseSet" validate:"required"`
	Injury              []LanguageDescription `bson:"injury"`
}

type ExerciseSetJoined struct {
	Exercise   Exercise `bson:"exercise" json:"exercise,omitempty"`
	TimePeriod int      `bson:"timePeriod,omitempty" validate:"required"`
	Reps       int      `bson:"reps,omitempty" validate:"required"`
}
