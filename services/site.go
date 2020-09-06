package services

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"os"
)

func GetSiteInfo()map[string]interface{}{

	database := infrastructure.GetDb()
	var configList []models.ShopConfig
	database.Model(&models.ShopConfig{}).Where("code in (?)",
		[]string{"shop_name","shop_title","shop_desc","shop_logo","shop_closed","service_phone"},
	).Find(&configList)
	configMap := map[string]interface{}{}
	for _,config := range configList{
		configMap[config.Code] = config.Value
	}
	return map[string]interface{}{
		"name":    configMap["shop_name"],
		"title":   configMap["shop_title"],
		"desc":    configMap["shop_desc"],
		"logo":    configMap["shop_logo"],
		"opening": configMap["shop_closed"] == 1,
		"telephone": configMap["service_phone"],
		"terms_url":os.Getenv("TERMS_URL"),
		"about_url":os.Getenv("ABOUT_URL"),
	}

}
