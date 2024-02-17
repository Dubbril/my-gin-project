package helper

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
)

func MapStruct(source interface{}, destination interface{}) (err error) {
	return copier.CopyWithOption(destination, source, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}

func ConvByteArrayToJsonOrArray(resp []byte) (any, error) {
	if resp != nil {
		// Convert byte to json raw message
		var rawJson json.RawMessage
		err := json.Unmarshal(resp, &rawJson)
		if err != nil {
			return nil, errors.New("not a json : cannot convert byte to json raw message")
		}

		// Convert json raw message to map
		var dataObj map[string]interface{}
		var dataArray []map[string]interface{}
		err = json.Unmarshal(rawJson, &dataObj)
		if err != nil {
			// Unmarshal the JSON array into the slice of maps
			err := json.Unmarshal(rawJson, &dataArray)
			if err != nil {
				return nil, errors.New("not a json : can not convert json raw raw message to map")
			}
			return dataArray, nil
		}
		return dataObj, nil
	}
	return nil, errors.New("response body is zero byte")
}
