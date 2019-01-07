package helper

import (
	"github.com/huichen/sego"
)
const (
	VERSION = "v1.0Beta"
	
	CACHE_CONF = `{"CachePath":"./cache/rutime","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`
	DEFAULT_STATIC_EXT = ".html,.js,.css,.ico,.jpeg,.jpg,.png,.txt,.xml,.gif"
	DEFAULT_COOKIE_SECRET = "wemanago"
)

var (
	// when run mode is "dev", Debug mode is open
	Debug = beego.AppConfig.String("runmode") == "dev"
	// allowed visit direct statis file's extension
	StaticExt = make(map[string]bool)
	// Word segmenter
	Segmenter sego.Segmenter
	// system configure map
	ConfigMap sync.Map
	IsInstalled = false
	AllowedUploadExt = ",xlsx,xls,jpeg,jpg,gif,png"
)