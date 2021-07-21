package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResourceInterface interface {
	GetId() int
}

func RepositoryFlush(collection *mongo.Collection, ctx context.Context, subject ResourceInterface) error {
	objID := subject.GetId()

	filter := bson.D{{"_id", objID}}
	update := bson.D{{"$set", subject}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}