package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/logger"
)

// GenSignDefault 生成签名
func GenSignDefault(appId string, appSecret string) (string, string) {
	timestamp := time.Now().UnixNano() / 1e6
	return GenSign(appId, appSecret, timestamp)
}

// GenSign 通过自定义时间戳生成签名
func GenSign(appId string, appSecret string, timestamp int64) (string, string) {
	timeStamp := fmt.Sprintf("%d", timestamp)
	signStr := fmt.Sprintf("appId=%s&appSecret=%s&timeStamp=%s", appId, appSecret, timeStamp)
	return timeStamp, strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(signStr))))
}

// AddHeader 添加通用请求头
func AddHeader(req *http.Request, corpId string, appId string, appSecret string) {
	timestamp, sign := GenSignDefault(appId, appSecret)
	req.Header.Set(common.ECRequestContentType, common.ECRequestContentTypeValue)
	req.Header.Set(common.ECRequestCorpId, corpId)
	req.Header.Set(common.ECRequestSign, sign)
	req.Header.Set(common.ECRequestTimeStamp, timestamp)
}

// RequestAndResolve 通用请求方式
func RequestAndResolve(req *http.Request, resData interface{}) (err error) {
	client := &http.Client{}
	logger.Debug("请求方法: ", req.Method)
	logger.Debug("请求头: ", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case 401:
		err = fmt.Errorf("无权限访问该服务")
		break
	case 404:
		err = fmt.Errorf("服务不存在")
		break
	case 502:
		err = fmt.Errorf("服务器正在维护")
		break
	}
	if err != nil {
		return
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	logger.Debug("响应体: ", string(resBody))
	err = json.Unmarshal(resBody, resData)
	return
}
