package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Success(c *gin.Context) {
	JSON(c, gin.H{
		"status": "success",
	})
}

func Data(c *gin.Context, data interface{}) {
	JSON(c, gin.H{
		"status": "success",
		"data":   data,
	})
}

func Failed(c *gin.Context, code int, message string) {
	JSON(c, gin.H{
		"status":  "error",
		"message": message,
		"code":    code,
	})
}

func FailedCommon(c *gin.Context, message string) {
	JSON(c, gin.H{
		"status":  "error",
		"message": message,
		"code":    0,
	})
}

func ParamError(c *gin.Context) {
	JSON(c, gin.H{
		"status":  "error",
		"message": "param error",
	})
}

// Unauthorized 响应 401，未传参 msg 时使用默认消息
// 登录失败、jwt 解析失败时调用
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": defaultMessage("unknown error", msg...),
	})
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
