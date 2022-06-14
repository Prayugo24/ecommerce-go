package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/joho/godotenv"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	var mongoUrl = ""
	if username != "" && password != ""{
		mongoUrl = "mongodb://"+username+":"+password+"@"+host+":"+port
	}else{
		mongoUrl = "mongodb://"+host+":"+port
	}
	
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to connect to database :(")
		return nil
	}

	fmt.Println("Connected to database :)")
	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("ecommerce").Collection(collectionName)
	return productCollection
}
