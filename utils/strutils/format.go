package strutils

import (
	"errors"
	"github.com/Ericwyn/EzeFormat/log"
	"strings"
)

type StrHandlerFunc func(str string) (string, error)

type FormatFunc struct {
	// 推测出来的格式类型
	strType GuessStrType
	// 压缩函数
	Compress StrHandlerFunc
	// 格式化函数
	Format StrHandlerFunc
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

// FormatSmart
// 智能格式化, 猜测出来 input 的格式之后，根据不同的格式进行格式化
func FormatSmart(input string) (string, error) {
	// 有这些特殊字符的话，就不要压缩了
	if strings.Contains(input, "time parse") {
		return input, nil
	}

	guessType := GuessType(input)
	if guessType == TypeUnknown {
		return input, errors.New("无法推测的格式, 请手动选择格式化类型")
	}

	log.I("guess type and format: ", guessType)

	return FormatType(input, guessType)
}

// CompressSmart
// 智能压缩, 猜测出来 input 的格式之后，根据不同的格式进行压缩
func CompressSmart(input string) (string, error) {
	// 有这些特殊字符的话，就不要压缩了
	if strings.Contains(input, "time parse") {
		return input, nil
	}

	guessType := GuessType(input)
	if guessType == TypeUnknown {
		return input, errors.New("无法推测的格式, 请手动选择格式化类型")
	}

	log.I("guess type and compress: ", guessType)

	return CompressType(input, guessType)
}

// FormatType
// 根据给定的格式，格式化
func FormatType(input string, strType GuessStrType) (string, error) {
	if formatFunc, ok := FormatFuncMap[strType]; ok {
		return formatFunc.Format(input)
	}
	return input, errors.New("无法推测的格式, 请手动选择格式化类型")
}

// CompressType
// 根据给定的格式，压缩
func CompressType(input string, strType GuessStrType) (string, error) {
	if formatFunc, ok := FormatFuncMap[strType]; ok {
		return formatFunc.Compress(input)
	}
	return input, errors.New("无法推测的格式, 请手动选择格式化类型")
}
