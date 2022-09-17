package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Engine struct {}

func (e *Engine) HandleCreateDocument(c *gin.Context) {
	// if c.Request.Method != "POST" {
	// 	return
	// }
	
	doc := Document{}
	customId := c.DefaultQuery("customId", "")
	err := c.ShouldBindJSON(&doc);
	if err != nil {
		log.Print(err)
		log.Println("Could not bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := Create(&doc, customId, c.Request)
	if err != nil {
		log.Println("Could not create document")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (e *Engine) HandleGetDocument(c *gin.Context) {
	url := c.DefaultQuery("customId", "")
	var doc primitive.M
	if len(url) > 0 {
		doc = FindOneDocument(url)
	}
	c.JSON(http.StatusOK, gin.H{"document": doc})
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/tasks/", handleGetTasks)
// 	r.Run()
// }