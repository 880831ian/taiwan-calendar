package main

import (
	"fmt"
	"taiwan-calendar/router"
)

func main() {
	r := router.SetupRouter()

	port := 80
	r.Run(fmt.Sprintf(":%d", port))
}
