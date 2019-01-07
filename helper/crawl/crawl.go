package crawl

import (
	"github.com/astaxie/beego/httplib"
	"strings"
)

// build up the http request
func BindRequest(method, url, referrer, cookie, os string, iscn, isjson bool, headers ...map[string]string) *httplib.BeegoHTTPRequest {
	var req *httplib.BeegoHTTPRequest
	switch strings.ToLower(method) {
	case "get":
		req = httplib.Get(url)
	case "post":
		req = httplib.Post(url)
	case "put":
		req = httplib.Put(url)
	case "delete":
		req = httplib.Delete(url)
	case "head":
		req = httplib.Head(url)
	default:
		req = httplib.Get(url)						
	}

	// setup the referrer of header
	if len(referrer) > 0 {
		req.Header("Referrer", referrer)
	}

	// setup the Cookie of http request
	if len(cookie) > 0 {
		req.Header("Cookie", cookie)
	}

	// setup the host of request
	hostSlice := strings.Split(url, "://")
	if len(hostSlice) > 1 {
		host := strings.Split(hostSlice[1], "/")[0]
		req.SetHost(host)
	}

	// setup the compressed
	req.Header("Accept-Encoding", "gzip, deflate, br")

	// setup support of language
	if iscn {
		req.Header("Accept-Language", "zh-CN,zh,q=0.8,en,q=0.6")
	} else {
		req.Header("Accept-Language", "en-US,en,q=0.8,zh,q=0.6")
	}

	// if it is json request, then setup the http header about json
	if isjson {
		req.Header("Accept", "application/json")
		req.Header("X-Request", "JSON")
		req.Header("X-Requested-With", "XMLHttpRequest")
	} else {
		req.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image.apng,*/*;q=0.8")
	}

	// simulate the type of client and service
	switch strings.ToLower(os) {
	case "windows":
		req.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")
	case "linux":
		req.Header("User-Agent", "Mozilla/5.0 (X11; U; Linux i686) AppleWebKit/534.15 (KHTML, like Gecko) Ubuntu/10.10 Chromium/10.0.613.0 Chrome/10.0.613.0 Safari/534.15")
	case "mac":
		req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3107.4 Safari/537.36")
	case "android":
		req.Header("User-Agent", "MQQBrowser/26 Mozilla/5.0 (Linux; U; Android 2.3.7; MB200 Build/GRJ22 CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")
	case "ios":
		req.Header("User-Agent", "Mozilla/5.0(iPhone; CPU iPhone OS 9_3_3 like Mac OS X)AppleWebKit/601.1.46(KHTML, like Gecko)Mobile/13G3")
	default:
		req.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1667.0 Safari/537.36")	
	}

	// setup the diy header part
	if len(headers) > 0 {
		for _, header := range headers {
			for key, item := range header {
				req.Header(key, item)
			}
		}
	}

	return req
}

// common crawl function
func Crawl(method, url, referrer, cookie, os string, iscn, isjson, isdebug bool, headers ...map[string]string) (string, error) {
	req := BindRequest(method, url, referrer, cookie, os, iscn, isjson, headers...)

	if isdebug {
		req.Debug(isdebug)
	}

	// setting the header
	if len(headers) > 0 {
		for _, header := range headers {
			for key, item := range header {
				req.Header(key, item)
			}
		}
	}

	return req.String()
}