package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Totp_Secret struct {
	ID     primitive.ObjectID `bson:"_id" binding:"required"`
	Secret string             `bson:"secret" binding:"required"`
}
