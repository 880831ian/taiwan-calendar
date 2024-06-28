package router

import (
	"taiwan-calendar/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/taiwan-calendar/", controller.GetApiDoc)
	r.GET("/taiwan-calendar/docs", controller.GetApiDoc)
	r.GET("/taiwan-calendar/:year/", controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/", controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/:day/", controller.GetCalendar)

	return r
}
