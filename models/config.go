package models

import "time"

type Config struct {
	ID	int	`gorm:"primary_key" json:"id"` //
	Name	string	`json:"name"` //
	Type	string	`json:"type"` //
	Description	string	`json:"description"` //
	Code	string	`json:"code"` //
	Config	string	`json:"config"` //
	Status	int8	`json:"status"` //
	CreatedAt	time.Time	`json:"created_at"` //
	UpdatedAt	time.Time	`json:"updated_at"` //
}

func (*Config)TableName() string {
	return "ecs_config"
}
