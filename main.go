package main

import (
	"log/slog"
	"os"

	"task-manage-api/api"

	"github.com/gin-gonic/gin"
)

// in memory storage

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("logger launched")

	svr := initRouter()
	svr.Run(":8000")

	// we might need gracefully shutdown
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthCheck", api.HealthCheck)
	return router
}
