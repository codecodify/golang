package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"wechat/handlers"
	"wechat/initRouter"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetRouter()
}

func TestIndexHandler(t *testing.T){
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
}


func TestGetResource(t *testing.T){
	resourceItem := handlers.GetResource("gexingtx")
	assert.Equal(t, len(resourceItem) > 0, true)
}

func TestApiHandler(t *testing.T){
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api?s=katongtx", nil)
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)
}