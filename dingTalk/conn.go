package dingTalk

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"bytes"
	"os"
	"strconv"
	"io/ioutil"
)

// Conn dingTalk conn
type Conn struct {
	config *Config
	client *http.Client
}

// init 初始化Conn
func (conn *Conn) Init(config *Config) error {
	// new Transport
	transport := newTransport(conn, config)

	conn.config = config
	conn.client = &http.Client{Transport: transport}

	return nil
}

// 发起http请求
func (conn Conn) doRequest(method, gateWay string, headers map[string]string, data io.Reader) (*Response, error) {
	method = strings.ToUpper(method)
	uri, err := url.ParseRequestURI(gateWay)
	if err != nil {
		return nil, err
	}
	req := &http.Request{
		Method:     method,
		URL:        uri,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       uri.Host,
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Set(HTTPHeaderDate, date)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	reader := data
	switch v := data.(type) {
	case *bytes.Buffer:
		req.ContentLength = int64(v.Len())
	case *bytes.Reader:
		req.ContentLength = int64(v.Len())
	case *strings.Reader:
		req.ContentLength = int64(v.Len())
	case *os.File:
		req.ContentLength = tryGetFileSize(v)
	case *io.LimitedReader:
		req.ContentLength = int64(v.N)
	}
	req.Header.Set(HTTPHeaderContentLength, strconv.FormatInt(req.ContentLength, 10))
	// http body
	r, ok := reader.(io.ReadCloser)
	if !ok && reader != nil {
		r = ioutil.NopCloser(reader)
	}
	req.Body = r

	resp, err := conn.client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       resp.Body,
	}, nil
}

// 获取文件大小
func tryGetFileSize(f *os.File) int64 {
	fInfo, _ := f.Stat()
	return fInfo.Size()
}
