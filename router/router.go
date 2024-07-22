package router

import (
	"net/http"
	"taiwan-calendar/controller"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	limiter := tollbooth.NewLimiter(5, nil) // 1 秒最多 5 次請求
	limiter.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})
	limiter.SetMessageContentType("application/json; charset=utf-8")
	limiter.SetMessage(`{"http_code": "429", "message": "API 請求頻率過快，請稍後再試！", "status": "error"}`)

	r.GET("/taiwan-calendar/", tollbooth_gin.LimitHandler(limiter), controller.GetApiDoc)
	r.GET("/taiwan-calendar/docs", tollbooth_gin.LimitHandler(limiter), controller.GetApiDoc)
	r.GET("/taiwan-calendar/:year/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/:day/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"http_code": "404",
			"message":   "查無資料",
			"status":    "success",
		})
	})

	return r
}
