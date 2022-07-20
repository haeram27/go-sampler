package util

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PaesslerAG/jsonpath"
)

// unescape html in Json
func JsonUnEscape(value string) string {
	buffer := &bytes.Buffer{}
	bufio.NewWriter(buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(value)

	return buffer.String()
}

func PrintPrettyJson(jsonBlob []byte) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, jsonBlob, "", "  ")
	if err != nil {
		fmt.Println("JSON parse error: ", err.Error())
	} else {
		fmt.Println(prettyJSON.String())
	}
}

func PrettyJson(jsonBlob []byte) *bytes.Buffer {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, jsonBlob, "", "  ")
	return &prettyJSON
}

func makeInterfaceArray(i interface{}) []interface{} {
	switch x := i.(type) {
	case []interface{}:
		return i.([]interface{})
	case interface{}:
		var arr []interface{}
		arr = append(arr, i)
		return arr
	default:
		_ = x
		fmt.Println(errors.New("invalid type"))
		return []interface{}{}
	}
}

func JsonPath(jsonBlob []byte, path string) []interface{} {
	jsonStrt := interface{}(nil)
	err := json.Unmarshal(jsonBlob, &jsonStrt)
	if err != nil {
		fmt.Println(err)
		return []interface{}{}
	}

	result, err := jsonpath.Get(path, jsonStrt)
	if err != nil {
		fmt.Println(err)
		return []interface{}{}
	}

	/*
		jsonpath's result is
		interface{} (if size of result is 1)
		or
		[]interface{} (if size of result is 0 or size>1)

		make the result always as []interface{} for the convenient after process
	*/
	return makeInterfaceArray(result)
}
