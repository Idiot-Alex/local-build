package test

import (
	. "local-build/internal/handler"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test1(t *testing.T) {
	// 创建一个 router
	router := gin.Default()
	router.GET("/api/env/init", InitEnv)
	router.GET("/api/tool/list", ListTool)

	// 创建一个 GET 请求并设置参数
	req1, _ := http.NewRequest(http.MethodGet, "/api/tool/list", nil)
	// req2, _ := http.NewRequest(http.MethodGet, "/api/env/init", nil)

	// 为每个测试用例创建一个新的 ResultRecorder
	res1 := httptest.NewRecorder()
	// res2 := httptest.NewRecorder()

	// 向 router 发送 GET 请求
	router.ServeHTTP(res1, req1)
	// router.ServeHTTP(res2, req2)

	log.Printf("res: %v\n", res1.Body.String())
	// 检查响应是否符合预期
	if res1.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res1.Body.String())
	}
	
}