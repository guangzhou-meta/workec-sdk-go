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
	//res = dto.NewCustomerFileUploadResponseDTO()
	//err = c.RequestByAutoResolve(common.CustomerFileUpload, reqData, res)
	//TODO  文件上传
	return
}

// CustomerShare 客户-共享客户
func (c *Config) CustomerShare(reqData *dto.CustomerShareRequestDTO) (res *dto.CustomerShareResponseDTO, err error) {
	res = dto.NewCustomerShareRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerShare, reqData, res)
	return
}

// CustomerChangeAbandon 放弃客户
func (c *Config) CustomerChangeAbandon(reqData *dto.CustomerChangeAbandonRequestDTO) (res *dto.CustomerChangeAbandonResponseDTO, err error) {
	res = dto.NewCustomerChangeAbandonRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerChangeAbandon, reqData, res)
	return
}

// CustomerChangeUser 变更跟进人(支持批量)
func (c *Config) CustomerChangeUser(reqData *dto.CustomerChangeUserRequestDTO) (res *dto.CustomerChangeUserResponseDTO, err error) {
	res = dto.NewCustomerChangeUserRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerChangeUser, reqData, res)
	return
}

// CustomerFace 客户-获取客户头像
func (c *Config) CustomerFace(reqData *dto.CustomerFaceRequestDTO) (res *dto.CustomerFaceResponseDTO, err error) {
	res = dto.NewCustomerFaceRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerFace, reqData, res)
	return
}

// CustomerQueryLabel 查询客户标签(支持批量)
func (c *Config) CustomerQueryLabel(reqData *dto.CustomerQueryLabelRequestDTO) (res *dto.CustomerQueryLabelResponseDTO, err error) {
	res = dto.NewCustomerQueryLabelRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryLabel, reqData, res)
	return
}

// CustomerQueryList 查询客户列表
func (c *Config) CustomerQueryList(reqData *dto.CustomerQueryListRequestDTO) (res *dto.CustomerQueryListResponseDTO, err error) {
	res = dto.NewCustomerQueryListRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryList, reqData, res)
	return
}

// CustomerFileList 客户资料-文件列表查询
func (c *Config) CustomerFileList(reqData *dto.CustomerFileListRequestDTO) (res *dto.CustomerFileListResponseDTO, err error) {
	res = dto.NewCustomerFileListRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerFileList, reqData, res)
	return
}

// CustomerFolderList 客户资料-文件目录查询
func (c *Config) CustomerFolderList(reqData *dto.CustomerFolderListRequestDTO) (res *dto.CustomerFolderListResponseDTO, err error) {
	res = dto.NewCustomerFolderListRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerFolderList, reqData, res)
	return
}

// CustomerQueryExist 手机查询客户(判重)
func (c *Config) CustomerQueryExist(reqData *dto.CustomerQueryExistRequestDTO) (res *dto.CustomerQueryExistResponseDTO, err error) {
	res = dto.NewCustomerQueryExistRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryExist, reqData, res)
	return
}

// CustomerQuery 分页查询客户信息 (建议使用 查询客户列表 接口代替)
func (c *Config) CustomerQuery(reqData *dto.CustomerQueryRequestDTO) (res *dto.CustomerQueryResponseDTO, err error) {
	res = dto.NewCustomerQueryRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerQuery, reqData, res)
	return
}

// CustomerPreciseQuery 批量查询客户信息(建议使用 查询客户列表 接口代替)
func (c *Config) CustomerPreciseQuery(reqData *dto.CustomerPreciseQueryRequestDTO) (res *dto.CustomerPreciseQueryResponseDTO, err error) {
	res = dto.NewCustomerPreciseQueryRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerPreciseQuery, reqData, res)
	return
}

// CustomerQueryCustomer 导出客户
func (c *Config) CustomerQueryCustomer(reqData *dto.CustomerQueryCustomerRequestDTO) (res *dto.CustomerQueryCustomerResponseDTO, err error) {
	res = dto.NewCustomerQueryCustomerRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerQueryCustomer, reqData, res)
	return
}

// CustomerGetCustomerGroup 查询客户分组
func (c *Config) CustomerGetCustomerGroup(reqData *dto.CustomerGetCustomerGroupRequestDTO) (res *dto.CustomerGetCustomerGroupResponseDTO, err error) {
	res = dto.NewCustomerGetCustomerGroupRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerGetCustomerGroup, reqData, res)
	return
}

// CustomerDelcrms 客户-获取删除的客户
func (c *Config) CustomerDelcrms(reqData *dto.CustomerDelcrmsRequestDTO) (res *dto.CustomerDelcrmsResponseDTO, err error) {
	res = dto.NewCustomerDelcrmsRequestDTO()
	err = c.RequestByAutoResolve(common.CustomerDelcrms, reqData, res)
	return
}
