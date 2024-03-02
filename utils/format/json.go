package format

import "encoding/json"

// ------------------------ JSON 格式化 ------------------------

// formatJson 格式化 JSON 字符串
func formatJson(jsonStr string) (string, error) {
	var jsonObj interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
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
	var jsonObj interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonObj)
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
