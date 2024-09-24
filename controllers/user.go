package controllers

import (
	"context"
	"prepathon-auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByEmail(mongoClient *mongo.Client, email string) (*models.User, error) {
	userCollection := mongoClient.Database("test").Collection("users")

	filter := bson.D{
		{"email", email},
	}
	var user models.User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)

	return &user, err

}

func CreateUser(mongoClient *mongo.Client, user *models.User) error {

	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	userCollection := mongoClient.Database("test").Collection("users")

	_, err := userCollection.InsertOne(context.TODO(), user)

	return err
}
