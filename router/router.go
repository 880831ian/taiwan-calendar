package router

import (
	"encoding/json"
	"net/http"
	"os"
	"taiwan-calendar/controller"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

// loadBlockedIPs loads the blocked IPs from a JSON file
func loadBlockedIPs(filepath string) map[string]struct{} {
	file, err := os.Open(filepath)
	if err != nil {
		return make(map[string]struct{})
	}
	defer file.Close()

	var blockedIPs []string
	if err := json.NewDecoder(file).Decode(&blockedIPs); err != nil {
		return make(map[string]struct{})
	}

	ipMap := make(map[string]struct{})
	for _, ip := range blockedIPs {
		ipMap[ip] = struct{}{}
	}
	return ipMap
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	blockedIPs := loadBlockedIPs("blocked_ips.json")

	r.Use(func(c *gin.Context) {
		clientIP := c.ClientIP()
		if _, exists := blockedIPs[clientIP]; exists {
			c.JSON(http.StatusForbidden, gin.H{
				"http_code": "403",
				"message":   "系統偵測到使用量過大，已先將該 IP 進行封鎖，如還需使用，請聯絡開發人員 https://t.me/pinyichuchu，若超過 3 天仍未聯絡，會直接從網路層永久封鎖該 IP",
				"status":    "error",
			})
			c.Abort()
			return
		}
		c.Next()
	})

	limiter := tollbooth.NewLimiter(2, nil) // 1 秒最多 2 次請求
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
