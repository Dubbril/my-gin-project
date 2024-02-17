package helper

import "encoding/json"

func IsValidJSON(data []byte) bool {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	return err == nil
}
