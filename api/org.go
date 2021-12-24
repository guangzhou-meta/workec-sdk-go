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

// OrgUserCreate 创建员工
func (c *Config) OrgUserCreate(reqData *dto.OrgUserCreateRequestDTO) (res *dto.OrgUserCreateResponseDTO, err error) {
	res = dto.NewOrgUserCreateResponseDTO()
	err = c.RequestByAutoResolve(common.OrgUserCreate, reqData, res)
	return
}

// OrgUserUpdateUser 修改员工
func (c *Config) OrgUserUpdateUser(reqData *dto.OrgUserUpdateUserRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.OrgUserUpdateUser, reqData, res)
	return
}

// OrgUserDelete 删除员工
func (c *Config) OrgUserDelete(reqData *dto.OrgUserDeleteRequestDTO) (res *dto.OrgUserDeleteResponseDTO, err error) {
	res = dto.NewOrgUserDeleteResponseDTO()
	err = c.RequestByAutoResolve(common.OrgUserDelete, reqData, res)
	return
}

// OrgUserUserOnOff 启用或禁用员工
func (c *Config) OrgUserUserOnOff(reqData *dto.OrgUserUserOnOffRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.OrgUserOnOff, reqData, res)
	return
}

// OrgUserFindUserInfoById 组织架构 - 获取指定员工信息
func (c *Config) OrgUserFindUserInfoById(reqData *dto.OrgUserFindUserInfoByIdRequestDTO) (res *dto.OrgUserFindUserInfoByIdResponseDTO, err error) {
	res = dto.NewOrgUserFindUserInfoByIdDTO()
	err = c.RequestByAutoResolve(common.OrgUserFindUserInfoById, reqData, res)
	return
}
