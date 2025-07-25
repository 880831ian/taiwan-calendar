package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"taiwan-calendar/model"
)

func GetSupportedYears() ([]string, error) {
	dataDir := "data"

	// 讀取 data 資料夾中的所有檔案
	files, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("無法讀取 data 資料夾: %v", err)
	}

	var years []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			// 從檔案名稱中提取年份（檔案名稱格式為 YYYY.json）
			fileName := file.Name()
			yearStr := strings.TrimSuffix(fileName, ".json")

			// 驗證是否為有效的年份
			if year, err := strconv.Atoi(yearStr); err == nil && year >= 1912 && year <= time.Now().Year() {
				years = append(years, yearStr)
			}
		}
	}

	// 排序年份
	sort.Strings(years)

	if len(years) == 0 {
		return nil, fmt.Errorf("在 data 資料夾中找不到任何有效的年份檔案")
	}

	return years, nil
}

func LoadCalendar(filename string) ([]model.Calendar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("資料取得失敗，請參考 API 文件目前支援的年份")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("資料讀取失敗，請聯繫開發人員")
	}

	var originalCalendar []model.OriginalCalendar
	err = json.Unmarshal(bytes, &originalCalendar)
	if err != nil {
		return nil, fmt.Errorf("無法解析 JOSN 檔案，請聯繫開發人員")
	}

	type MonthInfo struct {
		FullName     string
		Abbreviation string
	}

	monthMap := map[string]MonthInfo{
		"01": {FullName: "January", Abbreviation: "Jan"},
		"02": {FullName: "February", Abbreviation: "Feb"},
		"03": {FullName: "March", Abbreviation: "Mar"},
		"04": {FullName: "April", Abbreviation: "Apr"},
		"05": {FullName: "May", Abbreviation: "May"},
		"06": {FullName: "June", Abbreviation: "Jun"},
		"07": {FullName: "July", Abbreviation: "Jul"},
		"08": {FullName: "August", Abbreviation: "Aug"},
		"09": {FullName: "September", Abbreviation: "Sep"},
		"10": {FullName: "October", Abbreviation: "Oct"},
		"11": {FullName: "November", Abbreviation: "Nov"},
		"12": {FullName: "December", Abbreviation: "Dec"},
	}

	var calendar []model.Calendar
	for _, originalCalendar := range originalCalendar {
		isHoliday := originalCalendar.IsHoliday == "2"
		parsedDate, err := time.Parse("20060102", originalCalendar.Date)
		if err != nil {
			return nil, fmt.Errorf("日期解析錯誤，請聯繫開發人員")
		}

		dateFormat := parsedDate.Format("2006/01/02")
		year := parsedDate.Format("2006")
		yearInt, _ := strconv.Atoi(parsedDate.Format("2006"))
		roc_yearStr := strconv.Itoa(yearInt - 1911)
		month := parsedDate.Format("01")
		day := parsedDate.Format("02")

		monthInfo := monthMap[month]

		type WeekInfo struct {
			FullName     string
			Abbreviation string
		}

		weekMap := map[string]WeekInfo{
			"日": {FullName: "Sunday", Abbreviation: "Sun"},
			"一": {FullName: "Monday", Abbreviation: "Mon"},
			"二": {FullName: "Tuesday", Abbreviation: "Tue"},
			"三": {FullName: "Wednesday", Abbreviation: "Wed"},
			"四": {FullName: "Thursday", Abbreviation: "Thu"},
			"五": {FullName: "Friday", Abbreviation: "Fri"},
			"六": {FullName: "Saturday", Abbreviation: "Sat"},
		}

		weekInfo := weekMap[originalCalendar.Week]

		calendar = append(calendar, model.Calendar{
			Date:          originalCalendar.Date,
			DateFormat:    dateFormat,
			Year:          year,
			ROCYear:       roc_yearStr,
			Month:         month,
			Month_En:      monthInfo.FullName,
			Month_En_Abbr: monthInfo.Abbreviation,
			Day:           day,
			Week:          weekInfo.FullName,
			Week_Abbr:     weekInfo.Abbreviation,
			Week_Chinese:  originalCalendar.Week,
			IsHoliday:     isHoliday,
			Caption:       originalCalendar.Remark,
		})
	}

	return calendar, nil
}
