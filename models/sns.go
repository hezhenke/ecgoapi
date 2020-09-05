package models

import "time"

type Sns struct {
	UserId	int	`gorm:"primary_key" json:"user_id"` //
	OpenId	string	`json:"open_id"` //
	Vendor	int8	`json:"vendor"` //第三方平台类型
	CreatedAt	time.Time	`json:"created_at"` //
	UpdatedAt	time.Time	`json:"updated_at"` //
}
