package helper

import (
	"fmt"
)

//Get the configure item
//@param cate configure option type
//@param key configure option key
//@param def default value of configure item
//@return val value of configure item
func GetConfig(cate string, key string, def ...string) string {
	if val, ok := ConfigMap.Load(fmt.Sprintf("%v.%v", cate, key)); ok {
		return val.(string)
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

//Get Bool value configure item
//@param cate configure option type
//@param key configure option key
//@return val value of configure item
func GetConfigBool(cate string, key string) (val bool) {
	value := GetConfig(cate, key)

	if value == "true" || value == "1" {
		val = true
	}

	return
}


