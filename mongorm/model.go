package mongorm

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	Context        context.Context
	DB             *mongo.Database
	collectionName string
}
