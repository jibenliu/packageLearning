package impl

import (
	"context"
	"digExam/chapter2/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Person struct {
	Db *mongo.Database
}

func (p Person) FindAll() []*schema.Person {
	var result []*schema.Person
	cur, err := p.Db.Collection("person").Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem schema.Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &elem)
	}
	return result
}
