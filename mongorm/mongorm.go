package mongorm

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}
func (m *Model) New(ctx context.Context, db *mongo.Database, collectionName string) *Model {
	m.Context = ctx
	m.DB = db
	m.collectionName = collectionName
	return m
}

func (m *Model) Create(model interface{}) error {
	collection := m.DB.Collection(m.collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := collection.InsertOne(m.Context, model)
	if err != nil {
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (m *Model) FindAll(filter interface{}, result interface{}) error {
	collection := m.DB.Collection(m.collectionName)

	cursor, err := collection.Find(m.Context, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(m.Context)

	if err = cursor.All(m.Context, result); err != nil {
		return err
	}

	return nil
}

func (m *Model) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Update(filter interface{}, update interface{}) error {
	collection := m.DB.Collection(m.collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(m.Context, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) Delete(filter interface{}) error {
	collection := m.DB.Collection(m.collectionName)
	_, err := collection.DeleteOne(m.Context, filter)
	if err != nil {
		return err
	}

	return nil
}
