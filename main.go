package main

import (
	"fmt"
	"taiwan-calendar/router"
)

func main() {
	r := router.SetupRouter()

	port := 80
	fmt.Printf("啟動服務 Port 為 %d\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
