package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017"),
		options.Client().SetAuth(options.Credential{
			Username: "admin",
			Password: "123456",
		}))
	if err != nil {
		panic("connect to mongo failed")
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	filter := bson.D{}
	databases, err := client.ListDatabaseNames(ctx, filter)
	if err != nil {
		fmt.Printf("get mongo database names failed%v\n", err)
		return
	}
	fmt.Println(databases)
	db := client.Database("mongo-test")
	tCollect := db.Collection("test-collect")
	if tCollect == nil {
		err = db.CreateCollection(ctx, "test-collect")
		if err != nil {
			fmt.Printf("create mongo collection failed%v\n", err)
			return
		}
	}
	names, err := db.ListCollectionNames(ctx, filter)
	if err != nil {
		fmt.Printf("get mongo collection names failed%v\n", err)
		return
	}
	fmt.Println(names)
	collection := db.Collection("numbers")
	res, err := collection.InsertOne(ctx, bson.D{
		{"name", "pi"},
		{"value", 3.14159},
	})
	if err != nil {
		fmt.Printf("insert collection failed%v\n", err)
		return
	}
	id := res.InsertedID
	fmt.Println(id)
}
