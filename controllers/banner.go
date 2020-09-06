package controllers

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBannerList(c *gin.Context)  {
	bannerList,err := services.GetBannerList()
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,dtos.FormatBody(gin.H{
		"banners":bannerList,
	}))
}
