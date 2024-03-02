package format

import (
	"errors"
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/GoTools/date"
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
	t, err := date.ToTime(input, timeParseMap[guessType])
	if err != nil {
		return input, errors.New("按照日期格式 " + timeParseMap[guessType] + " 解析失败")
	}
	log.I("按照日期格式压缩", input, "->", timeParseMap[guessType])
	return formatTime(t)
}

// formatTime 格式化时间
// 将一个时间用多种格式进行格式化，然后返回
func formatTime(t time.Time) (string, error) {
	var result = "[time parse]\n"

	result += "\n"

	// 时间戳格式
	// t 转为 10位时间戳
	result += strconv.FormatInt(t.Unix(), 10) + "\n"
	// t 转为 13位时间戳
	result += strconv.FormatInt(t.UnixNano()/1e6, 10) + "\n"

	result += "\n"

	// 年月日格式
	result += date.Format(t, timeParseMap[TypeData]) + "\n"
	result += date.Format(t, timeParseMap[TypeDateTimeMinutes]) + "\n"
	result += date.Format(t, timeParseMap[TypeDateTimeSeconds]) + "\n"
	result += date.Format(t, timeParseMap[TypeDateTimeMills]) + "\n"

	return result, nil
}
