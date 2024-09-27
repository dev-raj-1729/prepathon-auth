package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id" `
	Email    string             `json:"email" bson:"email" binding:"required"`
	Name     string             `json:"name" bson:"name"`
	PhotoURL string             `json:"photo_url" bson:"photo_url"`
}
