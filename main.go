package main

import (
	"log/slog"
	"os"

	"task-manage-api/api"
	"task-manage-api/storage"

	"github.com/gin-gonic/gin"
)

var Logger *slog.Logger

func main() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	Logger.Info("logger launched")

	storage.StorageMgr = storage.NewStorageInstance()
	storage.StorageMgr.SetTaskPoolSize(100) // PoolSize should be from dockerfile
	storage.StorageMgr.InitTaskIDPool()
	Logger.Info("storage is ready")

	Logger.Info("the server is going to run")
	svr := initRouter()
	svr.Run(":8000")

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
