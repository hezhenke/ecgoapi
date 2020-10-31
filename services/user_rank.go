package services

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
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
