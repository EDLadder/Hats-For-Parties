package db

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/EDLadder/Hats-For-Parties/config"
	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbconnect() *mongo.Client {
	mongoUrl, err := config.GetEnvVariable("MONGO_URL")
	if err != nil {
		log.Fatal("❌ Error fetching db url")
	}
	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("❌ Connection Failed to Database")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("❌ Connection Failed to Database")
	}
	color.Green("✅ Connected to Database")

	hatsCollection := client.Database("party-hats").Collection("hat")
	hatsCount, _ := hatsCollection.CountDocuments(context.TODO(), bson.D{
		{Key: "_id", Value: bson.D{{Key: "$exists", Value: true}}},
	})

	if hatsCount == 0 {
		envHatsCount, err := config.GetEnvVariable("HATS_COUNT")
		if err != nil {
			color.Red("❌ Error fetching hats count")
		} else {
			partyHatsCount, _ := strconv.Atoi(envHatsCount)
			for i := 1; i <= partyHatsCount; i++ {
				hatsCollection.InsertOne(context.TODO(), bson.D{
					{Key: "name", Value: "Hat" + strconv.Itoa(i)},
					{Key: "firstUse", Value: nil},
					{Key: "createdAt", Value: time.Now()},
					{Key: "canBeUseAfter", Value: time.Now()},
					{Key: "partyId", Value: nil},
				})
			}
		}
	}

	return client
}
