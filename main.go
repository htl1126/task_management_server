package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"task-manage-api/api"
	"task-manage-api/storage"

	"github.com/gin-gonic/gin"
)

var Logger *slog.Logger

func main() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	Logger.Info("logger launched")

	storage.StorageMgr = storage.NewStorageInstance()
	// try to get task pool size from the environment variable
	// we will use 100 if can't get the value correctly
	taskPoolSizeStr := os.Getenv("TASKPOOLSIZE")
	if taskPoolSizeStr == "" {
		storage.StorageMgr.SetTaskPoolSize(100)
		Logger.Info("task pool size set to 100")
	} else {
		taskPoolSize, err := strconv.Atoi(taskPoolSizeStr)
		if err != nil {
			storage.StorageMgr.SetTaskPoolSize(100)
			Logger.Info("task pool size set to 100")
		} else {
			storage.StorageMgr.SetTaskPoolSize(taskPoolSize)
			Logger.Info(fmt.Sprintf("task pool size set to %d", taskPoolSize))
		}
	}
	storage.StorageMgr.InitTaskIDPool()
	Logger.Info("storage is ready")

	Logger.Info("the server is going to run")
	svr := initRouter()
	serverPort := os.Getenv("SERVERPORT")
	if serverPort != "" {
		svr.Run(":" + serverPort)
	} else {
		svr.Run(":8080")
	}

	// we might need gracefully shutdown
	// need to deal with panic?
	// custom errors?
	Logger.Info("the server has been shutdown")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	Logger.Info("start registering API routes")

	// health check
	router.GET("/healthCheck", api.HealthCheck)

	// task management APIs
	router.POST("/tasks", api.CreateTask)
	router.GET("/tasks", api.GetTasks)
	router.PUT("/tasks/:id", api.UpdateTask)
	router.DELETE("/tasks/:id", api.DeleteTask)

	Logger.Info("routes were registered successfully")
	return router
}
