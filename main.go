package main

import (
	"ecshopGoApi/controllers"
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

func main()  {
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("DB_NAME"))

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	count := 0
	//goods2 :=models.Goods{}
	database.Model(&models.Goods{}).Count(&count)

	var goods models.Goods

	database.Where("?=?","goods_id",1).Find(&goods)
	jsonStr,_ := json.Marshal(goods)
	fmt.Println(string(jsonStr))
	if isMatch,_ := regexp.MatchString("^https","http://www.baidu.com");isMatch{
		fmt.Println(isMatch)
	}
	r := gin.Default()

	v2 := r.Group("/v2")
	{
		v2.POST("/ecapi.home.product.list",controllers.Home)
		v2.POST("/ecapi.auth.social",controllers.Auth)
	}
	r.Run() // listen and serve on 0.0.0.0:8080

}