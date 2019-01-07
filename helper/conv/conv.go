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

// Upper first char
func UpperFirst(str string) string {
	return strings.Replace(str, str[0:1], strings.ToUpper(str[0:1], 1))
}

// convert url path parameter to type map
func PathParameterToMap(path string) map[string]string {
	var data = make(map[string]string)
	path = strings.Trim(path, "/")
	pathSlice := strings.Split(path, "/")
	pathSliceLen := len(pathSlice)

	if pathSliceLen % 2 == 1 {
		pathSliceLen--
	}

	if pathSliceLen > 0 {
		for i := 0; i < pathSliceLen; {
			data[pathSlice[i]] = pathSlice[(i + 1)]
			i = i + 2
		}
	}
	return data
}

// convert url query parameter to type map
func QueryParameterToMap(query string) map[string]string {
	var data = make(map[string]string)
	query = strings.Trim(query, "=&")
	querySlice := strings.Split(query, "&")

	if len(querySlice) > 0 {
		for _, queryStr := range querySlice {
			queryUnit := strings.Split(queryStt, "=")
			if len(queryUnit) == 2 {
				data[queryUnit[0]] = queryUnit[1]
			}
		}
	}

	return data
}
