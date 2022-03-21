package controllers

import (
	"github.com/gin-gonic/gin"
	//     "context"
	//     "time"
	//     "go.mongodb.org/mongo-driver/mongo"
	//     "go.mongodb.org/mongo-driver/mongo/options"
	//     "go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetOne(c *gin.Context) {
	c.JSON(200, gin.H{"data": "hello here"})
}
