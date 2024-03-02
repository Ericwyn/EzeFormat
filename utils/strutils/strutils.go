package strutils

import "strings"

func StringTrim(str string) string {
	str = strings.Trim(str, " ")
	//str = strings.Trim(str, "\n\t")
	str = strings.Trim(str, "\r\n")
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	//str = strings.Trim(str, "\t")
	return str
}
