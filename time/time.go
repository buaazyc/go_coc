package time

import (
	"fmt"
	"log"
	"time"
)

// TimeToDay time.Time格式转化为日期
func TimeToDay(utcTimeStr string) (string, error) {
	time, err := GetTimeFromCoc(utcTimeStr)
	if err != nil {
		return "", err
	}
	timeLayout := "20060102"
	return time.Format(timeLayout), nil
}

// TimeToMonth time.Time格式转化为月份
func TimeToMonth(utcTimeStr string) (string, error) {
	time, err := GetTimeFromCoc(utcTimeStr)
	if err != nil {
		return "", err
	}
	timeLayout := "200601"
	return time.Format(timeLayout), nil
}

// GetTimeFromCoc 将coc的时间字符串转换为本地时区的time.Time格式
func GetTimeFromCoc(timeStr string) (time.Time, error) {
	timeLayout := "20060102T150405.000Z"
	utcTime, err := time.Parse(timeLayout, timeStr)
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
