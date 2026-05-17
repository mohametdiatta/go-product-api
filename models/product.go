package models

import "gin-learning/mongorm"

type ProductSchema struct {
	mongorm.Model
	Name        string `bson:"name"`
	Price       int    `bson:"price"`
	Description string `bson:"description"`
}

var Product = &mongorm.Model{} // uniquement le Model de base
