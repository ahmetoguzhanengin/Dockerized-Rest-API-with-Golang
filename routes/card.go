package routes

import (
	"github.com/ahmetoguzhanengin/RestAPI/controller"
	"github.com/gin-gonic/gin"
)

func CardRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/card", controller.ListCards)
	v1.POST("/card", controller.CreateCard)
}
