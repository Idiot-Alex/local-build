package test

import (
	"bytes"
	"encoding/json"
	. "local-build/internal/handler"
	"local-build/internal/pkg/env"
	"local-build/internal/store/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestToolList(t *testing.T) {
	// 创建一个 router
	router := initRouter()

	// 创建一个 GET 请求并设置参数
	req1, _ := http.NewRequest(http.MethodGet, "/api/tool/list", nil)

	// 为每个测试用例创建一个新的 ResultRecorder
	res1 := httptest.NewRecorder()

	// 向 router 发送 GET 请求
	router.ServeHTTP(res1, req1)

	log.Printf("res: %v\n", res1.Body.String())
	// 检查响应是否符合预期
	if res1.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res1.Body.String())
	}
}

func TestSaveTool(t *testing.T) {
	router := initRouter()
	// 创建一个 GET 请求并设置参数
	tool := model.Tool{Name: "test", Type: env.JDK, Path: "/Library/Java/JavaVirtualMachines/jdk-17.jdk/Contents/Home"}
	body, err := json.Marshal(tool)
	if err != nil {
		panic(err)
	}
	log.Printf("body: %+v\n", bytes.NewBuffer(body))
	req, _ := http.NewRequest(http.MethodPost, "/api/tool/save", bytes.NewBuffer(body))

	// 为每个测试用例创建一个新的 ResultRecorder
	res := httptest.NewRecorder()

	// 向 router 发送请求
	router.ServeHTTP(res, req)

	log.Printf("code: %v, res: %v\n", res.Code, res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func TestProjectList(t *testing.T) {
	router := initRouter()

	req, _ := http.NewRequest(http.MethodGet, "/api/project/list", nil)

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	log.Printf("code: %v, res: %v\n", res.Code, res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func TestDelProject(t *testing.T) {
	router := initRouter()

	data := model.ReqIds{ Ids: []string{"1"} }
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/api/project/del", bytes.NewBuffer(body))

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	log.Printf("code: %v, res: %v\n", res.Code, res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func initRouter() *gin.Engine {
	// 创建一个 router
	router := gin.Default()
	router.GET("/api/env/init", InitEnv)
	router.GET("/api/tool/list", ToolList)
	router.POST("/api/tool/save", SaveTool)

	router.POST("/api/project/save", SaveProject)
	router.GET("/api/project/list", ProjectList)
	router.POST("/api/project/del", DelProject)
	return router
}
