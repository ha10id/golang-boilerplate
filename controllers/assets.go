package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

type AssetsController struct{}

var assetsModel = new(models.Assets)

func (u AssetsController) GetAll(filter string, c *gin.Context) {
	assets, err := assetsModel.GetAll(filter)
	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": assets[0]})
	return
}

func (u AssetsController) GetMainBanners(c *gin.Context) {
	u.GetAll("main_banner", c)
	//assets, err := assetsModel.GetAll("main_banner")
	//if err != nil {
	//	c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
	//	c.Abort()
	//	return
	//}
	//
	//c.JSON(200, gin.H{"data": assets})
	return
}

func (u AssetsController) GetMainActions(c *gin.Context) {
	u.GetAll("main_action", c)
	return
}