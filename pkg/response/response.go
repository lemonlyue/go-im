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
