package controllers

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context)  {


	var req dtos.HomeRequest
	if err := c.ShouldBindJSON(&req);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	fmt.Println(req.CategoryId)
	hotProducts,recentlyProducts,goodProducts := services.GetHomeList(req)
	c.JSON(200,dtos.CreateHomeProductListDto(hotProducts,recentlyProducts,goodProducts))
}