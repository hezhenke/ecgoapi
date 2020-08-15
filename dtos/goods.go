package dtos

import (
	"ecshopGoApi/models"
	"os"
	"strings"
)

type HomeRequest struct {
	CategoryId int `json:"category_id" binding:"omitempty,min=4,numeric"`
}

func createProductListDto(goodList []models.Goods) []interface{}{

	t := make([]interface{},len(goodList))
	for i,goods := range goodList{

		t[i] = map[string]interface{}{
			"id": goods.GoodsId,
			"name": goods.GoodsName,
			"price":goods.MarketPrice,
			"current_price":goods.MarketPrice,
			"default_photo":formatPhoto(goods.GoodsImg,"",""),
		}

	}
	return t
}


func CreateHomeProductListDto(hotProducts []models.Goods,recentlyProducts []models.Goods,goodProducts []models.Goods) map[string]interface{} {
	return map[string]interface{}{
		"error_code":0,
		"good_products":createProductListDto(goodProducts),
		"hot_products":createProductListDto(hotProducts),
		"recently_products":createProductListDto(recentlyProducts),
	}
}

func formatPhoto(img string,thumb string,domain string) map[string]interface{} {
	if len(img) == 0{
		return map[string]interface{}{}
	}
	if len(thumb) == 0{
		thumb = img
	}
	if len(domain) == 0{
		domain = os.Getenv("SHOP_URL")
	}
	if !strings.HasPrefix(thumb,"http") || !strings.HasPrefix(thumb,"https"){
		thumb = domain+"/"+thumb
	}
	if !strings.HasPrefix(img,"http") || !strings.HasPrefix(img,"https"){
		img = domain+"/"+img
	}
	return map[string]interface{}{
		"width": nil,
		"height": nil,
		"thumb": thumb,
		"large": img,
	}
}

