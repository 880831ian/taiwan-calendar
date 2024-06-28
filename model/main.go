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
	Date         string `json:"date"`
	DateFormat   string `json:"date_format"`
	Year         string `json:"year"`
	ROCYear      string `json:"roc_year"`
	Month        string `json:"month"`
	Day          string `json:"day"`
	Week         string `json:"week"`
	Week_Abbr    string `json:"week_abbr"`
	Week_Chinese string `json:"week_chinese"`
	IsHoliday    bool   `json:"isHoliday"`
	Caption      string `json:"caption"`
}
