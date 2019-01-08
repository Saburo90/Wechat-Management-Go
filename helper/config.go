package helper

import (
	"fmt"
	"strconv"
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

//Get Int64 value configure item
//@param cate configure option type
//@param key configure option key
//@return val value of configure item
func GetConfigInt64(cate string, key string) (val int64) {
	val, _ := strconv.ParseInt(GetConfig(cate, key), 10, 64)
	return
}

//Get Float64 value configure item
//@param cate configure option type
//@param key configure option key
//@return val value of configure item
func GetConfigFloat64(cate string, key string) (val float64) {
	val, _ := strconv.ParseFloat(GetConfig(cate, key), 64)
	return
}


func setDefaultConfig() {
	// the following configure item must configured by before the program has been installed
	if !IsInstalled {
		beego.BConfig.AppName = "Wemanago" // default name of your program
		beego.BConfig.Listen.HTTPPORT = 8888 // default http service listen port
		beego.BConfig.RunMode = "dev" // default setup development environment
		beego.BConfig.EnableGzip = true // default setup open gizp compressed function
		beego.BConfig.WebConfig.EnableXSRF = false // if the program is not installed, that unable the XSRF
		beego.BConfig.WebConfig.XSRFkey = 
	}
}