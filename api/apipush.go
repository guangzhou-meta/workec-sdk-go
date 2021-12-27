package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// ApiPushGetApiPush 查询主动推送接口
func (c *Config) ApiPushGetApiPush(reqData *dto.ApiPushGetApiPushRequestDTO) (res *dto.ApiPushGetApiPushResponseDTO, err error) {
	res = dto.NewApiPushGetApiPushRequestDTO()
	err = c.RequestByAutoResolve(common.ApiPushGetApiPush, reqData, res)
	return
}
