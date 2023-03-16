package main

import (
	"github.com/ahmetoguzhanengin/RestAPI/config"
	"github.com/ahmetoguzhanengin/RestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.CardRoute(router)
	router.Run("0.0.0.0:8181")
}
