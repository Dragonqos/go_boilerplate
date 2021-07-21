package repository

import (
	"github.com/Dragonqos/go_boilerplate/api/core/db"
	"github.com/Dragonqos/go_boilerplate/api/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	ChannelRepo ChannelRepositoryInterface
)

type ChannelRepositoryInterface interface {
	FindAll() []model.ChannelType
	FindById(id int) *model.ChannelType
	DeleteById(id int) bool
	Insert(channel *model.ChannelPostType) *int32
	GetNextId() *int
	Flush(action *model.ChannelType) error
}

type channelRepository struct {
	dbs *mongo.Database
	ctx context.Context
	collectionName string
}

func CreateChannelRepository(db *mongo.Database, ctx context.Context) *ChannelRepositoryInterface {
	if nil == ChannelRepo {
		ChannelRepo = &channelRepository{
			dbs: db,
			ctx: ctx,
			collectionName: "channels",
		}
	}

	return &ChannelRepo
}

func (r channelRepository) FindAll() []model.ChannelType {
	var results []model.ChannelType

	collection := r.dbs.Collection(r.collectionName)

	filter := bson.D{{}}
	cur, err := collection.Find(r.ctx, filter)

	if err != nil {
		fmt.Println(err)
		return results
	}

	for cur.Next(r.ctx) {
		var result model.ChannelType
		err = cur.Decode(&result)
		if err != nil {
			fmt.Println("Unable to decode element")
			return []model.ChannelType{}
		}

		results = append(results, result)
	}

	_ = cur.Close(r.ctx)

	return results
}

func (r channelRepository) FindById(id int) *model.ChannelType {
	var result model.ChannelType

	collection := r.dbs.Collection(r.collectionName)

	filter := bson.D{{"_id", id}}
	err := collection.FindOne(r.ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &result
}

func (r channelRepository) DeleteById(id int) bool {
	collection := r.dbs.Collection(r.collectionName)

	filter := bson.D{{"_id", id}}
	result, err := collection.DeleteOne(r.ctx, filter)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if result.DeletedCount == 0 {
		fmt.Println(err)
		return result.DeletedCount == 0
	}

	return result.DeletedCount > 0
}

func (r channelRepository) Insert(channel *model.ChannelPostType) *int32{
	now := time.Now()
	channel.CreatedAt = primitive.NewDateTimeFromTime(now)
	channel.UpdatedAt = primitive.NewDateTimeFromTime(now)

	collection := r.dbs.Collection(r.collectionName)
	res, err := collection.InsertOne(r.ctx, channel)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	result := res.InsertedID.(int32)
	//result := res.InsertedID.(primitive.ObjectID).Hex()
	return &result
}


func (r channelRepository) GetNextId() *int {
	var increment model.IncrementType

	incrementName := "channel_increment_ids"
	collection := r.dbs.Collection(incrementName)

	filter := bson.D{{"_id", "channels"}}
	insert := bson.M{
		"$setOnInsert": bson.M{
			"current_id": 0,
		},
	}
	update := bson.M{
		"$inc": bson.M{
			"current_id": 1,
		},
	}
	upsert := true
	opt := options.UpdateOptions{
		Upsert:         &upsert,
	}

	_, err := collection.UpdateOne(r.ctx, filter, insert, &opt)
	if err != nil {
		fmt.Println("insert initial fail", err)
		return nil
	}
	_, err = collection.UpdateOne(r.ctx, filter, update)
	if err != nil {
		fmt.Println("increment fail", err)
		return nil
	}

	filter = bson.D{{"_id", "channels"}}
	err = collection.FindOne(r.ctx, filter).Decode(&increment)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &increment.CurrentId
}

func (r channelRepository) Flush(channel *model.ChannelType) error {
	collection := r.dbs.Collection(r.collectionName)
	return db.RepositoryFlush(collection, r.ctx, channel)
}