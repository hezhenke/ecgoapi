package models

type ShopConfig struct {
	ID	int8	`gorm:"primary_key" json:"id"` //
	ParentId int8 `json:"parent_id"`
	Code string `json:"code"`
	Value string  `json:value`
}

func (*ShopConfig)TableName() string {
	return "shop_config"
}
