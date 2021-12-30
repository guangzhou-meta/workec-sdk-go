package dto

import (
	"fmt"
)

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
)

func rolePrivilegesTransform(src []*common.RolePrivilege) []string {
	privilegeList := make([]string, 0)
	if src != nil {
		for _, item := range src {
			if item != nil {
				privilegeList = append(privilegeList, item.GetId())
			}
		}
	}
	return privilegeList
}

type RoleAddRequestDTO struct {
	UserId        int64                   `json:"userId"`
	RoleName      string                  `json:"roleName"`
	PrivilegeList []*common.RolePrivilege `json:"privilegeList"`
}

type RoleAddRequestSimpleDTO struct {
	UserId        int64    `json:"userId"`
	RoleName      string   `json:"roleName"`
	PrivilegeList []string `json:"privilegeList"`
}

func NewRoleAddRequestSimpleDTO() *RoleAddRequestSimpleDTO {
	return &RoleAddRequestSimpleDTO{}
}

func (r *RoleAddRequestSimpleDTO) Assign(src *RoleAddRequestDTO) (*RoleAddRequestSimpleDTO, error) {
	if src == nil {
		return nil, fmt.Errorf("请求参数[dto.RoleAddRequestDTO]不能为空")
	}
	r.UserId = src.UserId
	r.RoleName = src.RoleName
	r.PrivilegeList = rolePrivilegesTransform(src.PrivilegeList)
	return r, nil
}

type RoleAddResponseDTO struct {
	CommonDTO
	Data *RoleAddResponseData `json:"data"`
}

func NewRoleAddResponseDTO() *RoleAddResponseDTO {
	return &RoleAddResponseDTO{}
}

type RoleAddResponseData struct {
	CorpId   int64  `json:"corpId"`
	UserId   int64  `json:"userId"`
	RoleName string `json:"roleName"`
	RoleId   int64  `json:"roleId"`
}

type RoleDeleteRequestDTO struct {
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}

type RoleChangeNameRequestDTO struct {
	UserId   int64  `json:"userId"`
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
}

type RoleChangePrivilegesRequestDTO struct {
	UserId        int64                   `json:"userId"`
	RoleId        int64                   `json:"roleId"`
	AddPrivileges []*common.RolePrivilege `json:"addPrivileges"`
	DelPrivileges []*common.RolePrivilege `json:"delPrivileges"`
}

type RoleChangePrivilegesRequestSimpleDTO struct {
	UserId        int64    `json:"userId"`
	RoleId        int64    `json:"roleId"`
	AddPrivileges []string `json:"addPrivileges"`
	DelPrivileges []string `json:"delPrivileges"`
}

func NewRoleChangePrivilegesRequestSimpleDTO() *RoleChangePrivilegesRequestSimpleDTO {
	return &RoleChangePrivilegesRequestSimpleDTO{}
}

func (r *RoleChangePrivilegesRequestSimpleDTO) Assign(src *RoleChangePrivilegesRequestDTO) (*RoleChangePrivilegesRequestSimpleDTO, error) {
	if src == nil {
		return nil, fmt.Errorf("请求参数[dto.RoleChangePrivilegesRequestDTO]不能为空")
	}
	r.UserId = src.UserId
	r.RoleId = src.RoleId
	r.AddPrivileges = rolePrivilegesTransform(src.AddPrivileges)
	r.DelPrivileges = rolePrivilegesTransform(src.DelPrivileges)
	return r, nil
}

type RoleChangeUsersRequestDTO struct {
	UserId     int64                            `json:"userId"`
	RoleId     int64                            `json:"roleId"`
	ModifyType *common.RoleChangeUserModifyType `json:"modifyType"`
	UserIds    []int64                          `json:"userIds"`
}

type RoleChangeUsersRequestSimpleDTO struct {
	UserId     int64   `json:"userId"`
	RoleId     int64   `json:"roleId"`
	ModifyType int     `json:"modifyType"`
	UserIds    []int64 `json:"userIds"`
}

func NewRoleChangeUsersRequestSimpleDTO() *RoleChangeUsersRequestSimpleDTO {
	return &RoleChangeUsersRequestSimpleDTO{}
}

func (r *RoleChangeUsersRequestSimpleDTO) Assign(src *RoleChangeUsersRequestDTO) (*RoleChangeUsersRequestSimpleDTO, error) {
	if src == nil {
		return nil, fmt.Errorf("请求参数[dto.RoleChangeUsersRequestDTO]不能为空")
	}
	r.UserId = src.UserId
	r.RoleId = src.RoleId
	r.UserIds = src.UserIds
	r.ModifyType = int(*src.ModifyType)
	return r, nil
}

type RoleListRequestDTO struct {
	UserId int64 `json:"userId"`
}

type RoleListResponseDTO struct {
	CommonDTO
	Data []*RoleListItem `json:"data"`
}

func NewRoleListResponseDTO() *RoleListResponseDTO {
	return &RoleListResponseDTO{}
}

type RoleListItem struct {
	CorpId   int64  `json:"corpId"`
	UserId   int64  `json:"userId"`
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	Sort     int64  `json:"sort"`
}
