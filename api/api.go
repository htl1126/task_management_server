package api

import "github.com/gin-gonic/gin"

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "server is running and healthy"})
}

func CreateTask(c *gin.Context) {
	// 1. get a task number from a pool
	// 2. create an in-memory task item
	// 3. return a message with the task ID
}
