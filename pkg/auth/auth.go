package auth

import (
	"errors"
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/logger"
	"github.com/gin-gonic/gin"
)

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) models.User {
	userModel, ok := c.MustGet("current_user").(models.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return models.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
