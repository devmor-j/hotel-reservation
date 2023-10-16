package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/devmor-j/hotel-reservation/api"
	"github.com/devmor-j/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	coll := client.Database(dbname).Collection(userColl)

	ctx = context.Background()
	user := types.User{
		Firstname: "morteza",
		Lastname:  "jamshidi",
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	user = types.User{}
	if err = coll.FindOne(ctx, bson.M{}).Decode(&user); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", user)

	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
