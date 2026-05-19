package models

import "github.com/mohametdiatta/gormongo"

type Commentschema struct {
	gormongo.Model
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Text  string `bson:"text" json:"text"`
}

var Comment = &gormongo.Model{}
