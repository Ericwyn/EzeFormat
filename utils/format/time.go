package format

import (
	"errors"
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/GoTools/date"
	"github.com/Ericwyn/GoTools/str"
	"strconv"
	"time"
)

// ------------------------- 日期格式化 -------------------------

// 定义的日期格式
var timeParseMap = map[GuessStrType]string{
	TypeData:            "yyyy-MM-dd",
	TypeDateTimeMinutes: "yyyy-MM-dd HH:mm",
	TypeDateTimeSeconds: "yyyy-MM-dd HH:mm:ss",
	TypeDateTimeMills:   "yyyy-MM-dd HH:mm:ss.SSS",
}

// genDateFormatFunc
// 根据给定的格式，生成对应的日期格式化函数
func genDateFormatFunc(typ GuessStrType) FormatFunc {
	return FormatFunc{
		strType: typ,
		Compress: func(input string) (string, error) {
			return formatDate(typ, input)
		},
		Format: func(input string) (string, error) {
			return formatDate(typ, input)
		},
	}
}

// 10位和 13 位 时间戳格式化和压缩
func formatTimeStamp(input string) (string, error) {
	// 先转为 int64
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return input, errors.New("时间戳解析失败")
	}
	// 时间戳转为时间
	if len(input) == 10 {
		t := time.Unix(i, 0)
		return formatTime(t)
	} else if len(input) == 13 {
		t := time.UnixMilli(i)
		return formatTime(t)
	}
	return input, errors.New("时间戳解析失败")
}

func formatDate(guessType GuessStrType, input string) (string, error) {
	t, err := bjTimeYyyyMMddToTimeStamp(input, timeParseMap[guessType])
	if err != nil {
		return input, errors.New("按照日期格式 " + timeParseMap[guessType] + " 解析失败")
	}
	log.I("按照日期格式压缩", input, "->", timeParseMap[guessType])
	return formatTime(t)
}

// bjTimeYyyyMMddToTimeStamp
// 所有 yyyyMMdd 之类的时间字符串, 这里都默认为北京时间来解析成 time.Time
func bjTimeYyyyMMddToTimeStamp(timeStr string, formatStr string) (time.Time, error) {
	formatStr = getFormatString(formatStr)

	// 加载 UTC+8 时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}

	// 直接在 UTC+8 时区解析时间
	parsedTime, err := time.ParseInLocation(formatStr, timeStr, location)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

// formatTime 格式化时间
// 将一个时间用多种格式进行格式化，然后返回
func formatTime(t time.Time) (string, error) {
	var result = "[time parse]\n"

	// 时间戳格式
	// t 转为 10位时间戳
	result += strconv.FormatInt(t.Unix(), 10) + "\n"
	// t 转为 13位时间戳
	result += strconv.FormatInt(t.UnixNano()/1e6, 10) + "\n"

	result += "\n"

	// UTC+0 时间格式
	utcTime := t.UTC()
	result += "UTC+0:\n"
	result += date.Format(utcTime, timeParseMap[TypeData]) + "\n"
	result += date.Format(utcTime, timeParseMap[TypeDateTimeMinutes]) + "\n"
	result += date.Format(utcTime, timeParseMap[TypeDateTimeSeconds]) + "\n"
	result += date.Format(utcTime, timeParseMap[TypeDateTimeMills]) + "\n"

	result += "\n"

	// 北京时间格式 (UTC+8)
	beijingLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", err
	}
	beijingTime := t.In(beijingLocation)
	result += "UTC+8:\n"
	result += date.Format(beijingTime, timeParseMap[TypeData]) + "\n"
	result += date.Format(beijingTime, timeParseMap[TypeDateTimeMinutes]) + "\n"
	result += date.Format(beijingTime, timeParseMap[TypeDateTimeSeconds]) + "\n"
	result += date.Format(beijingTime, timeParseMap[TypeDateTimeMills]) + "\n"

	return result, nil
}

var dateFormatMap = map[string]string{}

// 传入 yyyyMMdd 之类的
func getFormatString(formatStr string) string {
	// 原始 format string
	formatStrOld := formatStr
	// 构建一个 formatStr
	if formatStr := dateFormatMap[formatStrOld]; formatStr == "" {
		formatStr = str.ReplaceAll(formatStrOld, "yyyy", "2006")
		formatStr = str.ReplaceAll(formatStr, "yy", "06")
		formatStr = str.ReplaceAll(formatStr, "MM", "01")
		formatStr = str.ReplaceAll(formatStr, "dd", "02")
		formatStr = str.ReplaceAll(formatStr, "HH", "15")
		formatStr = str.ReplaceAll(formatStr, "hh", "3")
		formatStr = str.ReplaceAll(formatStr, "mm", "04")
		formatStr = str.ReplaceAll(formatStr, "ss", "05")
		formatStr = str.ReplaceAll(formatStr, "SSS", "000")
		dateFormatMap[formatStrOld] = formatStr
	}
	return dateFormatMap[formatStrOld]
}
