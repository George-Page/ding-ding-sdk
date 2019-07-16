package dingTalk

import (
	"crypto/cipher"
	"time"
	"fmt"
)

// HTTPTimeout http timeout
type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

// Config dingTalk configure
type Config struct {
	AccessTokenUrl string
	AppKey         string
	AppSecret      string
	AppAesKey      string
	Token          string
	EncodingAesKey string
	SuiteKey       string
	Block          cipher.Block
	Key            []byte
	RetryTimes     uint        // 失败重试次数
	Timeout        uint        // 超时时间 默认60s
	HTTPTimeout    HTTPTimeout // HTTP的超时时间设置
}

// 获取默认配置
func getDefaultDingTalkConfig() *Config {
	config := Config{}

	config.AccessTokenUrl = "https://oapi.dingtalk.com/gettoken"
	config.AppKey = ""
	config.AppSecret = ""
	config.Token = ""
	config.EncodingAesKey = ""
	config.SuiteKey = ""
	config.RetryTimes = 3
	config.Timeout = 60 // seconds

	config.HTTPTimeout.ConnectTimeout = time.Second * 30   // 30s
	config.HTTPTimeout.ReadWriteTimeout = time.Second * 60 // 60s
	config.HTTPTimeout.HeaderTimeout = time.Second * 60    // 60s
	config.HTTPTimeout.LongTimeout = time.Second * 300     // 300s
	config.HTTPTimeout.IdleConnTimeout = time.Second * 50  // 50s

	return &config
}

func BuildHttpGetParams(url string, val map[string]interface{}) string {
	var queryStr string
	for v, k := range val {
		queryStr += fmt.Sprintf("&%v=%v", v, k)
	}

	return fmt.Sprintf("%s?%s", url, queryStr[1:])
}
