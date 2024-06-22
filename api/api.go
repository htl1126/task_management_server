package api

import (
	"fmt"
	"net/http"
	"task-manage-api/model"
	"task-manage-api/storage"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "server is running and healthy"})
}

func CreateTask(c *gin.Context) {
	var req model.CreateReq

	// get a task number from the ID pool
	taskID := storage.StorageMgr.GetTaskID()
	if taskID == -1 {
		c.JSON(500, gin.H{"message": "Cannot create a new task"})
	}

	// create an in-memory task item
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request content"})
	}
	taskItem := storage.CreateTaskItem(taskID, req.Name, model.Incompleted) // create a task with incompleted status by default
	storage.StorageMgr.WriteToTaskPool(taskItem)

	// respond with task ID
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("a task was created successfully with ID %d", taskID)})
}
