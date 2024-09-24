package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id" `
	Email            string             `json:"email" bson:"email" binding:"required"`
	TwoFactorEnabled bool               `json:"two_factor_enabled" bson:"two_factor_enabled"`
}
