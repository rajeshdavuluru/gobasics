package main

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

)

type Student struct {

	StudentId     		int 	 `bson:"studentId" json:"studentId"`
	FirstName  			string 	`bson:"firstName" json:"firstName"`
	LastName 			string	`bson:"lastName" json:"lastName"`
	
}

func main(){

	log.Println("Mongo DB test");

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("test")

	coll := db.Collection("student")

	cursor, err := coll.Find(
		context.Background(),
		bson.D{{}},
	)


	var results []Student
	for cursor.Next(context.Background()) {
		log.Println(cursor.Current);
		

		var elem Student

		if err := cursor.Decode(&elem); err != nil {
			log.Fatal(err)
		}

		log.Println(elem)
		results=append(results, elem)
/*
		var s Student
		err = cursor.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}
		
		results=append(results, &s)

	/*	var doc bson.Raw
    	err = bson.Unmarshal(cursor.Current, &doc)
    	if err != nil {
            log.Fatal(err)
		}
		
		log.Println(doc.firstName)*/
		//s.firstName = cursor.Current.firstName

		//result = append(result, s)
	}
	log.Println(results)
	


}