package test

import (
	"gin-skeleton/bootstrap"
	bootstrapConfig "gin-skeleton/config"
	"gin-skeleton/pkg/app"
	"gin-skeleton/pkg/config"
	"gin-skeleton/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
)

var router *gin.Engine

func init() {
	router = bootstrap.SetupRoute()
	app.IsUnitTestEnv = true

	bootstrapConfig.Initialize()
	config.InitConfig("")

	// init logger
	bootstrap.SetupLogger()

	//init db
	bootstrap.SetupDB()

	// init redis
	bootstrap.SetupRedis()
}

func Get(url string, data map[string]string) *httptest.ResponseRecorder {
	url += utils.ParseToStr(data)
	request := httptest.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	return response
}

func Post(url string, data map[string]string) *httptest.ResponseRecorder {
	params := utils.ParseToStr(data)
	params = strings.TrimLeft(params, "?")
	request := httptest.NewRequest("POST", url, strings.NewReader(params))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	return response
}
