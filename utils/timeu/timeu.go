package timeu

import (
	"strconv"
	"time"
)

// TrimHour 对齐到小时
func TrimHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// TrimDate 对其到天，舍去时分秒
func TrimDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDate 格式化日期区间
func FormatDatePeriod(b, e time.Time) string {
	return b.Format("2006/01/02") + " - " + e.Format("2006/01/02")
}

// FormatTsDate 格式化日期
func FormatTsDate(ts int64) string {
	return time.Unix(ts, 0).Format("20060102")
}

func FormatDateFromStr(date *string) (*string, error) {
	if date == nil || *date == "" {
		return nil, nil
	}

	i, err := strconv.ParseInt(*date, 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(i, 0)
	formatTime := tm.Format("20060102")
	return &formatTime, nil
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 获取本周周一的时间
func GetFirstDateOfWeek(d time.Time) time.Time {
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()).AddDate(0, 0, offset)
}
