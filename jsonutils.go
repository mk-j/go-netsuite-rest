package netsuite

import (
	"bytes"
	"encoding/json"
)

func JsonEncode(input interface{}) string {
	//jsonValue, _ := json.Marshal(input)
	jsonValue, _ := json.MarshalIndent(input, "", "    ")
	return string(jsonValue)
}

func JsonDecode(input string) map[string]interface{} {
	var r map[string]interface{}
	json.Unmarshal([]byte(input), &r)
	return r
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

