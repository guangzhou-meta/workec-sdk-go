package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// SpecialSelect 查询指定的某个销售计划
func (c *Config) SpecialSelect(reqData *dto.SpecialSelectRequestDTO) (res *dto.SpecialSelectResponseDTO, err error) {
	res = dto.NewSpecialSelectResponseDTO()
	err = c.RequestByAutoResolve(common.SpecialSelect, reqData, res)
	return
}
