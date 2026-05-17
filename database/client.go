package database

import (
	"context"
	"gin-learning/mongorm"

	"go.mongodb.org/mongo-driver/mongo"
)

type ModelRegistery struct {
	Model          map[string]*mongorm.Model
	Context        context.Context
	DB             *mongo.Database
	collectionName string
}

func (r *ModelRegistery) New(ctx context.Context, db *mongo.Database) {
	r.Context = ctx
	r.DB = db

}
func (r *ModelRegistery) Register(m map[string]*mongorm.Model, collectionName string) {
	r.Model = make(map[string]*mongorm.Model)
	r.Model = m
	r.collectionName = collectionName

}
func (r *ModelRegistery) Init() *ModelRegistery {
	for key, value := range r.Model {
		r.Model[key] = value.New(r.Context, r.DB, r.collectionName)
	}
	return r

}
