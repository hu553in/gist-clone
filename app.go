package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// configure logging to console and file
	logFile, err := os.Create("gist-clone.log")
	if err == nil {
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	}

	// create router
	router := gin.New()

	// configure router to use middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// create GET /hello route
	router.GET("/hello", func(c *gin.Context) {
		name := strings.TrimSpace(c.Query("name"))

		if len(name) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "name must not be empty",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Hello, " + name + "!",
			})
		}
	})

	// temporary turn off "You trusted all proxies, this is NOT safe" warning
	router.SetTrustedProxies(nil)

	// listen to OS port
	router.Run()
}
