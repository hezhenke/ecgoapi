package dtos

import "ecshopGoApi/models"

type AuthRequest struct {
	Vendor int `json:"vendor" binding:"required"`
	DeviceId string `json:"device_id"`
	AccessToken string `json:"access_token"`
	JsCode string `json:"js_code"`
	OpenId string `json:"open_id"`
	InviteCode int `json:"invite_code"`
}


func CreateAuthDto(isNewUser bool,openid string,token string,user models.Users)map[string]interface{}{
	return map[string]interface{}{
		"is_new_user":isNewUser,
		"openid":openid,
		"token":token,
		"user":user,
	}
}
