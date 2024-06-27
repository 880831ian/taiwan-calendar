package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"taiwan-calendar/model"
)

func LoadCalendar(filename string) ([]model.Calendar, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("資料取得失敗，目前尚未更新檔案")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("資料讀取失敗，請聯繫開發人員")
	}

	var originalCalendar []model.OriginalCalendar
	err = json.Unmarshal(bytes, &originalCalendar)
	if err != nil {
		return nil, fmt.Errorf("無法解析 JOSN 檔案，請聯繫開發人員")
	}

	var calendar []model.Calendar
	for _, originalCalendar := range originalCalendar {
		isHoliday := originalCalendar.IsHoliday == "2"
		calendar = append(calendar, model.Calendar{
			Date:      originalCalendar.Date,
			Week:      originalCalendar.Week,
			IsHoliday: isHoliday,
			Remark:    originalCalendar.Remark,
		})
	}

	return calendar, nil
}
