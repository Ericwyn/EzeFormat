package cmd

import (
	"fmt"
	"github.com/Ericwyn/EzeFormat/utils/format"
)

func Format(input string, strType string) {
	if input == "" {
		//return "", nil
	}
	var formatType format.GuessStrType
	if strType == "json" {
		formatType = format.TypeJson
	} else if strType == "xml" {
		formatType = format.TypeXml
	} else {
		formatType = format.GuessType(input)
	}

	s, err := format.FormatType(input, formatType)
	if err != nil {
		printFormatResult("", "", err)
	} else {
		printFormatResult(s, formatType, nil)
	}
}

func printFormatResult(result string, strType format.GuessStrType, err error) {
	if err != nil {
		fmt.Println("[format error]: ", err)
		return
	}
	fmt.Println("[format type]: ", strType)
	fmt.Println()
	fmt.Println(result)
}
