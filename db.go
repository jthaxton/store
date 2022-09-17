package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Create a document
func Create(customId string, req *http.Request) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var model interface{}
	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		fmt.Println(err.Error())
	}

	doc := bson.D{{Key: "id", Value: primitive.NewObjectID()}, {Key: "customId", Value: customId}, {Key: "meta", Value: model}}
	result := client.Database("documents").Collection("documents")
	res, err := result.InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Could not create Document: %v", err)
		return primitive.NilObjectID, err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	log.Println("Created document")

	return oid, nil
}

func FindOneDocument(customId string) bson.M {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var result bson.M
	err := client.Database("documents").Collection("documents").FindOne(context.TODO(), bson.D{{Key: "customId", Value: customId}}).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		if err == mongo.ErrNoDocuments {
			fmt.Println("NO DOCS FOUND")
			return nil
		}
		return nil
	}
	return result
}