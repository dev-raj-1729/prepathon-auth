package controllers

import (
	"context"
	"prepathon-auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindSecretByUserId(mongoClient *mongo.Client, userId primitive.ObjectID) (*models.Totp_Secret, error) {
	filter := bson.D{{"_id", userId}}
	var totp_secret models.Totp_Secret
	err := mongoClient.Database("test").Collection("totp_secrets").FindOne(context.Background(), filter).Decode(&totp_secret)

	return &totp_secret, err
}

func CreateSecret(mongoClient *mongo.Client, secret string, userId primitive.ObjectID) error {

	totp_secret := models.Totp_Secret{
		ID:     userId,
		Secret: secret,
	}
	_, err := mongoClient.Database("test").Collection("totp_secrets").InsertOne(context.Background(), totp_secret)

	return err
}
