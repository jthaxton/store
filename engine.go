package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {}

func (handler *Handler) HandleCreateDocument(c *gin.Context) {
	customId := c.DefaultQuery("customId", "")

	id, err := Create(customId, c.Request)
	if err != nil {
		log.Println("Could not create document")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *Handler) HandleGetDocument(c *gin.Context) {
	url := c.DefaultQuery("customId", "")
	var doc primitive.M
	if len(url) > 0 {
		doc = FindOneDocument(url)
	}
	c.JSON(http.StatusOK, gin.H{"document": doc})
}
