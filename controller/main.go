package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"taiwan-calendar/model"
	"taiwan-calendar/repository"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 定義錯誤回應的格式
type ErrorResponse struct {
	HttpCode int    `json:"http_code"`
	Message  string `json:"message""`
	Status   string `json:"status" example:"error"`
}

// GetCalendar doc
// @Summary 取得行事曆資料
// @Tags taiwan-calendar
// @Param year path string true "年份 (西元)"
// @Param month path string false "月份 (01-12)"
// @Param day path string false "日期 (01-31)"
// @Param isHoliday query bool false "是否為假日"
// @Success 200 {array} model.Calendar "回傳行事曆資料"
// @Failure 400 {object} ErrorResponse "格式錯誤"
// @Failure 404 {object} ErrorResponse "範圍錯誤或是查無資料"
// @Failure 429 {object} ErrorResponse "頂到頻率限制
// @Router /taiwan-calendar/{year} [get]
// @Router /taiwan-calendar/{year}/{month} [get]
// @Router /taiwan-calendar/{year}/{month}/{day} [get]
func GetCalendar(c *gin.Context) {
	yearParam := c.Param("year")
	monthParam := c.Param("month")
	dayParam := c.Param("day")
	isHolidayQuery := c.Query("isHoliday")

	// 檢查 year 是否為有效數字
	if _, yearErr := strconv.Atoi(yearParam); yearErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"http_code": 400,
			"message":   "年份格式錯誤，請輸入有效的西元年份",
			"status":    "error",
		})
		return
	}

	// 檢查 month 是否為有效數字
	month, monthErr := strconv.Atoi(monthParam)
	if monthParam != "" {
		if monthErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"http_code": 400,
				"message":   "月份格式錯誤，請輸入 01-12 的月份",
				"status":    "error",
			})
			return
		}
		if month < 1 || month > 12 {
			c.JSON(http.StatusNotFound, gin.H{
				"http_code": 404,
				"message":   "月份範圍錯誤，請輸入 01-12 的月份",
				"status":    "error",
			})
			return
		}
	}

	// 檢查 day 是否為有效數字
	day, dayErr := strconv.Atoi(dayParam)
	if dayParam != "" {
		if dayErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"http_code": 400,
				"message":   "日期格式錯誤，請輸入 01-31 的日期",
				"status":    "error",
			})
			return
		}
		if day < 1 || day > 31 {
			c.JSON(http.StatusNotFound, gin.H{
				"http_code": 404,
				"message":   "日期範圍錯誤，請輸入 01-31 的日期",
				"status":    "error",
			})
			return
		}
	}

	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	filename := "data/" + yearParam + ".json"

	calendar, err := repository.LoadCalendar(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"http_code": 404,
			"message":   "資料取得失敗，請參考 API 文件目前支援的年份",
			"status":    "error",
		})
		return
	}

	var filteredCalendar []model.Calendar

	// 月份過濾
	if month > 0 {
		for _, cal := range calendar {
			if strings.HasPrefix(cal.Date, yearParam+monthStr) {
				filteredCalendar = append(filteredCalendar, cal)
			}
		}
		calendar = filteredCalendar
		filteredCalendar = []model.Calendar{}
	}

	// 日期過濾
	if day > 0 {
		for _, cal := range calendar {
			if strings.HasPrefix(cal.Date, yearParam+monthStr+dayStr) {
				filteredCalendar = append(filteredCalendar, cal)
			}
		}
		calendar = filteredCalendar
		filteredCalendar = []model.Calendar{}
	}

	// 是否放假過濾
	if isHolidayQuery != "" {
		isHoliday := isHolidayQuery == "true"
		for _, cal := range calendar {
			if cal.IsHoliday == isHoliday {
				filteredCalendar = append(filteredCalendar, cal)
			}
		}
		calendar = filteredCalendar
	}

	if len(calendar) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"http_code": 404,
			"message":   "查無資料，請檢查過濾內容是否正確",
			"status":    "error",
		})
		return
	}

	c.JSON(http.StatusOK, calendar)
}
