package controller

import (
	"net/http"

	"github.com/ahmetoguzhanengin/RestAPI/config"
	"github.com/ahmetoguzhanengin/RestAPI/models"
	"github.com/gin-gonic/gin"
)

func ListCards(ctx *gin.Context) {
	cards := []models.Card{}
	config.DB.Find(&cards)
	ctx.JSON(http.StatusOK, &cards)
}
func CreateCard(ctx *gin.Context) {
	var card models.Card
	ctx.BindJSON(&card)

	if err := config.DB.Create(&card).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &card)
	}

}
