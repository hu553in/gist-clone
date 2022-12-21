package server

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hu553in/gist-clone/config"
)

func Init() {
	configureLogging()
	NewRouter().Run(":" + config.GetConfig().GetString("server.port"))
}

func configureLogging() {
	logFile, err := os.Create(config.GetConfig().GetString("log.path"))
	if err != nil {
		log.Fatal("Failed to create log file", err)
	}

	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}
