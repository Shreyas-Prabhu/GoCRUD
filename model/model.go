package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty`
	Name string `json:"name,omitempty" bson:"name,omitempty`
	Company string `json:"company,omitempty" bson:"company,omitempty`
}