package model

// 原始行事曆資料結構
type OriginalCalendar struct {
	Date      string `json:"西元日期"`
	Week      string `json:"星期"`
	IsHoliday string `json:"是否放假"`
	Remark    string `json:"備註"`
}

// 轉換後的行事曆資料結構
type Calendar struct {
	Date      string `json:"date"`
	Week      string `json:"week"`
	IsHoliday bool   `json:"isHoliday"`
	Remark    string `json:"remark"`
}
