package services

import (
	"ecshopGoApi/dtos"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Banner struct {
	ID int `json:"id"`
	Photo map[string]interface{} `json:"photo"`
	Link string `json:"link"`
	Title string `json:"title"`
	Sort string `json:"sort"`
}

func GetBannerList() (bannerList []Banner,err error) {

	bannerList = []Banner{}
	client := &http.Client{}
	url := fmt.Sprintf("%v/data/flash_data.xml",os.Getenv("SHOP_URL"))
	request,err := http.NewRequest("GET",url,nil)
	if err !=nil{
		return
	}
	response,err := client.Do(request)
	if err !=nil{
		return
	}
	body,err := ioutil.ReadAll(response.Body)

	r,err := regexp.Compile(`item_url="([^"]+)"\slink="([^"]*)"\stext="([^"]*)"\ssort="([^"]*)"`)
	if err !=nil{
		return
	}
	matchList := r.FindAllStringSubmatch(string(body),-1)
	if len(matchList) != 0{
		for key,match := range matchList{
			banner := Banner{
				ID:key,
				Photo: dtos.FormatPhoto(match[1],"",""),
				Link: match[2],
				Title: match[3],
				Sort: match[4],
			}
			bannerList = append(bannerList,banner)
		}
	}
	return bannerList,nil
}
