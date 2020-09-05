package controllers

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context){
	var req dtos.AuthRequest
	if err := c.ShouldBindJSON(&req);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authResult,err := services.Auth(req)
	if err !=nil{
		c.JSON(http.StatusOK,dtos.FormatError(400,err.Error()))
	}
	c.JSON(http.StatusOK,dtos.FormatBody(authResult))
}
