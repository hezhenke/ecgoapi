package models

import "time"

type Users struct {
	UserId	int	`gorm:"primary_key" json:"user_id"` //
	Email	string	`json:"email"` //
	UserName	string	`json:"user_name"` //
	Password	string	`json:"password"` //
	Question	string	`json:"question"` //
	Answer	string	`json:"answer"` //
	Sex	int	`json:"sex"` //
	Birthday	time.Time	`json:"birthday"` //
	UserMoney	float32	`json:"user_money"` //
	FrozenMoney	float32	`json:"frozen_money"` //
	PayPoints	int	`json:"pay_points"` //
	RankPoints	int	`json:"rank_points"` //
	AddressId	int	`json:"address_id"` //
	RegTime	int	`json:"reg_time"` //
	LastLogin	int	`json:"last_login"` //
	LastTime	time.Time	`json:"last_time"` //
	LastIp	string	`json:"last_ip"` //
	VisitCount	int	`json:"visit_count"` //
	UserRank	int	`json:"user_rank"` //
	IsSpecial	int	`json:"is_special"` //
	EcSalt	string	`json:"ec_salt"` //
	Salt	string	`json:"salt"` //
	ParentId	int	`json:"parent_id"` //
	Flag	int	`json:"flag"` //
	Alias	string	`json:"alias"` //
	Msn	string	`json:"msn"` //
	Qq	string	`json:"qq"` //
	OfficePhone	string	`json:"office_phone"` //
	HomePhone	string	`json:"home_phone"` //
	MobilePhone	string	`json:"mobile_phone"` //
	IsValidated	int	`json:"is_validated"` //
	CreditLine	float32	`json:"credit_line"` //
	PasswdQuestion	string	`json:"passwd_question"` //
	PasswdAnswer	string	`json:"passwd_answer"` //
}
