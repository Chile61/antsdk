package utils

import (
	"encoding/json"
)

// ToJSON JSON
func ToJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}
