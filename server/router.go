package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hu553in/gist-clone/config"
	"github.com/hu553in/gist-clone/controllers"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.GetConfig().GetString("gin.mode"))

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)

	configureControllers(router)

	return router
}

func configureControllers(router *gin.Engine) {
	health := new(controllers.HealthController)
	router.GET("/health", health.Health)
}
