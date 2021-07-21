package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChannelType struct {
	ID            int                          `json:"id" bson:"_id"`
	Name          string                       `json:"name" bson:"name"`
	CreatedAt     primitive.DateTime           `json:"created_at" bson:"created_at"`
	UpdatedAt     primitive.DateTime           `json:"updated_at" bson:"updated_at"`
}

type ChannelPostType struct {
	ID          int                          `json:"id" bson:"_id"`
	Name        string                       `json:"name"`

	//	Description string             `json:"description"`
	CreatedAt primitive.DateTime `json:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at"`
}

type IncrementType struct {
	ID        string `json:"id" bson:"_id"`
	CurrentId int    `json:"current_id" bson:"current_id"`
}

func (l ChannelType) GetId() int {
	return l.ID
}
