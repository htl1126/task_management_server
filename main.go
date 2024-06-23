package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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

	Logger.Info("the server is going to start")
	router := initRouter()

	serverPort := os.Getenv("SERVERPORT")
	if serverPort != "" {
		serverPort = ":" + serverPort
	} else {
		serverPort = ":8000"
	}
	srv := &http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Error(fmt.Sprintf("listen: %v", err))
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	Logger.Info("Server is going to shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Logger.Error(fmt.Sprintf("Server shutdown: %v", err))
	}

	// need to deal with panic?
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
