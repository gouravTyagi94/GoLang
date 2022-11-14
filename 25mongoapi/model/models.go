package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct{
	//This id is _id which is provided by mongodb itself
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string `json:"movie,omitempty"`
	Watched bool `json:"watched,omitempty"`
}