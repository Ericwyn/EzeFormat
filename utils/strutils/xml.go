package strutils

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// ------------------------ XML 格式化 ------------------------

// FormatXml 格式化 XML 字符串
func formatXml(xmlStr string) (string, error) {
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
			//return err.Error() + "\n\n" + xmlStr
			return xmlStr, err
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

	return buf.String(), nil
}

// CompressXml 压缩 XML 字符串
func compressXml(xmlStr string) (string, error) {
	input := strings.NewReader(xmlStr)
	buf := &bytes.Buffer{}
	decoder := xml.NewDecoder(input)

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			//return err.Error() + "\n\n" + xmlStr
			return xmlStr, err
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

	return buf.String(), nil

}
