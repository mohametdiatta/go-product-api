package models

import "gin-learning/mongorm"

type Commentschema struct {
	mongorm.Model
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Dext  string `bson:"text"`
}

var Comment = &mongorm.Model{}
