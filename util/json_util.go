package util

import "encoding/json"

func Struct2String(v interface{}) string {
	content, err := json.Marshal(v)
	if len(content) != 0 {
		return string(content)
	}
	return err.Error()
}
