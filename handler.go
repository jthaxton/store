package main

import (
	"net/http"
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {}

func (handler *Handler) HandleCreateDocument(c *gin.Context) {
	customId := c.DefaultQuery("customId", "")
	if len(customId) == 0 {
		c.JSON(403, gin.H{"error": "customId not found"})
	}

	id, err := Create(customId, c.Request, c)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) HandleGetDocument(c *gin.Context) {
	url := c.DefaultQuery("customId", "")
	if len(url) == 0 {
		c.JSON(403, gin.H{"error": "customId not found"})
	}
	doc, err := FindOneDocument(url)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"document": *doc})
	}
}

//Create a document
func Create(customId string, req *http.Request, c *gin.Context) (*primitive.M, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	var model interface{}
	err := json.NewDecoder(req.Body).Decode(&model)
	
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return nil, err
	}

	doc := bson.D{{Key: "id", Value: primitive.NewObjectID()}, {Key: "customId", Value: customId}, {Key: "meta", Value: model}}
	result := client.Database("documents").Collection("documents")
	_, err = result.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	// oid := res.InsertedID.
	docReturned, err := FindOneDocument(customId)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return nil, err
	}
	return docReturned, nil
}

func FindOneDocument(customId string) (*bson.M, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	var result bson.M
	err := client.Database("documents").Collection("documents").FindOne(context.TODO(), bson.D{{Key: "customId", Value: customId}}).Decode(&result)
	
	if err != nil {
		return nil, err
	}
	return &result, nil
}