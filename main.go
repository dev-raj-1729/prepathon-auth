package main

import (
	"context"
	"log"
	"os"
	"prepathon-auth/handlers"

	"github.com/gofiber/fiber/v2"
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

	app.Post("/token/user", handler.CreateUserWithFirebaseToken)
	app.Get("/token/user", handler.FindUserWithFirebaseToken)

	app.Post("enable_2fa", handler.Enable2FA)
	app.Get("/verify_2fa", handler.Verify2FA)

	// const token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImUwM2E2ODg3YWU3ZjNkMTAyNzNjNjRiMDU3ZTY1MzE1MWUyOTBiNzIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vcHJlcGF0aG9uLTYyZGNmIiwiYXVkIjoicHJlcGF0aG9uLTYyZGNmIiwiYXV0aF90aW1lIjoxNzI3MDgwNjI3LCJ1c2VyX2lkIjoiMjM2UWhQYnZvZVFNWFMwVGo2UkE0QjRDYzhyMSIsInN1YiI6IjIzNlFoUGJ2b2VRTVhTMFRqNlJBNEI0Q2M4cjEiLCJpYXQiOjE3MjcxMjE3NjYsImV4cCI6MTcyNzEyNTM2NiwiZW1haWwiOiJkZXYucmFqLnIxNzI5QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImRldi5yYWoucjE3MjlAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.v5XMliptBkBKfwnTzunujepzCEaxJWj9_WDhgpLHJeJcatSbhcmWaIAVFr11z_os0y1UpgvRRwHrfCTnulGtrx2LOvGs8hq9WwJ9gbQPsrai7GtcYexbaK5gkPXEuoG3s4xbdKhMYJRVu4lfbGDgdLmRgCGJkvWgW3d7oWaYs2RRdEti_pku98u1nwlR0osXI2XXpLvO2-hr_ixcZ_Lvs6J5ql9sEy4UGPxxhW5LDGus6yTjwjVJpj3JzY5pYVjEv5nU3vf8gztO4jFJUp4DtSHuEw04V0JzzZ3E7QEe0E6z_kfXg9tvz3-SSDJUkwxRnZS8EfSwT72Xxm6n-hBIJg"
	// claims, err := utils.VerifyIDToken(token)
	// fmt.Println(claims.EmailVerified)
	// fmt.Println(err)
	// jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	// 	fmt.Println(token.Header["kid"])
	// 	return token, nil
	// })
	// app.Post("a")
	app.Listen("localhost:8080")
}
