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

func TestBuildPlanList(t *testing.T) {
	router := initRouter()
	jsonData := []byte(`
		{
			"pageNo": 1,
			"pageSize": 30
		}
	`)
	req, _ := http.NewRequest(http.MethodGet, "/api/build-plan/list", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	log.Printf("res: %v\n", res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func TestSaveBuildPlan(t *testing.T) {
	router := initRouter()

	buildPlan := model.BuildPlan{
		Name: "test1",
	}
	jsonData, _ := json.Marshal(buildPlan)
	req, _ := http.NewRequest(http.MethodPost, "/api/build-plan/save", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	log.Printf("res: %v\n", res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func TestDelBuildPlan(t *testing.T) {
	router := initRouter()

	data := model.ReqIds{Ids: []string{"20240123160639060"}}
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/api/build-plan/del", bytes.NewBuffer(body))

	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	log.Printf("code: %v, res: %v\n", res.Code, res.Body.String())
	// 检查响应是否符合预期
	if res.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got %v", res.Body.String())
	}
}

func TestToolList(t *testing.T) {
	// 创建一个 router
	router := initRouter()

	jsonData := []byte(`
		{
			"pageNo": 1,
			"pageSize": 30
		}
	`)
	// 创建一个 GET 请求并设置参数
	req1, _ := http.NewRequest(http.MethodGet, "/api/tool/list", bytes.NewBuffer(jsonData))
	// 设置 Content-Type 头部
	req1.Header.Set("Content-Type", "application/json")

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

	jsonData := []byte(`
		{
			"pageNo": 1,
			"pageSize": 30
		}
	`)
	req, _ := http.NewRequest(http.MethodGet, "/api/project/list", bytes.NewBuffer(jsonData))

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

	data := model.ReqIds{Ids: []string{"1"}}
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

func TestSaveProject(t *testing.T) {
	router := initRouter()
	// 创建一个 GET 请求并设置参数
	p := model.Project{
		Name: "test",
		Path: "/tmp/test/hello",
		RepoConfig: model.RepoConfig{
			Type:       env.GIT,
			AccessType: env.CREDENTIALS,
			Url:        "https://gitee.com/hotstrip/hello-world-java.git",
		},
	}
	body, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	log.Printf("body: %+v\n", bytes.NewBuffer(body))
	req, _ := http.NewRequest(http.MethodPost, "/api/project/save", bytes.NewBuffer(body))

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

func initRouter() *gin.Engine {
	// 创建一个 router
	router := gin.Default()
	router.GET("/api/env/init", InitEnv)
	router.GET("/api/tool/list", ToolList)
	router.POST("/api/tool/save", SaveTool)

	router.POST("/api/project/save", SaveProject)
	router.GET("/api/project/list", ProjectList)
	router.POST("/api/project/del", DelProject)

	router.GET("/api/build-plan/list", BuildPlanList)
	router.POST("/api/build-plan/save", SaveBuildPlan)
	router.POST("/api/build-plan/del", DelBuildPlan)

	return router
}
