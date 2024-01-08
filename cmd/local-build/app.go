package localbuild

import (
	"local-build/internal/api"
	"local-build/internal/config"
	"local-build/internal/store/sqlite"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// run server
func RunServer() {
	log.Println("app starting...")

	config.Load("config.toml")

	startSqlite()

	startService()
}

// start sqlite server
func startSqlite() {
	sqlite.InitTable()
	log.Println("init table success...")
}

// bind web router server
func startService() {
	log.Println("server start success...")
	// 1.创建路由
	r := gin.Default()

	// 进行跨域设置
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许所有源
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	
	// 2.绑定路由规则，执行的函数
	api.InitApis(r)
	// 3.监听端口，默认在8080
	r.Run(":8000")
}