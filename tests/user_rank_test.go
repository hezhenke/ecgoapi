package tests

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"fmt"
	"testing"
)

func TestGetUserRank(t *testing.T)  {
	database := infrastructure.GetDb()
	userRank := models.UserRank{}
	database.Model(&models.UserRank{}).First(&userRank)
	fmt.Println(userRank)
}
