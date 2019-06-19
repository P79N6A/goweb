package test

import (
	"gopkg.in/go-playground/assert.v1"
	router2 "goweb/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	router := router2.InitRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil) //测试接口请求
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "It works", w.Body.String())
}
