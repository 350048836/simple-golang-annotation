package jsonutils

import (
	"../logger"
	"encoding/json"
)

func Json2Obj(str string, v interface{}) {
	err := json.Unmarshal([]byte(str), v)
	if err != nil {
		logger.Info("Json2Obj Error: %v", err)
	}
}

func Obj2Json(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
