package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// SalesGetSalesFiles 获取订单附件
func (c *Config) SalesGetSalesFiles(reqData *dto.SalesGetSalesFilesRequestDTO) (res *dto.SalesGetSalesFilesResponseDTO, err error) {
	res = dto.NewSalesGetSalesFilesRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetSalesFiles, reqData, res)
	return
}

// SalesGetSalesFieldMapping 获取订单字段信息
func (c *Config) SalesGetSalesFieldMapping() (res *dto.SalesGetSalesFieldMappingResponseDTO, err error) {
	res = dto.NewSalesGetSalesFieldMappingRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetSalesFieldMapping, nil, res)
	return
}

// SalesUpdateSales 修改订单
func (c *Config) SalesUpdateSales(reqData *dto.SalesUpdateSalesRequestDTO) (res *dto.SalesUpdateSalesResponseDTO, err error) {
	res = dto.NewSalesUpdateSalesRequestDTO()
	err = c.RequestByAutoResolve(common.SalesUpdateSales, reqData, res)
	return
}

// SalesUpdateStatus 更新订单状态
func (c *Config) SalesUpdateStatus(reqData *dto.SalesUpdateStatusRequestDTO) (res *dto.SalesUpdateStatusResponseDTO, err error) {
	res = dto.NewSalesUpdateStatusRequestDTO()
	err = c.RequestByAutoResolve(common.SalesUpdateStatus, reqData, res)
	return
}

// SalesGetSales 查询订单列表
func (c *Config) SalesGetSales(reqData *dto.SalesGetSalesRequestDTO) (res *dto.SalesGetSalesResponseDTO, err error) {
	res = dto.NewSalesGetSalesRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetSales, reqData, res)
	return
}

// SalesGetSalesDetail 查询订单详情
func (c *Config) SalesGetSalesDetail(reqData *dto.SalesGetSalesDetailRequestDTO) (res *dto.SalesGetSalesDetailResponseDTO, err error) {
	res = dto.NewSalesGetSalesDetailRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetSalesDetail, reqData, res)
	return
}

// SalesDeleteProduct 删除产品信息
func (c *Config) SalesDeleteProduct(reqData *dto.SalesDeleteProductRequestDTO) (res *dto.SalesDeleteProductResponseDTO, err error) {
	res = dto.NewSalesDeleteProductRequestDTO()
	err = c.RequestByAutoResolve(common.SalesDeleteProduct, reqData, res)
	return
}

// SalesUpdateProduct 修改产品信息
func (c *Config) SalesUpdateProduct(reqData *dto.SalesUpdateProductRequestDTO) (res *dto.SalesUpdateProductResponseDTO, err error) {
	res = dto.NewSalesUpdateProductRequestDTO()
	err = c.RequestByAutoResolve(common.SalesUpdateProduct, reqData, res)
	return
}

// SalesAddProduct 新增产品
func (c *Config) SalesAddProduct(reqData *dto.SalesAddProductRequestDTO) (res *dto.SalesAddProductResponseDTO, err error) {
	res = dto.NewSalesAddProductRequestDTO()
	err = c.RequestByAutoResolve(common.SalesAddProduct, reqData, res)
	return
}

// SalesUpload 上传订单附件
func (c *Config) SalesUpload(reqData *dto.SalesUploadRequestDTO) (res *dto.SalesUploadResponseDTO, err error) {
	res = dto.NewSalesUploadRequestDTO()
	err = c.RequestByAutoResolve(common.SalesUpload, reqData, res)
	return
}

// SalesGetProductGroupList 获取产品分组
func (c *Config) SalesGetProductGroupList() (res *dto.SalesGetProductGroupListResponseDTO, err error) {
	res = dto.NewSalesGetProductGroupListRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetProductGroupList, nil, res)
	return
}

// SalesGetProductList 分页获取产品源数据
func (c *Config) SalesGetProductList(reqData *dto.SalesGetProductListRequestDTO) (res *dto.SalesGetProductListResponseDTO, err error) {
	res = dto.NewSalesGetProductListRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetProductList, reqData, res)
	return
}

// SalesAddSales 创建订单
func (c *Config) SalesAddSales(reqData *dto.SalesAddSalesRequestDTO) (res *dto.SalesAddSalesResponseDTO, err error) {
	res = dto.NewSalesAddSalesRequestDTO()
	err = c.RequestByAutoResolve(common.SalesAddSales, reqData, res)
	return
}

// SalesGetProductsDetail 订单产品明细查询
func (c *Config) SalesGetProductsDetail(reqData *dto.SalesGetProductsDetailRequestDTO) (res *dto.SalesGetProductsDetailResponseDTO, err error) {
	res = dto.NewSalesGetProductsDetailRequestDTO()
	err = c.RequestByAutoResolve(common.SalesGetProductsDetail, reqData, res)
	return
}
