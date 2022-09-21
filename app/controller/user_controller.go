package controller

import (
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/auth"
	"gin-skeleton/pkg/database"
	"gin-skeleton/pkg/jwt"
	"gin-skeleton/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// Register 注册
func (user *UserController) Register(c *gin.Context) {
	// nickname
	nickname := c.PostForm("nickname")
	// email
	email := c.PostForm("email")
	// password
	password := c.PostForm("password")
	// 验证码
	captcha := c.PostForm("captcha")

	if nickname == "" || email == "" || password == "" || captcha == "" {
		response.ParamError(c)
		return
	}

	// 验证验证码

	var userModel models.User
	database.DB.Where(&models.User{Email: email}).First(&userModel)

	if userModel.ID > 0 {
		response.FailedCommon(c, "该邮箱已被注册,请登录")
		return
	}

	userModel = models.User{
		Nickname: nickname,
		Email:    email,
		Password: password,
	}

	userModel.Create()

	if userModel.ID > 0 {
		id := userModel.GetStringID()
		token := jwt.NewJWT().IssueToken(id, userModel.Nickname)
		response.Data(c, map[string]string{
			"id":    id,
			"token": token,
		})
		return
	}

	response.FailedCommon(c, "注册失败,请重试")
}

// Login 登录
func (user *UserController) Login(c *gin.Context) {
	// email
	email := c.PostForm("email")
	// password
	password := c.PostForm("password")

	if email == "" || password == "" {
		response.ParamError(c)
		return
	}

	var userModel models.User
	database.DB.Where(&models.User{Email: email}).First(&userModel)

	if userModel.ID > 0 {
		id := userModel.GetStringID()
		token := jwt.NewJWT().IssueToken(id, userModel.Nickname)
		response.Data(c, map[string]string{
			"id":    id,
			"token": token,
		})
		return
	}

	response.FailedCommon(c, "该邮箱未注册,请注册后登录")
}

// GetUserInfo 获取用户信息
func (user *UserController) GetUserInfo(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
