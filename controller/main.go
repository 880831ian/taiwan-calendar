package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"taiwan-calendar/model"
	"taiwan-calendar/repository"
)

func GetCalendar(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	filename := "data/" + year + ".json"

	calendar, err := repository.LoadCalendar(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 过滤月份
	if month != "" {
		var filteredCalendar []model.Calendar
		for _, calendar := range calendar {
			if strings.HasPrefix(calendar.Date, year+month) {
				filteredCalendar = append(filteredCalendar, calendar)
			}
		}
		calendar = filteredCalendar
	}

	// 过滤是否放假
	isHolidayQuery := c.Query("isHoliday")
	if isHolidayQuery != "" {
		var filteredCalendar []model.Calendar
		isHoliday := isHolidayQuery == "true"
		for _, calendar := range calendar {
			if calendar.IsHoliday == isHoliday {
				filteredCalendar = append(filteredCalendar, calendar)
			}
		}
		calendar = filteredCalendar
	}

	c.JSON(http.StatusOK, calendar)
}
