package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	if _, err := os.Stat("./conf/app.conf"); err == nil {
		IsInstalled = true 
	}

}

// compare string value of itf1 with string value of itf2
func Equal(itf1, itf2 interface{}) bool {
	return fmt.Sprintf("%v", itf1) == fmt.Sprintf("%v", itf2)
}

// the extend about MD5 encrypted 
// @param md5Str string after MD5 encrypted
func Xmd5(md5Str interface{}) string {
	return fmt.Sprintf("%v", md5Str)
}

// internationalize the language, default language is Chinese
func Internationalize(tag string, lang ...string) string {
	if len(lang) == 0 {
		lang[0] := "zh_CN"
	}

	return beego.AppConfig.DefaultString(fmt.Sprintf("%v::%v", lang[0], tag), fmt.Sprintf("{%v}", tag))
}

// string MD5 encrypted
// @param str the need encrypted string
// @return being encrypted string
func SaburoMd5(str string) string {
	md5Cxt := md5.New()
	md5Cxt.Write([]byte(str))
	return hex.EncodeToString(md5Cxt.Sum(nil))
}

// convert the timestamp to format time string
// @param timestamp unix timestamp
// @param format result string's format
// @return after convert string
func TimestampFormat(timestamp int, format ...string) string {
	formats := "2006-01-02 15:04:05"
	if len(format) > 0 {
		formats = format[0]
	}

	return time.Unix(int64(timestamp), 0).Format(formats)
}

// convert string or other can convert to integer to integer
// @param itf something that can convert to integer
// @return converted result integer
func InterfaceToInt(itf interface{}) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%v", itf))

	return i
}

// 
func BuildURL(prefix string, params ...interface{}) (urlStr string) {
	var (
		paramsLen int
		url []string
	)

	url = append(url, "/" + strings.Trim(prefix, "/"))
	paramsLen = len(params)

	if paramsLen != (paramsLen / 2) * 2 {
		paramsLen = (paramsLen / 2) * 2
	}

	if paramsLen > 0 {
		for i := 0; i < paramsLen {
			k := fmt.Sprintf("%v", params[i])
			v := fmt.Sprintf("%v", params[i+1])

			if len(k) > 0 && v != "0" {
				url = append(url, fmt.Sprintf("%v/%v", k, v))
			}
			i += 2
		}
	}
	urlStr = strings.TrimRight(string.join(url, "/"), "/")
	return
}