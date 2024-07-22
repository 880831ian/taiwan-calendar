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

func GetApiDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"http_code": "200",
		"message":   "請參考 API 文件：https://documenter.getpostman.com/view/7653426/2sA3dsnZcb",
	})
}

func GetCalendar(c *gin.Context) {
	yearParam := c.Param("year")
	monthParam := c.Param("month")
	dayParam := c.Param("day")
	isHolidayQuery := c.Query("isHoliday")

	// 檢查 year, month, day 是否為有效數字
	_, yearErr := strconv.Atoi(yearParam)
	month, monthErr := strconv.Atoi(monthParam)
	day, dayErr := strconv.Atoi(dayParam)

	// 如果 year, month 或 day 輸入無效，返回 404
	if yearErr != nil || (monthParam != "" && monthErr != nil) || (dayParam != "" && dayErr != nil) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":    "error",
			"http_code": "404",
			"message":   "查無資料",
		})
		return
	}

	// 檢查 month 是否在 1 到 12 之間
	if month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "error",
			"http_code": "400",
			"message":   "輸入格式錯誤，月份必須在 1 到 12 之間",
		})
		return
	}

	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	filename := "data/" + yearParam + ".json"

	calendar, err := repository.LoadCalendar(filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":    "error",
			"http_code": "400",
			"message":   err.Error()})
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
		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"http_code": "404",
			"message":   "查無資料",
		})
		return
	}

	c.JSON(http.StatusOK, calendar)
}
