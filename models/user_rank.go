package models

type UserRank struct {
	RankId int32 `gorm:primary_key`
	RankName string
	MinPoints int
	MaxPoints int
	Discount int
	ShowPrice int8
	SpecialRank int8
}
