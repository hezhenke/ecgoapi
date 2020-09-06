package dtos

import (
	"ecshopGoApi/models"
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
			"default_photo":FormatPhoto(goods.GoodsImg,"",""),
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



