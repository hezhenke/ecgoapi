package services

import (
	"ecshopGoApi/dtos"
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
)

/**
 * 首页商品列表
 */

func GetHomeList(req dtos.HomeRequest)(hotProducts []models.Goods,recentlyProducts []models.Goods,goodProducts []models.Goods){

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