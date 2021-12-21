package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// OrgDeptCreate 创建部门
func (c *Config) OrgDeptCreate(reqData *dto.OrgDeptCreateRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.OrgDeptCreate, reqData, res)
	return
}

// OrgDeptEdit 编辑部门
func (c *Config) OrgDeptEdit(reqData *dto.OrgDeptEditRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.OrgDeptEdit, reqData, res)
	return
}

// OrgDeptDelete 删除部门
func (c *Config) OrgDeptDelete(reqData *dto.OrgDeptDeleteRequestDTO) (res *dto.OrgDeptDeleteResponseDTO, err error) {
	res = dto.NewOrgDeptDeleteResponseDTO()
	err = c.RequestByAutoResolve(common.OrgDeptDelete, reqData, res)
	return
}

// OrgStructInfo 架构信息
func (c *Config) OrgStructInfo() (res *dto.OrgStructInfoResponseDTO, err error) {
	res = dto.NewOrgStructInfoResponseDTO()
	err = c.RequestByAutoResolve(common.OrgStructInfo, nil, res)
	return
}

// OrgUserDelete 删除员工
func (c *Config) OrgUserDelete(reqData *dto.OrgUserDeleteRequestDTO) (res *dto.OrgUserDeleteResponseDTO, err error) {
	res = dto.NewOrgUserDeleteResponseDTO()
	err = c.RequestByAutoResolve(common.OrgUserDelete, reqData, res)
	return
}
