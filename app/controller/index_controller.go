package controller

import (
	"fmt"
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseController
}

func (index *IndexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": "Hello World",
	})
}

func (index *IndexController) Add(c *gin.Context)  {
	name := c.DefaultPostForm("name", "")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "param error",
			"status": "error",
		})
		return
	}

	fmt.Println(1)
	test := models.Test{Name: name}
	database.DB.Create(&test)
	fmt.Println(2)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": name,
	})
}