package format

import (
	"encoding/json"
	"errors"
	"github.com/Ericwyn/EzeFormat/utils/strutils"
	"strings"
)

// ------------------------ JSON 格式化 ------------------------

// formatJson 格式化 JSON 字符串
func formatJson(jsonStr string) (string, error) {
	err := jsonPreCheck(jsonStr)
	if err != nil {
		return jsonStr, err
	}

	var jsonObj interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		//return "format error in unmarshal : " + err.Error() + "\n\n" + jsonStr
		return jsonStr, err
	}
	resByte, err := json.MarshalIndent(jsonObj, "", "    ")
	if err != nil {
		//return "format error in Marshal : " + err.Error() + "\n\n" + jsonStr
		return jsonStr, err
	}
	return string(resByte), nil
}

func compressJson(jsonStr string) (string, error) {
	err := jsonPreCheck(jsonStr)
	if err != nil {
		return jsonStr, err
	}

	var jsonObj interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonObj)
	if err != nil {
		//return "format error in unmarshal : " + err.Error() + "\n\n" + jsonStr
		return jsonStr, err
	}
	resByte, err := json.Marshal(jsonObj)
	if err != nil {
		//return "format error in Marshal : " + err.Error() + "\n\n" + jsonStr
		return jsonStr, err
	}
	return string(resByte), nil
}

func jsonPreCheck(input string) error {
	input = strutils.StringTrim(input)

	if (strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]")) ||
		(strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}")) {
		return nil
	} else {
		return errors.New("俺寻思这看着不太像 JSON 啊?")
	}
}
