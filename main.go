package main

import (
	"fmt"
	"taiwan-calendar/docs"
	"taiwan-calendar/router"
)

// @title 台灣行事曆 API
// @version 1.0
// @description 台灣行事曆相關的 API 服務，可參考：https://github.com/880831ian/taiwan-calendar
// @host 127.0.0.1:80
// @BasePath /taiwan-calendar/
// @schemes http https
func main() {
	docs.SwaggerInfo.Title = "台灣行事曆 API"
	docs.SwaggerInfo.Description = "台灣行事曆相關的 API 服務，可參考：https://github.com/880831ian/taiwan-calendar"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:80"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := router.SetupRouter()

	port := 80
	r.Run(fmt.Sprintf(":%d", port))
}
