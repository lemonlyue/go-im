package users

import (
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/database"
)

func Get(userId string) (userModel models.User) {
	database.DB.Where("id", userId).First(&userModel)
	return
}
