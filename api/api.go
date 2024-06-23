package api

import (
	"fmt"
	"net/http"
	"strconv"
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
		return
	}
	taskItem := storage.CreateTaskItem(taskID, req.Name, model.Incompleted) // create a task with incompleted status by default
	storage.StorageMgr.WriteToTaskPool(taskItem)                            // task ID was newly generated, it's safe to write to the pool

	// respond with task ID
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("a task was created successfully with ID %d", taskID)})
}

func GetTasks(c *gin.Context) {
	taskItems := storage.StorageMgr.GetTaskPool()

	taskList := make(model.TaskItemList, len(taskItems))
	idx := 0
	for _, v := range taskItems {
		taskList[idx] = v
		idx += 1
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   fmt.Sprintf("%d task(s) in total", len(taskList)),
		"task_list": taskList,
	})
}

func UpdateTask(c *gin.Context) {
	var (
		req model.UpdateReq
		err error
	)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Incorrect request content"})
		return
	}

	if !storage.StorageMgr.TaskPoolHasID(id) {
		c.JSON(http.StatusOK, gin.H{"message": "task not found"})
		return
	}

	if req.Name == nil && req.Status == nil {
		c.JSON(http.StatusOK, gin.H{"message": "name and status are not given, updated nothing"})
		return
	}

	taskItem := storage.StorageMgr.GetTaskItemByID(id)
	if req.Name != nil {
		taskItem.Name = *req.Name
	}
	if req.Status != nil {
		taskItem.Status = *req.Status
	}
	storage.StorageMgr.WriteToTaskPool(taskItem) // the task ID was obtained from the pool, so we can write back to it

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID"})
		return
	}

	if !storage.StorageMgr.TaskPoolHasID(id) {
		c.JSON(http.StatusOK, gin.H{"message": "task not found"})
		return
	}

	storage.StorageMgr.DeleteFromTaskPool(id)
	storage.StorageMgr.RecycleTaskID(id)

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}
