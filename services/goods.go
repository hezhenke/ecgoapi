package services

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
)

/**
 * 首页商品列表
 */

func GetHomeList(req dtos.HomeRequest,userId int)(hotProducts []models.Goods,recentlyProducts []models.Goods,goodProducts []models.Goods){

	database := infrastructure.GetDb()
	mapCondition := make(map[string]interface{})
	mapCondition["is_delete"] = 0
	mapCondition["is_on_sale"] = 1
	mapCondition["is_alone_sale"] = 1
	if req.CategoryId > 0 {
		mapCondition["cat_id"] = req.CategoryId
	}

	//获取热销
	database.Model(&models.Goods{}).Where(mapCondition).Where("is_hot=?",1).Find(&hotProducts)

	//获取新品
	database.Model(&models.Goods{}).Where(mapCondition).Where("is_new=?",1).Find(&recentlyProducts)

	//获取精品
	database.Model(&models.Goods{}).Where(mapCondition).Where("is_best=?",1).Order("sort_order").Order("last_update desc").Find(&goodProducts)

	return hotProducts,recentlyProducts,goodProducts
}

// 获取单个商品的价格
func getCurrentPrice(goods * models.Goods,userId int32) float32{
	if goods == nil {
		return 0.0
	}
	promotionPrice := goods.GetPromotePrice()
	if promotionPrice > 0 {
		return promotionPrice
	}
	memberPrice,err := GetMemberRankPrice(goods,userId)
	if memberPrice > 0 && err == nil {
		return memberPrice
	}else{
		return goods.ShopPrice
	}
}
