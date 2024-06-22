package main

import (
	"log/slog"
	"os"

	"task-manage-api/api"
	"task-manage-api/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("logger launched")

	storage.StorageMgr = storage.NewStorageInstance()
	storage.StorageMgr.SetTaskPoolSize(100) // PoolSize should be from dockerfile
	storage.StorageMgr.InitTaskIDPool()

	svr := initRouter()
	svr.Run(":8000")

	// we might need gracefully shutdown
	// need to deal with panic?
}

func initRouter() *gin.Engine {
	router := gin.Default()

	// health check
	router.GET("/healthCheck", api.HealthCheck)

	// task management APIs
	router.POST("/tasks", api.CreateTask)

	return router
}
