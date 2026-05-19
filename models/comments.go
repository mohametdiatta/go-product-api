package models

import "gin-learning/mongorm"

type Commentschema struct {
	mongorm.Model
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Text  string `bson:"text" json:"text"`
}

var Comment = &mongorm.Model{}
