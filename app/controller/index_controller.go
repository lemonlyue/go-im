package controller

import (
	"fmt"
	"gin-skeleton/app/models"
	"gin-skeleton/pkg/database"
	"gin-skeleton/pkg/response"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (index *IndexController) Index(c *gin.Context) {
	response.Data(c, "Hello World")
}

func (index *IndexController) Add(c *gin.Context)  {
	name := c.DefaultPostForm("name", "")
	if name == "" {
		response.ParamError(c)
		return
	}

	test := models.Test{Name: name}
	database.DB.Create(&test)
	fmt.Println(2)

	response.Data(c, name)
}