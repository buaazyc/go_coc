package time

import (
	"fmt"
	"log"
	"time"
)

var (
	simpleDateLayout  = "20060102"
	simpleMonthLayout = "200601"
	cocTimeLayout     = "20060102T150405.000Z"
)

// CurMonth
func CurMonth() string {
	return time.Now().Format(simpleMonthLayout)
}

// TimeToDay time.Time格式转化为日期
func TimeToDay(utcTimeStr string) (string, error) {
	time, err := getTimeFromCoc(utcTimeStr)
	if err != nil {
		return "", err
	}
	return time.Format(simpleDateLayout), nil
}

// TimeToMonth time.Time格式转化为月份
func TimeToMonth(utcTimeStr string) (string, error) {
	time, err := getTimeFromCoc(utcTimeStr)
	if err != nil {
		return "", err
	}
	return time.Format(simpleMonthLayout), nil
}

// getTimeFromCoc 将coc的时间字符串转换为本地时区的time.Time格式
func getTimeFromCoc(timeStr string) (time.Time, error) {
	utcTime, err := time.Parse(cocTimeLayout, timeStr)
	if err != nil {
		log.Printf("%v", err)
		return time.Now(), err
	}
	// 获取当地时区
	local, err := time.LoadLocation("Local")
	if err != nil {
		log.Printf("%v", err)
		return time.Now(), err
	}
	return utcTime.In(local), nil
}

// SeasonStr 将赛季的年月“2021-10”以汉字形式输出“2021年10月”
func SeasonStr(s string) string {
	return fmt.Sprintf("%v年%v月", s[:4], s[5:])
}

// SeasonToStr 将赛季的年月“202110”以汉字形式输出“2021年10月”
func SeasonToStr(s string) string {
	return fmt.Sprintf("%v年%v月", s[:4], s[4:])
}

// GetCurSeason 获取当前年月，如202110
func GetCurSeason() string {
	return time.Now().Format(simpleMonthLayout)
}
