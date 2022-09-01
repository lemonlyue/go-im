package test

import (
	"gin-skeleton/bootstrap"
	"gin-skeleton/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

var router *gin.Engine

func init()  {
	router = bootstrap.SetupRoute()
}

func Get(url string, data map[string]string) *httptest.ResponseRecorder {
	url += utils.ParseToStr(data)
	request := httptest.NewRequest("GET", url, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	return response
}

func Post(url string, data map[string]string) *httptest.ResponseRecorder {
	request := httptest.NewRequest("POST", url + utils.ParseToStr(data), nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)
	return response
}