package models

type Description struct {
	Name        string `bson:"name,omitempty" validate:"required"`
	Description string `bson:"description,omitempty" validate:"required"`
}

type LanguageDescription struct {
	En_Description Description `bson:"en_description,omitempty" validate:"required"`
	Th_Description Description `bson:"th_description,omitempty" validate:"required"`
}
