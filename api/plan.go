package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// PlanGetPlanTemplate 销售计划-获取销售计划企业模板信息
func (c *Config) PlanGetPlanTemplate(reqData *dto.PlanGetPlanTemplateRequestDTO) (res *dto.PlanGetPlanTemplateResponseDTO, err error) {
	res = dto.NewPlanGetPlanTemplateResponseDTO()
	err = c.RequestByAutoResolve(common.PlanGetPlanTemplate, reqData, res)
	return
}
