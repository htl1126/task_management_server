package main

import (
	"log/slog"
	"os"

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
	router.GET("/healthCheck", healthCheck)
	return router
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"message": "server is running and healthy"})
}
