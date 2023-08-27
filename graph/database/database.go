package database

import (
	"context"
	"log"
	"time"

	"gitlab.com/pragmaticreviews/graphql-server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://msbalaji10875:abirami@cluster0.0gkyido.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}
func (db *DB) Save(input model.NewDog) *model.Dog {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db.client.Database("animals").Collection("dogs")
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return &model.Dog{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		IsGoodBoi: input.IsGoodBoi,
	}
}

func (db *DB) FindById(ID string) *model.Dog {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	collection := db.client.Database("animals").Collection("dogs")
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}

	dog := model.Dog{}
	res.Decode(&dog)
	return &dog

}
