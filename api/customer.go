package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// CustomerAddCustomer 新增客户
func (c *Config) CustomerAddCustomer(reqData *dto.CustomerAddCustomerRequestDTO) (res *dto.CustomerAddCustomerResponseDTO, err error) {
	res = dto.NewCustomerAddCustomerResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerAddCustomer, reqData, res)
	return
}

// CustomerUpdateCustomer 修改客户信息
func (c *Config) CustomerUpdateCustomer(reqData *dto.CustomerUpdateCustomerRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.CustomerUpdateCustomer, reqData, res)
	return
}

// CustomerBatchUpdateCustomer 批量修改客户信息
func (c *Config) CustomerBatchUpdateCustomer(reqData *dto.CustomerBatchUpdateCustomerRequestDTO) (res *dto.CustomerBatchUpdateCustomerResponseDTO, err error) {
	res = dto.NewCustomerBatchUpdateCustomerResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerBatchUpdateCustomer, reqData, res)
	return
}

// CustomerCombine 合并客戶
func (c *Config) CustomerCombine(reqData *dto.CustomerCombineRequestDTO) (res *dto.CustomerCombineResponseDTO, err error) {
	res = dto.NewCustomerCombineResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerCombine, reqData, res)
	return
}

// CustomerFileUpload 客户资料-文件上传(支持批量)
func (c *Config) CustomerFileUpload(reqData *dto.CustomerFileUploadRequestDTO) (res *dto.CustomerFileUploadResponseDTO, err error) {
	res = dto.NewCustomerFileUploadResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerFileUpload, reqData, res)
	return
}

// CustomerShare 客户-共享客户
func (c *Config) CustomerShare(reqData *dto.CustomerShareRequestDTO) (res *dto.CustomerShareResponseDTO, err error) {
	res = dto.NewCustomerShareResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerShare, reqData, res)
	return
}

// CustomerChangeAbandon 放弃客户
func (c *Config) CustomerChangeAbandon(reqData *dto.CustomerChangeAbandonRequestDTO) (res *dto.CustomerChangeAbandonResponseDTO, err error) {
	res = dto.NewCustomerChangeAbandonResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerChangeAbandon, reqData, res)
	return
}

// CustomerChangeUser 变更跟进人(支持批量)
func (c *Config) CustomerChangeUser(reqData *dto.CustomerChangeUserRequestDTO) (res *dto.CustomerChangeUserResponseDTO, err error) {
	res = dto.NewCustomerChangeUserResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerChangeUser, reqData, res)
	return
}

// CustomerFace 客户-获取客户头像
func (c *Config) CustomerFace(reqData *dto.CustomerFaceRequestDTO) (res *dto.CustomerFaceResponseDTO, err error) {
	res = dto.NewCustomerFaceResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerFace, reqData, res)
	return
}

// CustomerQueryLabel 查询客户标签(支持批量)
func (c *Config) CustomerQueryLabel(reqData *dto.CustomerQueryLabelRequestDTO) (res *dto.CustomerQueryLabelResponseDTO, err error) {
	res = dto.NewCustomerQueryLabelResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryLabel, reqData, res)
	return
}

// CustomerQueryList 查询客户列表
func (c *Config) CustomerQueryList(reqData *dto.CustomerQueryListRequestDTO) (res *dto.CustomerQueryListResponseDTO, err error) {
	res = dto.NewCustomerQueryListResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryList, reqData, res)
	return
}

// CustomerFileList 客户资料-文件列表查询
func (c *Config) CustomerFileList(reqData *dto.CustomerFileListRequestDTO) (res *dto.CustomerFileListResponseDTO, err error) {
	res = dto.NewCustomerFileListResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerFileList, reqData, res)
	return
}

// CustomerFolderList 客户资料-文件目录查询
func (c *Config) CustomerFolderList(reqData *dto.CustomerFolderListRequestDTO) (res *dto.CustomerFolderListResponseDTO, err error) {
	res = dto.NewCustomerFolderListResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerFolderList, reqData, res)
	return
}

// CustomerQueryExist 手机查询客户(判重)
func (c *Config) CustomerQueryExist(reqData *dto.CustomerQueryExistRequestDTO) (res *dto.CustomerQueryExistResponseDTO, err error) {
	res = dto.NewCustomerQueryExistResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryExist, reqData, res)
	return
}

// CustomerQuery 分页查询客户信息 (建议使用 查询客户列表 接口代替)
func (c *Config) CustomerQuery(reqData *dto.CustomerQueryRequestDTO) (res *dto.CustomerQueryResponseDTO, err error) {
	res = dto.NewCustomerQueryResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerQuery, reqData, res)
	return
}

// CustomerPreciseQuery 批量查询客户信息(建议使用 查询客户列表 接口代替)
func (c *Config) CustomerPreciseQuery(reqData *dto.CustomerPreciseQueryRequestDTO) (res *dto.CustomerPreciseQueryResponseDTO, err error) {
	res = dto.NewCustomerPreciseQueryResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerPreciseQuery, reqData, res)
	return
}

// CustomerQueryCustomer 导出客户
func (c *Config) CustomerQueryCustomer(reqData *dto.CustomerQueryCustomerRequestDTO) (res *dto.CustomerQueryCustomerResponseDTO, err error) {
	res = dto.NewCustomerQueryCustomerResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryCustomer, reqData, res)
	return
}

// CustomerGetCustomerGroup 查询客户分组
func (c *Config) CustomerGetCustomerGroup(reqData *dto.CustomerGetCustomerGroupRequestDTO) (res *dto.CustomerGetCustomerGroupResponseDTO, err error) {
	res = dto.NewCustomerGetCustomerGroupResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerGetCustomerGroup, reqData, res)
	return
}

// CustomerDelcrms 客户-获取删除的客户
func (c *Config) CustomerDelcrms(reqData *dto.CustomerDelcrmsRequestDTO) (res *dto.CustomerDelcrmsResponseDTO, err error) {
	res = dto.NewCustomerDelcrmsResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerDelcrms, reqData, res)
	return
}

// CustomerGetTrajectory 分页查询客户轨迹
func (c *Config) CustomerGetTrajectory(reqData *dto.CustomerGetTrajectoryRequestDTO) (res *dto.CustomerGetTrajectoryResponseDTO, err error) {
	res = dto.NewCustomerGetTrajectoryResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerGetTrajectory, reqData, res)
	return
}

// CustomerGetCrmVisitDetails 客户 - 获取员工签到记录
func (c *Config) CustomerGetCrmVisitDetails(reqData *dto.CustomerGetCrmVisitDetailsRequestDTO) (res *dto.CustomerGetCrmVisitDetailsResponseDTO, err error) {
	res = dto.NewCustomerGetCrmVisitDetailsResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerGetCrmVisitDetails, reqData, res)
	return
}

// CustomerGetChannelSource 获取来源信息
func (c *Config) CustomerGetChannelSource() (res *dto.CustomerGetChannelSourceResponseDTO, err error) {
	res = dto.NewCustomerGetChannelSourceResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerGetChannelSource, nil, res)
	return
}

// CustomerGetCasCadeFieldMapping 客户 - 获取级联字段信息
func (c *Config) CustomerGetCasCadeFieldMapping(reqData *dto.CustomerGetCasCadeFieldMappingRequestDTO) (res *dto.CustomerGetCasCadeFieldMappingResponseDTO, err error) {
	res = dto.NewCustomerGetCasCadeFieldMappingResponseDTO()
	err = c.RequestByAutoResolve(common.CustomerGetCasCadeFieldMapping, reqData, res)
	return
}
