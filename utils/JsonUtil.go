package utils

import (
	"encoding/json"
	"strings"
)

// ToJSON JSON
func ToJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

// JSONUnescapeString JSONUnescapeString
func JSONUnescapeString(in string) (out string) {
	out = in

	// out = strings.Replace(out, `\\`, `\`, -1)
	// out = strings.Replace(out, `\"`, `"`, -1)
	out = strings.Replace(out, `\/`, `/`, -1)
	// out = strings.Replace(out, `\b`, "\b", -1)
	// out = strings.Replace(out, `\f`, "\f", -1)
	// out = strings.Replace(out, `\n`, "\n", -1)
	// out = strings.Replace(out, `\r`, "\r", -1)
	// out = strings.Replace(out, `\t`, "\t", -1)

	return
}
