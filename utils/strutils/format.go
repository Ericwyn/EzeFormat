package strutils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Ericwyn/EzeFormat/log"
	"github.com/Ericwyn/GoTools/date"
	"io"
	"strconv"
	"strings"
	"time"
)

type StrHandlerFunc func(str string) string

type FormatFunc struct {
	strType  GuessStrType
	Compress StrHandlerFunc
	Format   StrHandlerFunc
}

// FormatFuncMap 推测出类型之后, 用这个方法来做解析
var FormatFuncMap = map[GuessStrType]FormatFunc{
	TypeData:            genDateFormatFunc(TypeData),
	TypeDateTimeMinutes: genDateFormatFunc(TypeDateTimeMinutes),
	TypeDateTimeSeconds: genDateFormatFunc(TypeDateTimeSeconds),
	TypeDateTimeMills:   genDateFormatFunc(TypeDateTimeMills),
	TypeTimeStamp: {
		strType:  TypeTimeStamp,
		Compress: formatTimeStamp,
		Format:   formatTimeStamp,
	},
	TypeTimeStampMills: {
		strType:  TypeTimeStampMills,
		Compress: formatTimeStamp,
		Format:   formatTimeStamp,
	},

	TypeJson: {
		strType:  TypeJson,
		Compress: compressJson,
		Format:   formatJson,
	},

	TypeXml: {
		strType:  TypeXml,
		Compress: compressXml,
		Format:   formatXml,
	},
}

// FormatSmart 智能格式化, 猜测出来 input 的格式之后，根据不同的格式进行格式化
func FormatSmart(input string) string {
	// 有这些特殊字符的话，就不要压缩了
	if strings.Contains(input, "time parse") {
		return input
	}

	guessType := GuessType(input)
	if guessType == TypeUnknown {
		return "无法推测的格式, 请手动选择格式化类型\n\n" + input
	}

	log.I("guess type and format: ", guessType)

	return FormatType(input, guessType)
}

// CompressSmart 智能压缩, 猜测出来 input 的格式之后，根据不同的格式进行压缩
func CompressSmart(input string) string {
	// 有这些特殊字符的话，就不要压缩了
	if strings.Contains(input, "time parse") {
		return input
	}

	guessType := GuessType(input)
	if guessType == TypeUnknown {
		return "无法推测的格式, 请手动选择格式化类型\n\n" + input
	}

	log.I("guess type and compress: ", guessType)

	return CompressType(input, guessType)
}

func FormatType(input string, strType GuessStrType) string {
	if formatFunc, ok := FormatFuncMap[strType]; ok {
		return formatFunc.Format(input)
	}
	return "无法推测的格式, 请手动选择格式化类型\n\n" + input
}

func CompressType(input string, strType GuessStrType) string {
	if formatFunc, ok := FormatFuncMap[strType]; ok {
		return formatFunc.Compress(input)
	}
	return "无法推测的格式, 请手动选择格式化类型\n\n" + input
}

// ------------------------ JSON 格式化 ------------------------

// formatJson 格式化 JSON 字符串
func formatJson(jsonStr string) string {
	var jsonObj interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		return "format error in unmarshal : " + err.Error() + "\n\n" + jsonStr
	}
	resByte, err := json.MarshalIndent(jsonObj, "", "    ")
	if err != nil {
		return "format error in Marshal : " + err.Error() + "\n\n" + jsonStr
	}
	return string(resByte)
}

func compressJson(jsonStr string) string {
	var jsonObj interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		return "format error in unmarshal : " + err.Error() + "\n\n" + jsonStr
	}
	resByte, err := json.Marshal(jsonObj)
	if err != nil {
		return "format error in Marshal : " + err.Error() + "\n\n" + jsonStr
	}
	return string(resByte)
}

// ------------------------ XML 格式化 ------------------------

// FormatXml 格式化 XML 字符串
func formatXml(xmlStr string) string {
	input := strings.NewReader(xmlStr)
	buf := &bytes.Buffer{}
	decoder := xml.NewDecoder(input)

	// 初始化缩进计数器
	indentCounter := 0

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err.Error() + "\n\n" + xmlStr
		}

		switch element := token.(type) {
		case xml.StartElement:
			// 在每个新的开始标签之前添加缩进
			buf.WriteString(strings.Repeat(" ", indentCounter*4))
			buf.WriteString("<" + element.Name.Local)

			// 添加属性（如果有）
			for _, attr := range element.Attr {
				attrLine := fmt.Sprintf(" %s=\"%s\"", attr.Name.Local, attr.Value)
				buf.WriteString(attrLine)
			}
			buf.WriteString(">\n")

			// 增加缩进计数器
			indentCounter++

		case xml.EndElement:
			// 减少缩进计数器
			indentCounter--

			// 在每个结束标签之前添加缩进
			buf.WriteString(strings.Repeat(" ", indentCounter*4))
			buf.WriteString("</" + element.Name.Local + ">\n")

		case xml.CharData:
			// 添加文本内容之前去除可能的空白字符
			content := strings.TrimSpace(string(element))
			if content != "" {
				buf.WriteString(strings.Repeat(" ", indentCounter*4))
				buf.WriteString(content + "\n")
			}
		}
	}

	return buf.String()
}

// CompressXml 压缩 XML 字符串
func compressXml(xmlStr string) string {
	input := strings.NewReader(xmlStr)
	buf := &bytes.Buffer{}
	decoder := xml.NewDecoder(input)

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err.Error() + "\n\n" + xmlStr
		}

		switch element := token.(type) {
		case xml.StartElement:
			buf.WriteString("<" + element.Name.Local)
			for _, attr := range element.Attr {
				attrLine := fmt.Sprintf(" %s=\"%s\"", attr.Name.Local, attr.Value)
				buf.WriteString(attrLine)
			}
			buf.WriteString(">")

		case xml.EndElement:
			buf.WriteString("</" + element.Name.Local + ">")

		case xml.CharData:
			content := strings.TrimSpace(string(element))
			if content != "" {
				buf.WriteString(content)
			}
		}
	}

	return buf.String()

}

func formatSql(sqlStr string) string {
	return sqlStr
}

func compressSql(sqlStr string) string {
	return sqlStr
}

// ------------------------- 日期格式化 -------------------------

var timeParseMap = map[GuessStrType]string{
	TypeData:            "yyyy-MM-dd",
	TypeDateTimeMinutes: "yyyy-MM-dd HH:mm",
	TypeDateTimeSeconds: "yyyy-MM-dd HH:mm:ss",
	TypeDateTimeMills:   "yyyy-MM-dd HH:mm:ss.SSS",
}

// 通用的日期格式化函数
func genDateFormatFunc(typ GuessStrType) FormatFunc {
	return FormatFunc{
		strType: typ,
		Compress: func(input string) string {
			return formatDate(typ, input)
		},
		Format: func(input string) string {
			return formatDate(typ, input)
		},
	}
}

// 10位和 13 位 时间戳格式化和压缩
func formatTimeStamp(input string) string {
	// 先转为 int64
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "时间戳解析失败: " + input
	}
	// 时间戳转为时间
	if len(input) == 10 {
		t := time.Unix(i, 0)
		return formatTime(t)
	} else if len(input) == 13 {
		t := time.UnixMilli(i)
		return formatTime(t)
	}
	return "时间戳解析失败: " + input
}

func formatDate(guessType GuessStrType, input string) string {
	t, err := date.ToTime(input, timeParseMap[guessType])
	if err != nil {
		return "按照日期格式 " + timeParseMap[guessType] + " 解析失败: " + input
	}
	log.I("按照日期格式压缩", input, "->", timeParseMap[guessType])
	return formatTime(t)
}

// formatTime 格式化时间
// 将一个时间用多种格式进行格式化，然后返回
func formatTime(t time.Time) string {
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

	return result
}
