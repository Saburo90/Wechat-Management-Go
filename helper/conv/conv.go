package conv

import (
	"encoding/json"
	"fmt"
	"strings"
)

// convert interface data into json data
func InterfaceToJson(itf interface{}) (string error) {
	bytes, err := json.Marshal(itf)
	return string(bytes), err
}


// convert interface into integer
func InterfaceToInteger(itf interface{}) (num int, err error) {
	return strconv.Atoi(fmt.Sprintf("%v", itf))
}