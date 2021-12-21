package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
	"github.com/guangzhou-meta/workec-sdk-go/logger"
)

// RolePrivilegeList 获取权限点列表
func (c *Config) RolePrivilegeList() []*common.RolePrivilege {
	return common.GetRolePrivilegeList()
}

// RoleAdd 添加角色
func (c *Config) RoleAdd(reqData *dto.RoleAddRequestDTO) (res *dto.RoleAddResponseDTO, err error) {
	reqSimpleData, err := dto.NewRoleAddRequestSimpleDTO().Assign(reqData)
	if err != nil {
		logger.Error(err)
		return
	}
	res = dto.NewRoleAddResponseDTO()
	err = c.RequestByAutoResolve(common.RoleAdd, reqSimpleData, res)
	return
}

// RoleDelete 删除角色
func (c *Config) RoleDelete(reqData *dto.RoleDeleteRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RoleDelete, reqData, res)
	return
}

// RoleChangeName 修改角色名称
func (c *Config) RoleChangeName(reqData *dto.RoleChangeNameRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RoleChangeName, reqData, res)
	return
}

// RoleChangePrivileges 新增，删除角色的权限点
func (c *Config) RoleChangePrivileges(reqData *dto.RoleChangePrivilegesRequestDTO) (res *dto.CommonDTO, err error) {
	reqSimpleData, err := dto.NewRoleChangePrivilegesRequestSimpleDTO().Assign(reqData)
	if err != nil {
		logger.Error(err)
		return
	}
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RoleChangePrivileges, reqSimpleData, res)
	return
}

// RoleChangeUsers 给角色分配成员（或者移除成员）
func (c *Config) RoleChangeUsers(reqData *dto.RoleChangeUsersRequestDTO) (res *dto.CommonDTO, err error) {
	reqSimpleData, err := dto.NewRoleChangeUsersRequestSimpleDTO().Assign(reqData)
	if err != nil {
		logger.Error(err)
		return
	}
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RoleChangeUsers, reqSimpleData, res)
	return
}

// RoleList 查询企业所有的角色信息
func (c *Config) RoleList(reqData *dto.RoleListRequestDTO) (res *dto.RoleListResponseDTO, err error) {
	res = dto.NewRoleListResponseDTO()
	err = c.RequestByAutoResolve(common.RoleList, reqData, res)
	return
}
