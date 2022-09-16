package models

import "gin-skeleton/pkg/database"

type User struct {
	BaseModel

	Nickname string `gorm:"type:varchar(255);not null;index"`
	Email    string `gorm:"type:varchar(255);index;default:null"`
	Password string `gorm:"type:varchar(255)"`

	CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(&user)
}
