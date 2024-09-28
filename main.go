package main

import (
	"context"
	"log"
	"os"
	"prepathon-auth/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// MongoDB initialization
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal(" 'MONGODB_URI' environment variable not set")
	}

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	handler := &handlers.Handler{
		MongoClient: mongoClient,
	}

	// Fiber initialization
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/token/user", handler.CreateUserWithFirebaseToken)
	app.Get("/token/user", handler.FindUserWithFirebaseToken)

	app.Post("enable_2fa", handler.Enable2FA)
	app.Get("/verify_2fa", handler.Verify2FA)

	port := os.Getenv("PORT")
	port = ":" + port
	app_env := os.Getenv("APP_ENV")
	if app_env == "dev" {
		// only listen on localhost for development
		app.Listen("localhost" + port)
	} else {
		app.Listen(port)
	}
}
