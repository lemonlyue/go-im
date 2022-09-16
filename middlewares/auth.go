// Package middlewares Gin 中间件
package middlewares

import (
	"gin-skeleton/app/service/users"
	"gin-skeleton/pkg/jwt"
	"gin-skeleton/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, "请登录后重试")
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := users.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "用户不存在")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Nickname)
		c.Set("current_user", userModel)

		c.Next()
	}
}
