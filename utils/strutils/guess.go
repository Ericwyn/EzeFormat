package strutils

import (
	"github.com/Ericwyn/EzeFormat/log"
	"strconv"
	"strings"
)

// 猜测输入的文字的类型, 是 json 还是 xml 还是 timeStamp 还是 millsTimeStamp 还是未知

type GuessStrType string

const TypeJson GuessStrType = "JSON"
const TypeXml GuessStrType = "XML"

const TypeTimeStamp GuessStrType = "TimeStamp"           // 数字, 秒时间戳
const TypeTimeStampMills GuessStrType = "MillsTimeStamp" // 数字, 毫秒时间戳

const TypeData = "Data"                       // 日期 '2023-04-03'
const TypeDateTimeMinutes = "DateTimeMinutes" // 时间 '2023-04-03 00:00'
const TypeDateTimeSeconds = "DateTimeSeconds" // 时间 '2023-04-03 00:00:00'
const TypeDateTimeMills = "DateTimeMills"     // 时间 '2023-04-03 00:00:00.000'

const TypeUnknown GuessStrType = "Unknown"

func GuessType(str string) GuessStrType {
	log.D("guess ", str)

	// 先要去除开头和结尾的空格
	str = strings.Trim(str, " ")

	if (strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}")) ||
		(strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")) {
		return TypeJson
	}
	if strings.Contains(str, "?xml") || (strings.HasPrefix(str, "<") && strings.HasSuffix(str, ">")) {
		return TypeXml
	}

	// 判断是否是 unix 时间戳或毫秒时间戳
	if isNum(str) {
		if len(str) == 10 {
			return TypeTimeStamp
		} else if len(str) == 13 {
			return TypeTimeStampMills
		}
	}

	if (strings.Contains(str, ":") || strings.Contains(str, "-")) &&
		(strings.HasPrefix(str, "20") || strings.HasPrefix(str, "19")) {
		// 有日期或时间
		if strings.Contains(str, ".") && len(str) > 20 {
			// 有毫秒
			return TypeDateTimeMills
		} else if strings.Contains(str, ":") && len(str) > 16 {
			// 有秒
			return TypeDateTimeSeconds
		} else if strings.Contains(str, ":") && len(str) > 10 {
			// 有秒
			return TypeDateTimeMinutes
		} else if strings.Contains(str, "-") {
			// 有日期
			return TypeData
		}

	}

	return TypeUnknown
}

func isNum(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
