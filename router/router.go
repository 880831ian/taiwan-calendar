package router

import (
	"encoding/json"
	"net/http"
	"os"
	"taiwan-calendar/controller"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// IP 封鎖檢查
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

	// 請求限制
	limiter := tollbooth.NewLimiter(2, nil) // 1 秒最多 2 次請求
	limiter.SetIPLookups([]string{"X-Forwarded-For", "X-Real-IP", "RemoteAddr"})
	limiter.SetMessageContentType("application/json; charset=utf-8")
	limiter.SetMessage(`{"http_code": 429, "message": "API 請求頻率過快，請稍後再試！", "status": "error"}`)

	// Swagger 文檔路由
	r.GET("/taiwan-calendar/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/taiwan-calendar/:year/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)
	r.GET("/taiwan-calendar/:year/:month/:day/", tollbooth_gin.LimitHandler(limiter), controller.GetCalendar)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"http_code": 404,
			"message":   "路徑錯誤，API 文件請參考 https://api.pin-yi.me/taiwan-calendar/swagger/index.html",
			"status":    "error",
		})
	})

	return r
}
