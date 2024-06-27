package router

import (
	"github.com/gin-gonic/gin"
	"taiwan-calendar/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/taiwan-calendar/:year/", controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/", controller.GetCalendar)

	return r
}
