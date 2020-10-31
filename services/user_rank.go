package services

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"errors"
)

type UserRankVo struct {
	RankId int32
	Discount float32
}

//获取用户的会员等级ID和所享受的折扣
func  GetUserRankByUserId( userId int32) (userRankVo UserRankVo,err error) {
	var(
		db = infrastructure.GetDb()
		user = models.Users{}
	)

	err = db.Model(&models.Users{}).Where("user_id=?",userId).First(&user).Error
	if err != nil{
		return
	}
	userRank := models.UserRank{}
	if user.UserRank == 0 {
		err = db.Model(&models.UserRank{}).Where(
			"special_rank=0 AND min_points <=? AND max_points>= ?",
			user.RankPoints, user.RankPoints).First(&userRank).Error
	}else{
		err = db.Model(&models.UserRank{}).Where("rank_id=?",user.UserRank).First(&userRank).Error
	}
	if err != nil{
		return
	}
	userRankVo.RankId = userRank.RankId
	userRankVo.Discount = float32(userRank.Discount) * 0.01
	return userRankVo,nil
}

//获取会员价
func GetMemberRankPrice(goods *models.Goods,userId int32) (memberPrice float32,err error) {
	if goods == nil {
		return 0.0,errors.New("商品对象为nil")
	}
	userRankVo,err := GetUserRankByUserId(userId)
	if err != nil{
		return goods.ShopPrice,nil
	}
	if userRankVo.Discount > 0 {
		memberPrice = goods.ShopPrice * userRankVo.Discount
	}else{
		memberPrice = goods.ShopPrice
	}
	return memberPrice,nil
}
