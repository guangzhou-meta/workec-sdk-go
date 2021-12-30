package api

import (
	"fmt"
)

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/logger"
	"github.com/guangzhou-meta/workec-sdk-go/utils"
)

// Config 组合配置
type Config struct {
	r *utils.Request
}

// NewConfig 创建配置
func NewConfig(appId string, appSecret string, corpId string) *Config {
	instance := &Config{
		r: utils.NewRequest(appId, appSecret, corpId, nil),
	}
	return instance.SetServerBaseUrl(common.DefaultServerBaseUrl)
}

// SetAppId 设置AppId
func (c *Config) SetAppId(appId string) *Config {
	c.r.SetAppId(appId)
	return c
}

// SetAppSecret 设置AppSecret
func (c *Config) SetAppSecret(appSecret string) *Config {
	c.r.SetAppSecret(appSecret)
	return c
}

// SetCorpId 设置企业Id
func (c *Config) SetCorpId(corpId string) *Config {
	c.r.SetCorpId(corpId)
	return c
}

// SetServerBaseUrl 设置服务器地址，默认为(https://open.workec.com/v2/)
func (c *Config) SetServerBaseUrl(baseServer string) *Config {
	c.r.SetServerBaseUrl(baseServer)
	return c
}

// Request 通用请求方式
func (c *Config) Request(api *common.ApiModel, reqData interface{}, resData interface{}) (err error) {
	path := api.Path
	switch api.Method {
	case common.GET:
		var data map[string]string
		if reqData != nil {
			data = reqData.(map[string]string)
		}
		return c.r.GET(path, data, resData)
	case common.POST:
		return c.r.POST(path, reqData, resData)
	case common.FILE:
		return c.r.Upload(path, reqData, resData)
	}
	return fmt.Errorf("未配置相关服务")
}

// RequestByAutoResolve 通用请求方式(错误信息自动写入日志)
func (c *Config) RequestByAutoResolve(api *common.ApiModel, reqData interface{}, resData interface{}) (err error) {
	err = c.Request(api, reqData, resData)
	if err != nil {
		logger.Error(err)
	}
	return
}
