package tests

import (
	"ecshopGoApi/services"
	"fmt"
	"testing"
)

func TestGetBannerList(t *testing.T){
	bannerList,err := services.GetBannerList()
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(bannerList)

}
