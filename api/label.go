package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// LabelAddLabelGroup 客户标签 - 创建标签分组
func (c *Config) LabelAddLabelGroup(reqData *dto.LabelAddLabelGroupRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.LabelAddLabelGroup, reqData, res)
	return
}

// LabelGetLabelInfo 获取标签信息
func (c *Config) LabelGetLabelInfo(reqData *dto.LabelGetLabelInfoRequestDTO) (res *dto.LabelGetLabelInfoResponseDTO, err error) {
	res = dto.NewLabelGetLabelInfoResponseDTO()
	err = c.RequestByAutoResolve(common.LabelGetLabelInfo, reqData, res)
	return
}

// LabelAddLabel 客户标签 - 创建标签
func (c *Config) LabelAddLabel(reqData *dto.LabelAddLabelRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.LabelAddLabel, reqData, res)
	return
}

// LabelUpdate 修改客户标签（支持批量）
func (c *Config) LabelUpdate(reqData *dto.LabelUpdateRequestDTO) (res *dto.LabelUpdateResponseDTO, err error) {
	res = dto.NewLabelUpdateResponseDTO()
	err = c.RequestByAutoResolve(common.LabelUpdate, reqData, res)
	return
}

// LabelDeleteCrmLabels 删除客户标签
func (c *Config) LabelDeleteCrmLabels(reqData *dto.LabelDeleteCrmLabelsRequestDTO) (res *dto.LabelDeleteCrmLabelsResponseDTO, err error) {
	res = dto.NewLabelDeleteCrmLabelsResponseDTO()
	err = c.RequestByAutoResolve(common.LabelDeleteCrmLabels, reqData, res)
	return
}

// LabelUpdateLabelGroupName 修改标签组名称
func (c *Config) LabelUpdateLabelGroupName(reqData *dto.LabelUpdateLabelGroupNameRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.LabelUpdateLabelGroupName, reqData, res)
	return
}

// ContactBookAdd 新增联系人
func (c *Config) ContactBookAdd(reqData *dto.ContactBookAddRequestDTO) (res *dto.ContactBookAddResponseDTO, err error) {
	res = dto.NewContactBookAddResponseDTO()
	err = c.RequestByAutoResolve(common.ContactBookAdd, reqData, res)
	return
}

// ContactBookDelete 删除联系人
func (c *Config) ContactBookDelete(reqData *dto.ContactBookDeleteRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.ContactBookDelete, reqData, res)
	return
}

// ContactBookUpdate 更新联系人
func (c *Config) ContactBookUpdate(reqData *dto.ContactBookUpdateRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.ContactBookUpdate, reqData, res)
	return
}

// ContactBookList 查询联系人
func (c *Config) ContactBookList(reqData *dto.ContactBookListRequestDTO) (res *dto.ContactBookListResponseDTO, err error) {
	res = dto.NewContactBookListResponseDTO()
	err = c.RequestByAutoResolve(common.ContactBookList, reqData, res)
	return
}

// ImMessageSend 发送通知
func (c *Config) ImMessageSend(reqData *dto.ImMessageSendRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.ImMessageSend, reqData, res)
	return
}
