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
		/**
		 * 1.如果有活动价优先返回活动价
		 * 2.如果用户已经登录了查看当前用户的会员等级
		 * 2.1 如果当前商品设置专门的会员等级价，返回该商品的等级价
		 * 2.2 如果没有专门设置，价格=当前等级下的折扣 * 商品的售卖价格
		 * 3. 如果没有登录或登录后没有命中会员等级，则返回商品的原始售卖价
		 */
		price := goods.GetPromotePrice()
		if price <= 0.0{

		}else{

		}
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



