package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	GoodsId	int32	`gorm:"primary_key" json:"id"` //
	CatId	int16	`json:"cat_id"` //
	GoodsSn	string	`json:"goods_sn"` //
	GoodsName	string	`json:"goods_name"` //
	GoodsNameStyle	string	`json:"goods_name_style"` //
	ClickCount	int	`json:"click_count"` //
	BrandId	int16	`json:"brand_id"` //
	ProviderName	string	`json:"provider_name"` //
	GoodsNumber	int32	`json:"goods_number"` //
	GoodsWeight	float32	`json:"goods_weight"` //
	MarketPrice	float32	`json:"market_price"` //
	VirtualSales	int16	`json:"virtual_sales"` //
	ShopPrice	float32	`json:"shop_price"` //
	PromotePrice	float32	`json:"promote_price"` //
	PromoteStartDate	int	`json:"promote_start_date"` //
	PromoteEndDate	int	`json:"promote_end_date"` //
	WarnNumber	int8	`json:"warn_number"` //
	Keywords	string	`json:"keywords"` //
	GoodsBrief	string	`json:"goods_brief"` //
	GoodsDesc	string	`json:"goods_desc"` //
	GoodsThumb	string	`json:"goods_thumb"` //
	GoodsImg	string	`json:"goods_img"` //
	OriginalImg	string	`json:"original_img"` //
	IsReal	int8	`json:"is_real"` //
	ExtensionCode	string	`json:"extension_code"` //
	IsOnSale	int8	`json:"is_on_sale"` //
	IsAloneSale	int8	`json:"is_alone_sale"` //
	IsShipping	int8	`json:"is_shipping"` //
	Integral	int	`json:"integral"` //
	AddTime	int	`json:"add_time"` //
	SortOrder	int16	`json:"sort_order"` //
	IsDelete	int8	`json:"is_delete"` //
	IsBest	int8	`json:"is_best"` //
	IsNew	int8	`json:"is_new"` //
	IsHot	int8	`json:"is_hot"` //
	IsPromote	int8	`json:"is_promote"` //
	BonusTypeId	int8	`json:"bonus_type_id"` //
	LastUpdate	int	`json:"last_update"` //
	GoodsType	int16	`json:"goods_type"` //
	SellerNote	string	`json:"seller_note"` //
	GiveIntegral	int	`json:"give_integral"` //
	RankIntegral	int	`json:"rank_integral"` //
	SuppliersId	int16	`json:"suppliers_id"` //
	IsCheck	int8	`json:"is_check"` //
}
var scope *gorm.Scope


//获取活动价
func (good *Goods) GetPromotePrice() float32  {
	if good.PromotePrice == 0.0{
		return 0.0
	}
	//获取当前时间戳
	now := time.Now().Unix()
	if now >= int64(good.PromoteStartDate) && now <= int64(good.PromoteEndDate) {
		return good.PromotePrice
	}
	return 0.0
}

/**
 * 1.如果有活动价优先返回活动价
 * 2.如果用户已经登录了查看当前用户的会员等级
 * 2.1 如果当前商品设置专门的会员等级价，返回该商品的等级价
 * 2.2 如果没有专门设置，价格=当前等级下的折扣 * 商品的售卖价格
 * 3. 如果没有登录或登录后没有命中会员等级，则返回商品的原始售卖价
 */

/*
func (*Goods)TableName() string {
	return "goods"
}
 */


