package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// StepUpdate 修改客户阶段(支持批量)
func (c *Config) StepUpdate(reqData *dto.StepUpdateRequestDTO) (res *dto.StepUpdateResponseDTO, err error) {
	res = dto.NewStepUpdateResponseDTO()
	err = c.RequestByAutoResolve(common.StepUpdate, reqData, res)
	return
}
