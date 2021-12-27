package dto

type SalesGetSalesFilesRequestDTO struct {
	CrmId  int64 `json:"crmId"`
	SaleId int   `json:"saleId"`
	Type   int   `json:"type"`
}

type SalesGetSalesFilesResponseDTO struct {
	CommonDTO
	Data *SalesGetSalesFilesResponseDTOData `json:"data"`
}

type SalesGetSalesFilesResponseDTOData struct {
	CrmID          int64  `json:"crmID"`
	FolderId       int64  `json:"folderId"`
	CreateUserId   int64  `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
	Path           string `json:"path"`
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Size           int64  `json:"size"`
	UpdateTime     string `json:"updateTime"`
}

func NewSalesGetSalesFilesRequestDTO() *SalesGetSalesFilesResponseDTO {
	return &SalesGetSalesFilesResponseDTO{}
}

type SalesGetSalesFieldMappingResponseDTO struct {
	CommonDTO
	Data []*SalesGetSalesFieldMappingResponseDTOData `json:"data"`
}

type SalesGetSalesFieldMappingResponseDTOData struct {
	GroupId   int                                               `json:"groupId"`
	GroupName string                                            `json:"groupName"`
	Fields    []*SalesGetSalesFieldMappingResponseDTODataFields `json:"fields"`
}

type SalesGetSalesFieldMappingResponseDTODataFields struct {
	Id        int                                                     `json:"id"`
	Name      string                                                  `json:"name"`
	Type      int                                                     `json:"type"`
	Stype     int                                                     `json:"stype"`
	Isnotnull int                                                     `json:"isnotnull"`
	Status    int                                                     `json:"status"`
	Params    []*SalesGetSalesFieldMappingResponseDTODataFieldsParams `json:"params"`
}

type SalesGetSalesFieldMappingResponseDTODataFieldsParams struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func NewSalesGetSalesFieldMappingRequestDTO() *SalesGetSalesFieldMappingResponseDTO {
	return &SalesGetSalesFieldMappingResponseDTO{}
}

type SalesUpdateSalesRequestDTO struct {
}

type SalesUpdateStatusRequestDTO struct {
	SaleId    int64 `json:"saleId"`
	OptUserId int64 `json:"optUserId"`
	3 int
}

type SalesUpdateStatusResponseDTO struct {
	CommonDTO
	Data int `json:"data"`
}

func NewSalesUpdateStatusRequestDTO() *SalesUpdateStatusResponseDTO {
	return &SalesUpdateStatusResponseDTO{}
}

type SalesGetSalesRequestDTO struct {
	UserIds        string `json:"userIds"`
	CrmIds         string `json:"crmIds"`
	DeptIds        string `json:"deptIds"`
	CreatTime      string `json:"creatTime"`
	LastModifyTime string `json:"lastModifyTime"`
	SalesStatus    int    `json:"salesStatus"`
	PageNo         int    `json:"pageNo"`
}

type SalesGetSalesResponseDTO struct {
	CommonDTO
	Total     int                             `json:"total"`
	PageSize  int                             `json:"pageSize"`
	MaxPageNo int                             `json:"maxPageNo"`
	PageNo    int                             `json:"pageNo"`
	Data      []*SalesGetSalesResponseDTOData `json:"data"`
}

type SalesGetSalesResponseDTOData struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"userId"`
	UserName    string `json:"userName"`
	CreateTime  string `json:"createTime"`
	CHangeTime  string `json:"CHangeTime"`
	ProjectName string `json:"projectName"`
	Status      int    `json:"status"`
	CrmId       string `json:"crmId"`
	Name        string `json:"name"`
	ProductId   string `json:"productId"`
	Money       int64  `json:"money"`
	GuessTime   string `json:"guessTime"`
	GroupId     string `json:"groupId"`
	DealDate    string `json:"dealDate"`
	ProductDes  string `json:"productDes"`
	Remark      string `json:"remark"`
	ChangeUser  string `json:"changeUser"`
	DoUserid    string `json:"doUserid"`
}

func NewSalesGetSalesRequestDTO() *SalesGetSalesResponseDTO {
	return &SalesGetSalesResponseDTO{}
}

type SalesGetSalesDetailRequestDTO struct {
	SaleId int64 `json:"saleId"`
}

type SalesGetSalesDetailResponseDTO struct {
	CommonDTO
	Data []*SalesGetSalesDetailResponseDTOData `json:"data"`
}

type SalesGetSalesDetailResponseDTOData struct {
	IdName            int64                                                  `json:"idName"`
	UserId            int64                                                  `json:"userId"`
	UserName          string                                                 `json:"userName"`
	CreateTime        string                                                 `json:"createTime"`
	ChangeTime        string                                                 `json:"changeTime"`
	ProjectName       string                                                 `json:"projectName"`
	Status            int                                                    `json:"status"`
	CrmId             string                                                 `json:"crmId"`
	Name              string                                                 `json:"name"`
	ProductId         string                                                 `json:"productId"`
	Money             int64                                                  `json:"money"`
	GuessTime         string                                                 `json:"guessTime"`
	GorupID           string                                                 `json:"gorupID"`
	DealDate          string                                                 `json:"dealDate"`
	ProductDes        string                                                 `json:"productDes"`
	Remark            string                                                 `json:"remark"`
	ChangeUser        string                                                 `json:"changeUser"`
	DoUserid          string                                                 `json:"doUserid"`
	CustomFieldValues []*SalesGetSalesDetailResponseDTODataCustomFieldValues `json:"customFieldValues"`
}

type SalesGetSalesDetailResponseDTODataCustomFieldValues struct {
	FieldId    int    `json:"fieldId"`
	FieldValue string `json:"fieldValue"`
}

func NewSalesGetSalesDetailRequestDTO() *SalesGetSalesDetailResponseDTO {
	return &SalesGetSalesDetailResponseDTO{}
}

type SalesDeleteProductRequestDTO struct {
	OptUserId  int64 `json:"optUserId"`
	ProductIds []int `json:"productIds"`
}

type SalesDeleteProductResponseDTO struct {
	CommonDTO
	Data map[string]string
}

func NewSalesDeleteProductRequestDTO() *SalesDeleteProductResponseDTO {
	return &SalesDeleteProductResponseDTO{}
}

type SalesUpdateProductRequestDTO struct {
	OptUserID int64                                `json:"optUserID"`
	Product   *SalesUpdateProductRequestDTOproduct `json:"product"`
}

type SalesUpdateProductRequestDTOproduct struct {
	ProductId        int64                                               `json:"productId"`
	GroupId          int64                                               `json:"groupId"`
	ProductName      string                                              `json:"productName"`
	ProductUnit      string                                              `json:"productUnit"`
	Money            int64                                               `json:"money"`
	CostMoney        int64                                               `json:"costMoney"`
	OnOff            int                                                 `json:"onOff"`
	Specs            int                                                 `json:"specs"`
	SpecsRequestList []int                                               `json:"specsRequestList"`
	ProductSpecsDto  *SalesUpdateProductRequestDTOproductProductSpecsDto `json:"productSpecsDto"`
}

type SalesUpdateProductRequestDTOproductProductSpecsDto struct {
	ProductSpecsName  string  `json:"productSpecsName"`
	ProductSpecsMoney float64 `json:"productSpecsMoney"`
	CostMoney         float64 `json:"costMoney"`
	OnOff             int     `json:"onOff"`
}

type SalesUpdateProductResponseDTO struct {
	CommonDTO
	Data int64 `json:"data"`
}

func NewSalesUpdateProductRequestDTO() *SalesUpdateProductResponseDTO {
	return &SalesUpdateProductResponseDTO{}
}

type SalesAddProductRequestDTO struct {
	OptUserId int64                             `json:"optUserId"`
	Product   *SalesAddProductRequestDTOProduct `json:"product"`
}

type SalesAddProductRequestDTOProduct struct {
	ProductId        int64                                            `json:"productId"`
	GroupId          int64                                            `json:"groupId"`
	ProductName      string                                           `json:"productName"`
	ProductUnit      string                                           `json:"productUnit"`
	Money            float64                                          `json:"money"`
	CostMoney        float64                                          `json:"costMoney"`
	OnOff            int                                              `json:"onOff"`
	Specs            int                                              `json:"specs"`
	SpecsRequestList []int                                            `json:"specsRequestList"`
	ProductSpecsDto  *SalesAddProductRequestDTOProductProductSpecsDto `json:"productSpecsDto"`
}

type SalesAddProductRequestDTOProductProductSpecsDto struct {
	ProductSpecsName  string  `json:"productSpecsName"`
	ProductSpecsMoney float64 `json:"productSpecsMoney"`
	CostMoney         float64 `json:"costMoney"`
	OnOff             int     `json:"onOff"`
}

type SalesAddProductResponseDTO struct {
	CommonDTO
	Data int64 `json:"data"`
}

func NewSalesAddProductRequestDTO() *SalesAddProductResponseDTO {
	return &SalesAddProductResponseDTO{}
}

type SalesUploadRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	CrmId     int   `json:"crmId"`
	FolderId  int64 `json:"folderId"`
	SalesId   int64 `json:"salesId"`
}

type SalesUploadResponseDTO struct {
	CommonDTO
	Data map[string]string `json:"data"`
}

func NewSalesUploadRequestDTO() *SalesUploadResponseDTO {
	return &SalesUploadResponseDTO{}
}

type SalesGetProductGroupListResponseDTO struct {
	CommonDTO
	Data map[string]string `json:"data"`
}

func NewSalesGetProductGroupListRequestDTO() *SalesGetProductGroupListResponseDTO {
	return &SalesGetProductGroupListResponseDTO{}
}

type SalesGetProductListRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	PageNo    int64 `json:"pageNo"`
	PageSize  int64 `json:"pageSize"`
}

type SalesGetProductListResponseDTO struct {
	CommonDTO
	Data map[string]string `json:"data"`
}

func NewSalesGetProductListRequestDTO() *SalesGetProductListResponseDTO {
	return &SalesGetProductListResponseDTO{}
}

type SalesAddSalesRequestDTO struct {
	OptUserId         int64   `json:"optUserId"`
	CrmId             int64   `json:"crmId"`
	OrderIsSubject    string  `json:"1"`
	OrderStatus       int     `json:"3"`
	SystemOrderAmount float64 `json:"4"`
	2101                string  `json:"2101"`
	2109 strint `json:"2109"`
	Product *SalesAddSalesRequestDTOProduct `json:"product"`
}

type SalesAddSalesRequestDTOProduct struct {
	CommonDTO
	Data int64 `json:"data"`
}

func NewSalesAddSalesRequestDTO() *SalesAddSalesRequestDTOProduct {
	return &SalesAddSalesRequestDTOProduct{}
}

type SalesGetProductsDetailRequestDTO struct {
	UserId  string  `json:"userId"`
	SaleIds []int64 `json:"saleIds"`
}

type SalesGetProductsDetailResponseDTO struct {
	CommonDTO
	Data *SalesGetProductsDetailResponseDTOData `json:"data"`
}

type SalesGetProductsDetailResponseDTOData struct {
	Id               int64   `json:"id"`
	SaleId           int64   `json:"saleId"`
	SpecsId          int64   `json:"specsId"`
	ProductSpecsName string  `json:"productSpecsName"`
	Unit             string  `json:"unit"`
	ProductGroupId   int64   `json:"productGroupId"`
	ProductId        int64   `json:"productId"`
	ProductName      string  `json:"productName"`
	ProductNum       int     `json:"productNum"`
	OriginalMoney    float64 `json:"originalMoney"`
	SaleMoney        float64 `json:"saleMoney"`
	SaleDiscount     int     `json:"saleDiscount"`
	OnOff            int     `json:"onOff"`
	Del              int     `json:"del"`
}

func NewSalesGetProductsDetailRequestDTO() *SalesGetProductsDetailResponseDTO {
	return &SalesGetProductsDetailResponseDTO{}
}
