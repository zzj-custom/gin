package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gookit/goutil/timex"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func ToIntSlice(str string, separation string) []int {
	splits := strings.Split(str, separation)
	var slice = make([]int, 0, len(splits))
	for _, split := range splits {
		atoi, err := strconv.Atoi(split)
		if err != nil {
			continue
		}
		slice = append(slice, atoi)
	}
	return slice
}

func ToStringSlice(str string, separation string) []string {
	splits := strings.Split(str, separation)
	return splits
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func ConvertEndTime(t time.Time) time.Time {
	return t.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
}

func Decimal(value float64) float64 {
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return res
}

func TimeFormatDay(time time.Time) string {
	return time.Format("2006-01-02")
}

func TimeFormatSecond(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func MarkToTime(str string) time.Time {
	var t time.Time
	switch str {
	case "today":
		t = time.Now()
	case "yesterday":
		t = timex.DayEnd(timex.AddDay(time.Now(), -1))
	case "current_week":
		// 获取当天是周几
		d := int(time.Now().Weekday())
		if d == 0 {
			d = 7
		}
		t = timex.DayEnd(timex.AddDay(time.Now(), -d))
	case "current_month":
		t = timex.DayEnd(timex.AddDay(time.Now(), -time.Now().Day()))
	case "last_week":
		d := int(time.Now().Weekday())
		if d == 0 {
			d = 14
		} else {
			d += 7
		}
		t = timex.DayEnd(timex.AddDay(time.Now(), -d))
	case "last_month":
		t = timex.DayEnd(time.Now().AddDate(0, -1, -time.Now().Day()))
	default:
		// 今年最后一天
		t = timex.DayEnd(timex.AddDay(time.Now(), -timex.Now().YearDay()))
	}
	return t
}

// TimeStringToTime 字符串时间转为time.Time
func TimeStringToTime(tm string) time.Time {
	var tps = []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006-01-02",
		"2006/01/02",
		"20060102",
		"2006",
		"15:04:05",
		"20060102",
		"2006-01-02",
		"2006-01-02 15",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"RFC3339",
		"RFC822Z",
		"RFC822",
		"2006/01/02",
		"2006/01/02 15",
		"2006/01/02 15:04",
		"2006/01/02 15:04:05",
		"2006.01.02",
		"2006.01.02 15",
		"2006.01.02 15:04",
		"2006.01.02 15:04:05",
	}
	for i := range tps {
		t, err := time.ParseInLocation(tps[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

// RunFuncName 获取正在运行的函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	arrName := strings.Split(f.Name(), ".")
	return arrName[len(arrName)-1]
}

// MomDate 获取同比时间
func MomDate(start time.Time, end time.Time) (time.Time, time.Time) {
	return start.AddDate(-1, 0, 0), end.AddDate(-1, 0, 0)
}
