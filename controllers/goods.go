package controllers

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context)  {

	userId := c.GetInt("currentUserId")
	var req dtos.HomeRequest
	if err := c.ShouldBindJSON(&req);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotProducts,recentlyProducts,goodProducts := services.GetHomeList(req,userId)
	c.JSON(200,dtos.CreateHomeProductListDto(hotProducts,recentlyProducts,goodProducts))
}