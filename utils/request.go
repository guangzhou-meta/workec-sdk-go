package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
)

type Request struct {
	appId         string
	appSecret     string
	corpId        string
	serverBaseUrl string
}

// NewRequest 创建请求
func NewRequest(appId string, appSecret string, corpId string, serverBaseUrl *string) *Request {
	instance := &Request{}
	if serverBaseUrl != nil {
		instance.SetServerBaseUrl(*serverBaseUrl)
	}
	return instance.
		SetAppId(appId).
		SetAppSecret(appSecret).
		SetCorpId(corpId)
}

func (m *Request) SetAppId(appId string) *Request {
	m.appId = appId
	return m
}

func (m *Request) SetAppSecret(appSecret string) *Request {
	m.appSecret = appSecret
	return m
}

func (m *Request) SetCorpId(corpId string) *Request {
	m.corpId = corpId
	return m
}

func (m *Request) SetServerBaseUrl(serverBaseUrl string) *Request {
	m.serverBaseUrl = serverBaseUrl
	return m
}

func (m *Request) GET(path string, reqData map[string]string, resData interface{}) (err error) {
	reqUrl := fmt.Sprintf("%s/%s", m.serverBaseUrl, path)
	params := url.Values{}
	for k, v := range reqData {
		params.Set(k, v)
	}
	requestUrl, err := url.Parse(reqUrl)
	if err != nil {
		return
	}
	requestUrl.RawQuery = params.Encode()
	reqUrl = requestUrl.String()
	req, err := http.NewRequest("GET", reqUrl, nil)
	AddHeader(req, m.corpId, m.appId, m.appSecret)
	return RequestAndResolve(req, resData)
}

func (m *Request) POST(path string, reqData interface{}, resData interface{}) (err error) {
	reqUrl := fmt.Sprintf("%s/%s", m.serverBaseUrl, path)
	bodyByte, err := json.Marshal(reqData)
	if err != nil {
		return
	}
	body := bytes.NewReader(bodyByte)
	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		return
	}
	AddHeader(req, m.corpId, m.appId, m.appSecret)
	return RequestAndResolve(req, resData)
}

func (m *Request) Upload(path string, reqData interface{}, resData interface{}) (err error) {
	reqUrl := fmt.Sprintf("%s/%s", m.serverBaseUrl, path)
	bodyByte, err := json.Marshal(reqData)
	if err != nil {
		return
	}
	body := bytes.NewReader(bodyByte)
	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		return
	}
	AddHeader(req, m.corpId, m.appId, m.appSecret)
	req.Header.Set(common.ECRequestContentType, common.ECRequestContentTypeFile)
	return RequestAndResolve(req, resData)
}
