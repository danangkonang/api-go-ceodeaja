package helper

import (
	"encoding/json"
	"net/http"
)

func MakeResponJson(value []byte, w http.ResponseWriter) {
	res := string(value)
	var responRequired map[string]interface{}
	json.Unmarshal([]byte(res), &responRequired)
	w.Write(value)
}
