package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID          primitive.ObjectID
	CustomId         string
	// Meta        interface{}
}

