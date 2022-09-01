package test

import (
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIndex(t *testing.T)  {
	response := Get("/index/index", nil)
	var responseData map[string]string
	_ = json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Hello World", responseData["data"])
}
