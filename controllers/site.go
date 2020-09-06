package controllers

import (
	"ecshopGoApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSiteInfo(c *gin.Context)  {
	siteInfo := services.GetSiteInfo()
	c.JSON(http.StatusOK,gin.H{
		"site_info" : siteInfo,
	})
}
