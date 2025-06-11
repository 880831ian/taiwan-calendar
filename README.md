# Taiwan-Calendar 臺灣行事曆

## 說明

最近在做一個專案，會需要台灣的國定假日，可惜政府資料開放平台沒有提供 API，只能下載檔案 (CSV、XML、JSON)，也沒有我想要的格式，因此我將政府資料開放平台提供的 JSON 資料格式透過 Go 寫了一隻簡單的 API，方便大家使用。

<br>

## 資料來源

本資料取自於[政府資料開放平台](https://data.gov.tw/dataset/14718)之「中華民國政府行政機關辦公日曆表」，欄位內容包含日期、星期、是否放假、說明。

<br>

## 使用說明

1. 由於是開源專案，歡迎大家直接使用，或是 Fork 進行修改。
2. 本人免費提供 API，但是不保證服務穩定性，API：[https://api.pin-yi.me/taiwan-calendar](https://api.pin-yi.me/taiwan-calendar)。
3. API 有每秒 2 次請求的限制，請大家不要把我的 API 打爆 QQ。
4. 使用說明：
   - `GET` `/taiwan-calendar/swagger/index.html/#`：查看 API 文件。
   - `GET` `/taiwan-calendar/{year}`：取得指定年份的行事曆資料。
   - `GET` `/taiwan-calendar/{year}/{month}`：取得指定年份+月份的行事曆資料。
   - `GET` `/taiwan-calendar/{year}/{month}/{day}`：取得指定年份+月份+日的行事曆資料。
   - `GET` `/taiwan-calendar/{year}/?isHoliday={true/false}`：取得指定年份是否為假日的行事曆資料。
   - `GET` `/taiwan-calendar/{year}/{month}/?isHoliday={true/false}`：取得指定年份+月份是否為假日的行事曆資料。
   - `GET` `/taiwan-calendar/{year}/{month}/{day}/?isHoliday={true/false}`：取得指定年份+月份+日是否為假日的行事曆資料。

    其餘詳細說明請參考 [API 文件](https://api.pin-yi.me/taiwan-calendar/swagger/index.html/#)。

<br>

## 顯示資料格式

```json
[
  {
    "date": "20250131", // 年月日
    "date_format": "2025/01/31", // 年/月/日
    "year": "2025", // 西元年
    "roc_year": "114", // 民國年
    "month": "01", // 月
    "month_en": "January", // 月 (英文)
    "month_en_abbr": "Jan", // 月縮寫 (英文)
    "day": "31", // 日
    "week": "Friday", // 星期
    "week_abbr": "Fri", // 星期縮寫
    "week_chinese": "五", // 星期 (中文)
    "isHoliday": true, // 是否為假日
    "caption": "春節" // 說明
  }
]
```

<br>

## 資料更新

由於中華民國政府行政機關辦公日曆表更新日期約為每年 6 月底更新下一年度的行事曆，因此本專案也會盡量於每年 6 月底更新下一年度的行事曆。

**目前收納資料範圍為 2017 年至 2025 年，未來會持續更新。**

(資料更新：2025/05/09 18:22)

<br>

## 資料內容備註

由於政府資料開放平台提供資料，有些無法下載 JSON 格式，因此我先下載 CSV 格式，再轉換成 JSON 格式。

以下為有問題的資料：

- 2025 年：沒有 JSON 格式，下載 CSV 會出現亂碼
- 2022 年：沒有 JSON 格式
- 2020 年：沒有 JSON 格式，下載 CSV 會出現亂碼
- 2019 年：沒有 JSON 格式，下載 CSV 會出現亂碼

<br>

## 補充

過程中有發現 GitHub 有人已經做了類似的專案，[TaiwanCalendar](https://github.com/ruyut/TaiwanCalendar/tree/master)，有參考部分 README.md 內容。
