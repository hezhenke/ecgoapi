package tests

import (
	"ecshopGoApi/infrastructure"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestMain(m *testing.M)  {
	err := godotenv.Load("../.env")
	if err != nil{
		log.Fatal(err)
	}
	database := infrastructure.OpenDbConnection()
	defer func() {
		if err := database.Close(); err!=nil{
			log.Fatal(err.Error())
		}
	}()
	m.Run()
}
