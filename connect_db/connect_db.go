package connectdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "admin"
const collecon = "watchlist"

var Collection *mongo.Collection

func Startdb() {
	clietOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clietOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("mongo connected")
	Collection = client.Database(dbName).Collection(collecon)

}
func PrintCollectionData() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	for _, doc := range results {
		fmt.Println(doc) // this prints each document as a map
	}
}
