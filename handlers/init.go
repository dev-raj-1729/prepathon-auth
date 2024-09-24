package handlers

import (
	"prepathon-auth/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	MongoClient *mongo.Client
	Session     *Session
}

type Session struct {
	User models.User
}
