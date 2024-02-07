package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zajicekn/hotel-reservation/api"
	"github.com/zajicekn/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbname = "hotel-reservation"
const userCollection = "users"

func main() {
	// Load  .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Get the MONGODB_URI from .env file
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No value set in 'MONGODB_URI' environment variable.")
	}

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	// Create the MongoDB client to interact with
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Initalize handlers
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	err = app.Listen(*listenAddr)
	if err != nil {
		log.Fatal(err)
	}
}
