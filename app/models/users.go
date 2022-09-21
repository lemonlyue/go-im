package models

import "gin-skeleton/pkg/database"

type User struct {
	BaseModel

	Nickname  string `gorm:"type:varchar(255);not null;index"`
	AvatarUrl string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password  string `gorm:"type:varchar(255);not null"`

	CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(&user)
}
