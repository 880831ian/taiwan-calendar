package main

import (
	"fmt"
	"os"
	"taiwan-calendar/docs"
	"taiwan-calendar/router"
)

// @title 台灣行事曆 API
// @version 1.0
// @description 台灣行事曆相關的 API 服務，可參考：https://github.com/880831ian/taiwan-calendar
// @host 127.0.0.1
// @BasePath /
// @schemes http https
func main() {
	docs.SwaggerInfo.Title = "台灣行事曆 API"
	docs.SwaggerInfo.Description = "台灣行事曆相關的 API 服務，可參考：https://github.com/880831ian/taiwan-calendar"
	docs.SwaggerInfo.Version = "1.0"

	// 根據環境變數設定 host
	if os.Getenv("ENV") == "production" {
		docs.SwaggerInfo.Host = "api.pin-yi.me"
		docs.SwaggerInfo.BasePath = "/taiwan-calendar"
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Host = "127.0.0.1"
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	r := router.SetupRouter()

	port := 80
	r.Run(fmt.Sprintf(":%d", port))
}
