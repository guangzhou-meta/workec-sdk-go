package dto

import "github.com/guangzhou-meta/workec-sdk-go/common"

type CustomerAddCustomerRequestDTO struct {
	OptUserID int64                                       `json:"optUserID"`
	List      []*CustomerAddCustomerRequestCustomerDetail `json:"list"`
	Notify    bool                                        `json:"notify"`
	Repeat    bool                                        `json:"repeat"`
}

type CustomerAddCustomerRequestCustomerDetail struct {
	FollowUserId      int64                          `json:"followUserId"`
	CreateCrmType     common.CustomerCustomerCrmType `json:"createCrmType"`
	GroupID           int64                          `json:"groupID"`
	Name              string                         `json:"name"`
	Call              string                         `json:"call"`
	Gender            int                            `json:"gender"`
	Birthday          string                         `json:"birthday"`
	IsBirthdayLunar   int                            `json:"isBirthdayLunar"`
	BirthdayLunarLeap int                            `json:"birthdayLunarLeap"`
	Title             string                         `json:"title"`
	Qq                string                         `json:"qq"`
	Mobile            string                         `json:"mobile"`
	Phone             string                         `json:"phone"`
	Fax               string                         `json:"fax"`
	MobileCode        string                         `json:"mobileCode"`
	PhoneCode         string                         `json:"phoneCode"`
	FaxCode           string                         `json:"faxCode"`
	Email             string                         `json:"email"`
	Company           string                         `json:"company"`
	CompanyUrl        string                         `json:"companyUrl"`
	CompanyAddress    string                         `json:"companyAddress"`
	Memo              string                         `json:"memo"`
	Vocation          string                         `json:"vocation"`
	Channelld         int                            `json:"channelld"`
	Prefecture        string                         `json:"prefecture"`
	Wechat            string                         `json:"wechat"`
	Founddate         string                         `json:"founddate"`
	Fields            map[string]string              `json:"fields"`
}

type CustomerAddCustomerResponseDTO struct {
	CommonDTO
	Data *CustomerAddCustomerResponseData `json:"data"`
}

func NewCustomerAddCustomerResponseDTO() *CustomerAddCustomerResponseDTO {
	return &CustomerAddCustomerResponseDTO{}
}

type CustomerAddCustomerResponseData struct {
	FailureList   []*CustomerAddCustomerResponseDataFailureListItem   `json:"failureList"`
	SuccessIdList []*CustomerAddCustomerResponseDataSuccessIdListItem `json:"successIdList"`
}

type CustomerAddCustomerResponseDataFailureListItem struct {
	FailureCause string `json:"failureCause"`
	Index        int    `json:"index"`
}
type CustomerAddCustomerResponseDataSuccessIdListItem struct {
	CrmId int64 `json:"crmId"`
	Index int   `json:"index"`
}

type CustomerUpdateCustomerRequestDTO struct {
	OptUserId         int64                         `json:"optUserId"`
	Repeat            bool                          `json:"repeat"`
	CrmId             int64                         `json:"crmId"`
	Name              string                        `json:"name"`
	Call              string                        `json:"call"`
	Memo              string                        `json:"memo"`
	Gender            common.CustomerCustomerGender `json:"gender"`
	Title             string                        `json:"title"`
	Qq                string                        `json:"qq"`
	Mobile            string                        `json:"mobile"`
	Email             string                        `json:"email"`
	Wechat            string                        `json:"wechat"`
	Company           string                        `json:"company"`
	CompanyUrl        string                        `json:"companyUrl"`
	CompanyAddress    string                        `json:"companyAddress"`
	Vocation          string                        `json:"vocation"`
	Channelld         int                           `json:"channelld"`
	Prefecture        string                        `json:"prefecture"`
	Fields            map[string]string             `json:"fields"`
	GroupId           int64                         `json:"groupId"`
	PublicPondld      int64                         `json:"publicPondld"`
	Phone             string                        `json:"phone"`
	Fax               string                        `json:"fax"`
	MobileCode        string                        `json:"mobileCode"`
	PhoneCode         string                        `json:"phoneCode"`
	FaxCode           string                        `json:"faxCode"`
	Birthday          string                        `json:"birthday"`
	IsBirthDayLunar   common.CustomerCustomerLunar  `json:"isBirthDayLunar"`
	BirthdayLunarLeap common.CustomerCustomerLunar  `json:"birthdayLunarLeap"`
	FoundDate         string                        `json:"founddate"`
}

type CustomerBatchUpdateCustomerRequestDTO struct {
	OptUserId         int64                         `json:"optUserId"`
	Repeat            bool                          `json:"repeat"`
	CrmId             int64                         `json:"crmId"`
	GroupId           int64                         `json:"groupId"`
	PublicPondId      int64                         `json:"publicPondId"`
	Name              string                        `json:"name"`
	Call              string                        `json:"call"`
	Gender            common.CustomerCustomerGender `json:"gender"`
	Birthday          string                        `json:"birthday"`
	IsBirthdayLunar   int                           `json:"isBirthdayLunar"`
	BirthdayLunarLeap int                           `json:"birthdayLunarLeap"`
	Title             string                        `json:"title"`
	Qq                string                        `json:"qq"`
	Moblie            string                        `json:"moblie"`
	Phone             string                        `json:"phone"`
	Fax               string                        `json:"fax"`
	MobileCode        string                        `json:"mobileCode"`
	PhoneCode         string                        `json:"phoneCode"`
	FaxCode           string                        `json:"faxCode"`
	Email             string                        `json:"email"`
	Comany            string                        `json:"company"`
	ComanyUrl         string                        `json:"comanyUrl"`
	CompanyAddress    string                        `json:"companyAddress"`
	Memo              string                        `json:"memo"`
	Vocation          string                        `json:"vocation"`
	Channelld         int                           `json:"channelld"`
	Prefecture        string                        `json:"prefecture"`
	Wechat            string                        `json:"wechat"`
	FoundDate         string                        `json:"founddate"`
	FieLds            map[string]string             `json:"fieLds"`
}

type CustomerBatchUpdateCustomerResponseDTO struct {
	CommonDTO
	FailureList []int `json:"failureList"`
	SuccessList []int `json:"successList"`
}

func NewCustomerBatchUpdateCustomerResponseDTO() *CustomerBatchUpdateCustomerResponseDTO {
	return &CustomerBatchUpdateCustomerResponseDTO{}
}

type CustomerCombineRequestDTO struct {
	OptUserId         int64                         `json:"optUserId"`
	CrmIds            string                        `json:"crmIds"`
	FollowUserId      int64                         `json:"followUserId"`
	GroupId           int64                         `json:"groupId"`
	PublicPondId      int64                         `json:"publicPondId"`
	Name              string                        `json:"name"`
	Call              string                        `json:"call"`
	Gender            common.CustomerCustomerGender `json:"gender"`
	Birthday          string                        `json:"birthday"`
	IsBirthdayLunar   common.CustomerCustomerLunar  `json:"isBirthdayLunar"`
	BirthdayLunarLeap common.CustomerCustomerLunar  `json:"birthdayLunarLeap"`
	Title             string                        `json:"title"`
	Qq                string                        `json:"qq"`
	Mobile            string                        `json:"mobile"`
	Phone             string                        `json:"phone"`
	Fax               string                        `json:"fax"`
	MobileCode        string                        `json:"mobileCode"`
	PhoneCode         string                        `json:"phoneCode"`
	FaxCode           string                        `json:"faxCode"`
	Email             string                        `json:"email"`
	Company           string                        `json:"company"`
	CompanyUrl        string                        `json:"companyUrl"`
	CompanyAddress    string                        `json:"companyAddress"`
	Memo              string                        `json:"memo"`
	Vocation          string                        `json:"vocation"`
	Channelld         int                           `json:"channelld"`
	Prefecture        string                        `json:"prefecture"`
	Wechat            string                        `json:"wechat"`
	FieIds            map[string]string             `json:"fieIds"`
}

type CustomerCombineResponseDTO struct {
	CommonDTO
	data int64 `json:"data"`
}

func NewCustomerCombineResponseDTO() *CustomerCombineResponseDTO {
	return &CustomerCombineResponseDTO{}
}

type CustomerFileUploadRequestDTO struct {
	OptUserId int64  `json:"optUserId"`
	CrmId     string `json:"crmId"`
	FolderId  int    `json:"folderId"`
	File      []int  `json:"file"`
}

type CustomerFileUploadResponseDTO struct {
	CommonDTO
	Data *CustomerFileUploadResponseDTOData `json:"data"`
}

type CustomerFileUploadResponseDTOData struct {
	CorpId    int64                                                   `json:"corpId"`
	OptUserId int64                                                   `json:"optUserId"`
	CrmId     int64                                                   `json:"crmId"`
	FileList  *CustomerFileUploadResponseDTOCrmFileUploadItemFileInfo `json:"fileList"`
}

type CustomerFileUploadResponseDTOCrmFileUploadItemFileInfo struct {
	FileName string `json:"fileName"`
	Size     int64  `json:"size"`
}

func NewCustomerFileUploadResponseDTO() *CustomerFileUploadResponseDTO {
	return &CustomerFileUploadResponseDTO{}
}

type CustomerShareRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	CrmIds    []int `json:"crmIds"`
	UserIds   []int `json:"userIds"`
	Type      int   `json:"type"`
}

type CustomerShareResponseDTO struct {
	CommonDTO
	Data *CustomerShareResponseDTOData
}

type CustomerShareResponseDTOData struct {
	CrmId  int64  `json:"crmId"`
	UserId int64  `json:"userId"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func NewCustomerShareRequestDTO() *CustomerShareResponseDTO {
	return &CustomerShareResponseDTO{}
}

type CustomerChangeAbandonRequestDTO struct {
	OptUserId int64  `json:"optUserId"`
	CrmIds    string `json:"crmIds"`
	Type      int    `json:"type"`
}

type CustomerChangeAbandonResponseDTO struct {
	CommonDTO
	Data []*CustomerChangeAbandonResponseDTOData `json:"data"`
}

type CustomerChangeAbandonResponseDTOData struct {
	CommonDTO
	CrmId int64  `json:"crmId"`
	Name  string `json:"name"`
}

func NewCustomerChangeAbandonRequestDTO() *CustomerChangeAbandonResponseDTO {
	return &CustomerChangeAbandonResponseDTO{}
}

type CustomerChangeUserRequestDTO struct {
	FollowUserId int64  `json:"followUserId"`
	CrmIds       string `json:"crmIds"`
	OptUserId    int64  `json:"optUserId"`
}

type CustomerChangeUserResponseDTO struct {
	CommonDTO
	Data *CustomerChangeUserResponseDTOData `json:"data"`
}
type CustomerChangeUserResponseDTOData struct {
	CommonDTO
	CrmID int64  `json:"crmID"`
	Name  string `json:"name"`
}

func NewCustomerChangeUserRequestDTO() *CustomerChangeUserResponseDTO {
	return &CustomerChangeUserResponseDTO{}
}

type CustomerFaceRequestDTO struct {
	CrmIdList []int `json:"crmIdList"`
	PageSize  int   `json:"pageSize"`
	PageNo    int   `json:"pageNo"`
}
type CustomerFaceResponseDTO struct {
	CommonDTO
	Data *CustomerFaceRequestDTO `json:"data"`
}

type CustomerFaceRequestDTOData struct {
	CrmId int64  `json:"crmId"`
	Face  string `json:"face"`
	Name  string `json:"name"`
}

func NewCustomerFaceRequestDTO() *CustomerFaceResponseDTO {
	return &CustomerFaceResponseDTO{}
}

type CustomerQueryLabelRequestDTO struct {
	CrmIds string `json:"crmIds"`
}

type CustomerQueryLabelResponseDTO struct {
	CommonDTO
	Data []*CustomerQueryLabelResponseDTODdata `json:"data"`
}

type CustomerQueryLabelResponseDTODdata struct {
	CrmId    int64  `json:"crmId"`
	Labellds string `json:"labellds"`
}

func NewCustomerQueryLabelRequestDTO() *CustomerQueryLabelResponseDTO {
	return &CustomerQueryLabelResponseDTO{}
}

type TimeRangeBaseVO struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type CustomerOrderBy struct {
	SortField *CustomerOrderBySortField
	SortType  *CustomerOrderBySortType
}
type CustomerOrderBySortField struct {
	ModifyTime   string `json:"modifyTime"`
	ContactTime  string `json:"contactTime"`
	CreateTime   string `json:"createTime"`
	CrmId        string `json:"crmId"`
	FollowUserId string `json:"followUserId"`
	Channel      string `json:"channel"`
}

type CustomerOrderBySortType struct {
	Asc  string `json:"asc"`
	Desc string `json:"desc"`
}

type CustomerQueryListRequestDTO struct {
	Mobile           string                        `json:"mobile"`
	Email            string                        `json:"email"`
	Step             []int                         `json:"step"`
	Labellds         []int                         `json:"labellds"`
	FollowUserId     []int64                       `json:"followUserId"`
	PubilcPondId     []int64                       `json:"pubilcPondId"`
	ModifyTime       *TimeRangeBaseVO              `json:"modifyTime"`
	ContactTime      *TimeRangeBaseVO              `json:"contactTime"`
	CreateTime       *TimeRangeBaseVO              `json:"createTime"`
	CrmIds           []int64                       `json:"crmIds"`
	Gender           common.CustomerCustomerGender `json:"gender"`
	Fax              string                        `json:"fax"`
	Msn              string                        `json:"msn"`
	Qq               string                        `json:"qq"`
	Phone            string                        `json:"phone"`
	Name             string                        `json:"name"`
	Memo             string                        `json:"memo"`
	Company          string                        `json:"company"`
	EmailNums        int                           `json:"emailNums"`
	Channel          []int64                       `json:"channel"`
	Shared           bool                          `json:"shared"`
	ApiAdd           bool                          `json:"apiAdd"`
	CrmType          int                           `json:"crmType"`
	ShareUserId      []int64                       `json:"shareUserId"`
	CreateUserIdList []int64                       `json:"createUserIdList"`
	PageSize         int                           `json:"pageSize"`
	PageNo           int                           `json:"pageNo"`
	Includes         string                        `json:"includes"`
	OrderBy          *CustomerOrderBy              `json:"orderBy"`
}

type CustomerQueryListResponseDTO struct {
	CommonDTO
	Data map[*CustomerQueryListResponseDTOData]string `json:"data"`
}

type CustomerQueryListResponseDTOData struct {
	List []*CustomerQueryListResponseDTODataList `json:"list"`
	Page *CustomerQueryListResponseDTODataPage   `json:"page"`
}

type CustomerQueryListResponseDTODataList struct {
	PageSize  int   `json:"pageSize"`
	PageNo    int   `json:"pageNo"`
	Total     int64 `json:"total"`
	MaxPageNo int   `json:"maxPageNo"`
}

type CustomerQueryListResponseDTODataPage struct {
	FollowUserId     int64             `json:"followUserId"`
	GroupId          int64             `json:"groupId"`
	Name             string            `json:"name"`
	Call             string            `json:"call"`
	Gender           string            `json:"gender"`
	Birthday         string            `json:"birthday"`
	IsLunarBirthday  int               `json:"isLunarBirthday"`
	Title            string            `json:"title"`
	Qq               string            `json:"qq"`
	Mobile           string            `json:"mobile"`
	Phone            string            `json:"phone"`
	Fax              string            `json:"fax"`
	Email            string            `json:"email"`
	Company          string            `json:"company"`
	CompanyUrl       string            `json:"companyUrl"`
	CompanyAddress   string            `json:"companyAddress"`
	Memo             string            `json:"memo"`
	Vocation         string            `json:"vocation"`
	Channel          string            `json:"channel"`
	Prefecture       string            `json:"prefecture"`
	FieldInfos       map[string]string `json:"fieldInfos"`
	CrmId            string            `json:"crmId"`
	ModifyTime       string            `json:"modifyTime"`
	ContactTime      string            `json:"contactTime"`
	CreateTime       string            `json:"createTime"`
	LastFollowUserId string            `json:"lastFollowUserId"`
	Step             int               `json:"step"`
	CreateUserId     int64             `json:"createUserId"`
	Wechat           string            `json:"wechat"`
	Wechats          []string          `json:"wechats"`
	Emails           []string          `json:"emails"`
	Mobiles          []string          `json:"mobiles"`
	StorageTime      string            `json:"storageTime"`
	PublicPondId     int64             `json:"publicPondId"`
	ApiAdd           bool              `json:"apiAdd"`
	CrmType          int               `json:"crmType"`
	ShareUserId      int64             `json:"shareUserId"`
	LastContactTime  *CustomerQueryListResponseDTODataPageLastContactTime
	Stars            int `json:"stars"`
}

type CustomerQueryListResponseDTODataPageLastContactTime struct {
	Time string `json:"time"`
	Type int    `json:"type"`
}

func NewCustomerQueryListRequestDTO() *CustomerQueryListResponseDTO {
	return &CustomerQueryListResponseDTO{}
}

type CustomerFileListRequestDTO struct {
	CrmIds   int64  `json:"crmIds"`
	FolderId string `json:"folderId"`
}

type CustomerFileListResponseDTO struct {
	CommonDTO
	Data *CustomerFileListResponseDTOData `json:"data"`
}

type CustomerFileListResponseDTOData struct {
	CrmId          int64  `json:"crmId"`
	FolderId       int64  `json:"folderId"`
	CreateUserId   int64  `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
	Path           string `json:"path"`
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Size           int64  `json:"size"`
	UpdateTime     string `json:"updateTime"`
}

func NewCustomerFileListRequestDTO() *CustomerFileListResponseDTO {
	return &CustomerFileListResponseDTO{}
}

type CustomerFolderListRequestDTO struct {
	CrmIds int64 `json:"crmIds"`
}

type CustomerFolderListResponseDTO struct {
	CommonDTO
	Data *CustomerFolderListResponseDTOCata `json:"data"`
}

type CustomerFolderListResponseDTOCata struct {
	CrmId          int64  `json:"crmId"`
	CreateUserId   int64  `json:"createUserId"`
	CreateUserName string `json:"createUserName"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Size           int64  `json:"size"`
	UpdateTime     string `json:"updateTime"`
}

func NewCustomerFolderListRequestDTO() *CustomerFolderListResponseDTO {
	return &CustomerFolderListResponseDTO{}
}

type CustomerQueryExistRequestDTO struct {
	Mobile           string   `json:"mobile"`
	MaxNumsPreMobile int      `json:"maxNumsPreMobile"`
	Includes         []string `json:"includes"`
}

type CustomerQueryExistResponseDTO struct {
	CommonDTO
	Data CustomerQueryExistResponseDTOData `json:"data"`
}

type CustomerQueryExistResponseDTOData struct {
	Num    int                                    `json:"num"`
	Mobile string                                 `json:"mobile"`
	List   *CustomerQueryExistResponseDTODatalist `json:"list"`
}

type CustomerQueryExistResponseDTODatalist struct {
	CrmId            string                                                `json:"crmId"`
	FieIdId          int                                                   `json:"fieIdId"`
	FollowUserId     int64                                                 `json:"followUserId"`
	GroupId          int64                                                 `json:"groupId"`
	Name             string                                                `json:"name"`
	Call             string                                                `json:"call"`
	Gender           common.CustomerCustomerGender                         `json:"gender"`
	Title            string                                                `json:"title"`
	Qq               string                                                `json:"qq"`
	Phone            string                                                `json:"phone"`
	Fax              string                                                `json:"fax"`
	Email            string                                                `json:"email"`
	Company          string                                                `json:"company"`
	CompanyUrl       string                                                `json:"companyUrl"`
	CompanyAddress   string                                                `json:"companyAddress"`
	Memo             string                                                `json:"Memo"`
	Vocation         string                                                `json:"vocation"`
	Channel          int                                                   `json:"channel"`
	Prefecture       string                                                `json:"prefecture"`
	ModifyTime       string                                                `json:"modifyTime"`
	ContactTime      string                                                `json:"contactTime"`
	CreateTime       string                                                `json:"createTime"`
	LastFollowUserId int64                                                 `json:"lastFollowUserId"`
	Step             int                                                   `json:"step"`
	CreateUserId     int64                                                 `json:"createUserId"`
	Wechat           string                                                `json:"wechat"`
	Wechats          string                                                `json:"wechats"`
	Emails           string                                                `json:"emails"`
	Mobiles          string                                                `json:"mobiles"`
	StorageTime      string                                                `json:"storageTime"`
	PublicPondId     int64                                                 `json:"publicPondId"`
	ApiAdd           bool                                                  `json:"apiAdd"`
	CrmType          common.CustomerQueryExistCrmType                      `json:"crmType"`
	ShareUserId      int64                                                 `json:"shareUserId"`
	LastContactTime  *CustomerQueryExistResponseDTODatalistLastContactTime `json:"lastContactTime"`
}

type CustomerQueryExistResponseDTODatalistLastContactTime struct {
	Time string `json:"time"`
	Type int    `json:"type"`
}

func NewCustomerQueryExistRequestDTO() *CustomerQueryExistResponseDTO {
	return &CustomerQueryExistResponseDTO{}
}

type CustomerQueryRequestDTO struct {
	Mobile       string                           `json:"mobile"`
	Email        string                           `json:"email"`
	Step         int                              `json:"step"`
	LabelIds     string                           `json:"labelIds"`
	FollowUserId int64                            `json:"followUserId"`
	PublicPondId int64                            `json:"publicPondId"`
	ModifyTime   *TimeRangeBaseVO                 `json:"modifyTime"`
	ContactTime  *TimeRangeBaseVO                 `json:"contactTime"`
	CreateTime   *TimeRangeBaseVO                 `json:"createTime"`
	OrderBy      *CustomerQueryRequestDTOOrderBy  `json:"orderBy"`
	PageInfo     *CustomerQueryRequestDTOPageInfo `json:"pageInfo"`
}

type CustomerQueryRequestDTOOrderBy struct {
	SortField string `json:"sortField"`
	SortType  string `json:"sortType"`
}

type CustomerQueryRequestDTOPageInfo struct {
	PageSize int `json:"pageSize"`
	PageNo   int `json:"pageNo"`
}

type CustomerQueryResponseDTO struct {
	CommonDTO
	Data *CustomerQueryResponseDTOData `json:"data"`
}

type CustomerQueryResponseDTOData struct {
	CustomerInfoList *CustomerQueryResponseDTODataCustomerInfoList `json:"customerInfoList"`
	PageInfo         *CustomerQueryResponseDTODataPageInfo         `json:"pageInfo"`
}

type CustomerQueryResponseDTODataPageInfo struct {
	PageSize  int   `json:"pageSize"`
	PageNo    int   `json:"pageNo"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"totalPage"`
}

type CustomerQueryResponseDTODataCustomerInfoList struct {
	FollowUserId    int64                         `json:"followUserId"`
	GroupId         int64                         `json:"groupId"`
	Name            string                        `json:"name"`
	Call            string                        `json:"call"`
	Gender          common.CustomerCustomerGender `json:"gender"`
	Birthday        string                        `json:"birthday"`
	IsLunarBirthday common.CustomerCustomerLunar  `json:"isLunarBirthday"`
	Title           string                        `json:"title"`
	Qq              string                        `json:"qq"`
	Mobile          string                        `json:"mobile"`
	Phone           string                        `json:"phone"`
	Fax             string                        `json:"fax"`
	Email           string                        `json:"email"`
	Company         string                        `json:"company"`
	CompanyUrl      string                        `json:"companyUrl"`
	CompanyAddress  string                        `json:"companyAddress"`
	Memo            string                        `json:"memo"`
	Vocation        string                        `json:"vocation"`
	Channel         int                           `json:"channel"`
	Prefecture      string                        `json:"prefecture"`
	FieldInfos      map[string]string             `json:"fieldInfos"`
	CrmIs           string                        `json:"crmIs"`
	ModifyTime      *TimeRangeBaseVO              `json:"modifyTime"`
	ContactTime     *TimeRangeBaseVO              `json:"contactTime"`
	CreateTime      *TimeRangeBaseVO              `json:"createTime"`
	Step            int                           `json:"step"`
}

func NewCustomerQueryRequestDTO() *CustomerQueryResponseDTO {
	return &CustomerQueryResponseDTO{}
}

type CustomerPreciseQueryRequestDTO struct {
	CrmIds string `json:"crmIds"`
}

type CustomerPreciseQueryResponseDTO struct {
	CommonDTO
	Data *CustomerPreciseQueryResponseDTOData `json:"data"`
}

type CustomerPreciseQueryResponseDTOData struct {
	FollowUserId    int64                         `json:"followUserId"`
	GroupId         int64                         `json:"groupId"`
	Name            string                        `json:"name"`
	Call            string                        `json:"call"`
	Gender          common.CustomerCustomerGender `json:"gender"`
	Birthday        string                        `json:"birthday"`
	IsLunarBirthday common.CustomerCustomerLunar  `json:"isLunarBirthday"`
	Title           string                        `json:"title"`
	Qq              string                        `json:"qq"`
	Mobile          string                        `json:"mobile"`
	Phone           string                        `json:"phone"`
	Fax             string                        `json:"fax"`
	Email           string                        `json:"email"`
	Company         string                        `json:"company"`
	CompanyUrl      string                        `json:"companyUrl"`
	CompanyAddress  string                        `json:"companyAddress"`
	Memo            string                        `json:"memo"`
	Vocation        string                        `json:"vocation"`
	Channel         int                           `json:"channel"`
	Prefecture      string                        `json:"prefecture"`
	FieldInfos      map[string]string             `json:"fieldInfos"`
	CrmId           string                        `json:"crmId"`
	ModifyTime      string                        `json:"modifyTime"`
	ContactTime     string                        `json:"contactTime"`
	CreateTime      int64                         `json:"createTime"`
	Step            int                           `json:"step"`
}

func NewCustomerPreciseQueryRequestDTO() *CustomerPreciseQueryResponseDTO {
	return &CustomerPreciseQueryResponseDTO{}
}

type CustomerQueryCustomerRequestDTO struct {
	Mobile           string                                     `json:"mobile"`
	Email            string                                     `json:"email"`
	Step             []int                                      `json:"step"`
	LabelIds         []int                                      `json:"labelIds"`
	FollowUserId     []int64                                    `json:"followUserId"`
	PublicPondId     []int64                                    `json:"publicPondId"`
	ModifyTime       *TimeRangeBaseVO                           `json:"modifyTime"`
	ContactTime      *TimeRangeBaseVO                           `json:"contactTime"`
	CreateTime       *TimeRangeBaseVO                           `json:"createTime"`
	CrmIds           int64                                      `json:"crmIds"`
	Gender           common.CustomerCustomerGender              `json:"gender"`
	Fax              string                                     `json:"fax"`
	Msn              string                                     `json:"msn"`
	Qq               string                                     `json:"qq"`
	LastContactTime  string                                     `json:"lastContactTime"`
	Phone            string                                     `json:"phone"`
	Name             string                                     `json:"name"`
	Memo             string                                     `json:"memo"`
	Company          string                                     `json:"company"`
	EmailNums        int                                        `json:"emailNums"`
	MobileNums       int                                        `json:"mobileNums"`
	Channel          []int64                                    `json:"channel"`
	Shared           bool                                       `json:"shared"`
	ApiAdd           bool                                       `json:"apiAdd"`
	CrmType          common.CustomerQueryExistCrmType           `json:"crmType"`
	ShareUserId      []int64                                    `json:"shareUserId"`
	CreateUserIdList []int64                                    `json:"createUserIdList"`
	PageSize         int                                        `json:"pageSize"`
	PageNo           int                                        `json:"pageNo"`
	Includes         []string                                   `json:"includes"`
	OrderBy          *CustomerQueryCustomerRequestDTOSortBaseVO `json:"orderBy"`
	Sort             string                                     `json:"orderBy"`
}

type CustomerQueryCustomerRequestDTOSortBaseVO struct {
	SortField *CustomerQueryCustomerRequestDTOSortBaseVOSortField
	SortType  *CustomerQueryCustomerRequestDTOSortBaseVOSortType
}
type CustomerQueryCustomerRequestDTOSortBaseVOSortField struct {
	ModifyTime   string `json:"modifyTime"`
	ContactTime  string `json:"contactTime"`
	CreateTime   string `json:"createTime"`
	CrmId        string `json:"crmId"`
	FollowUserId string `json:"followUserId"`
	Channel      string `json:"channel"`
}
type CustomerQueryCustomerRequestDTOSortBaseVOSortType struct {
	Asc  string `json:"asc"`
	Desc string `json:"desc"`
}

type CustomerQueryCustomerResponseDTO struct {
	CommonDTO
	Data *CustomerQueryCustomerResponseDTOData `json:"data"`
	Sort []string                              `json:"sort"`
}

type CustomerQueryCustomerResponseDTOData struct {
	List []*CustomerQueryCustomerResponseDTODataList `json:"list"`
	Page *CustomerQueryCustomerResponseDTODataPage   `json:"page"`
}

type CustomerQueryCustomerResponseDTODataList struct {
	FollowUserId int64                         `json:"followUserId"`
	GroupId      int64                         `json:"groupId"`
	Name         string                        `json:"name"`
	Call         string                        `json:"call"`
	Gender       common.CustomerCustomerGender `json:"gender"`
	Birthday     string                        `json:"birthday"`
}

type CustomerQueryCustomerResponseDTODataPage struct {
	PageSize  int   `json:"pageSize"`
	PageNo    int   `json:"pageNo"`
	Total     int64 `json:"total"`
	MaxPageNo int   `json:"maxPageNo"`
}

func NewCustomerQueryCustomerRequestDTO() *CustomerQueryCustomerResponseDTO {
	return &CustomerQueryCustomerResponseDTO{}
}

type CustomerGetCustomerGroupRequestDTO struct {
	UserId int64
}

type CustomerGetCustomerGroupResponseDTO struct {
	CommonDTO
	Data *CustomerGetCustomerGroupResponseDTOData `json:"data"`
}

type CustomerGetCustomerGroupResponseDTOData struct {
	GroupId   string `json:"groupId"`
	GroupName string `json:"groupName"`
}

func NewCustomerGetCustomerGroupRequestDTO() *CustomerGetCustomerGroupResponseDTO {
	return &CustomerGetCustomerGroupResponseDTO{}
}

type CustomerDelcrmsRequestDTO struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	LastId    string `json:"lastId"`
}

type CustomerDelcrmsResponseDTO struct {
	ErrCode int                             `json:"errCode"`
	ErrMsg  string                          `json:"errMsg"`
	Data    *CustomerDelcrmsResponseDTOData `json:"data"`
}

type CustomerDelcrmsResponseDTOData struct {
	CrmId   int64  `json:"crmId"`
	DelTime string `json:"delTime"`
	Id      string `json:"id"`
}

func NewCustomerDelcrmsRequestDTO() *CustomerDelcrmsResponseDTO {
	return &CustomerDelcrmsResponseDTO{}
}

type CustomerGetTrajectoryRequestDTO struct {
	CrmIds         string                               `json:"crmIds"`
	Date           *CustomerGetTrajectoryRequestDTODate `json:"date"`
	LastId         int64                                `json:"lastId"`
	LastTime       string                               `json:"lastTime"`
	PageSize       int                                  `json:"pageSize"`
	TrajectoryType int                                  `json:"trajectoryType"`
}

type CustomerGetTrajectoryRequestDTODate struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type CustomerGetTrajectoryResponseDTO struct {
	CommonDTO
	Data *CustomerGetTrajectoryResponseDTODate
}

type CustomerGetTrajectoryResponseDTODate struct {
	NextPageDTO    *CustomerGetTrajectoryResponseDTODatetrajectoryList
	TrajectoryList *[]CustomerGetTrajectoryResponseDTODatetrajectoryList
}
type CustomerGetTrajectoryResponseDTODatenextPageDTO struct {
	HasNextPage  int    `json:"hasNextPage"`
	NextLastId   int64  `json:"nextLastId"`
	NextLastTime string `json:"nextLastTime"`
	PageSize     int    `json:"pageSize"`
}

type CustomerGetTrajectoryResponseDTODatetrajectoryList struct {
	Content         string `json:"content"`
	CreateTime      int64  `json:"createTime"`
	CrmId           int64  `json:"crmId"`
	ReceiveUserIds  int64  `json:"receiveUserIds"`
	TrajectoryId    int64  `json:"trajectoryId"`
	TrajectoryTypeT int64  `json:"trajectoryTypeT"`
	UserId          int64  `json:"userId"`
}

func NewCustomerGetTrajectoryRequestDTO() *CustomerGetTrajectoryResponseDTO {
	return &CustomerGetTrajectoryResponseDTO{}
}

type CustomerGetCrmVisitDetailsRequestDTO struct {
	UserIds   []int64 `json:"userIds"`
	DeptIds   []int64 `json:"deptIds"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	PageNo    int     `json:"pageNo"`
	PageSize  int     `json:"pageSize"`
}

type CustomerGetCrmVisitDetailsResponseDTO struct {
	CommonDTO
	Data *CustomerGetCrmVisitDetailsResponseDTOData `json:"data"`
}
type CustomerGetCrmVisitDetailsResponseDTOData struct {
	UserId      int64   `json:"userId"`
	UserName    string  `json:"userName"`
	CrmId       int64   `json:"crmId"`
	CrmName     string  `json:"crmName"`
	CompanyId   int64   `json:"companyId"`
	CompanyName string  `json:"companyName"`
	Time        string  `json:"time"`
	Address     string  `json:"address"`
	LongItUde   float64 `json:"longItude"`
	Latitude    float64 `json:"latitude"`
}

func NewCustomerGetCrmVisitDetailsRequestDTO() *CustomerGetCrmVisitDetailsResponseDTO {
	return &CustomerGetCrmVisitDetailsResponseDTO{}
}

type CustomerGetChannelSourceResponseDTO struct {
	CommonDTO
	Data *CustomerGetChannelSourceResponseDTOData `json:"data"`
}

type CustomerGetChannelSourceResponseDTOData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	UserId     int64  `json:"userId"`
	CreateTime int64  `json:"createTime"`
}

func NewCustomerGetChannelSourceRequestDTO() *CustomerGetChannelSourceResponseDTO {
	return &CustomerGetChannelSourceResponseDTO{}
}

type CustomerGetCasCadeFieldMappingRequestDTO struct {
	FieldIds []int `json:"fieldIds"`
	LastId   int   `json:"lastId"`
}

type CustomerGetCasCadeFieldMappingResponseDTO struct {
	CommonDTO
	Data *CustomerGetCasCadeFieldMappingResponseDTOData `json:"data"`
}

type CustomerGetCasCadeFieldMappingResponseDTOData struct {
	NextPageDTO *CustomerGetCasCadeFieldMappingResponseDTODataNextPageDTO `json:"nextPageDTO"`
	FieldParams *CustomerGetCasCadeFieldMappingResponseDTODataFieldParams `json:"fieldParams"`
}

type CustomerGetCasCadeFieldMappingResponseDTODataNextPageDTO struct {
	PageSize    int `json:"pageSize"`
	HasNextPage int `json:"hasNextPage"`
	NextLastId  int `json:"nextLastId"`
}
type CustomerGetCasCadeFieldMappingResponseDTODataFieldParams struct {
	FieldId       int    `json:"fieldId"`
	ParamId       int    `json:"paramId"`
	ParentParamId int    `json:"parentParamId"`
	ParamName     string `json:"paramName"`
	Sort          int    `json:"sort"`
}

func NewCustomerGetCasCadeFieldMappingRequestDTO() *CustomerGetCasCadeFieldMappingResponseDTO {
	return &CustomerGetCasCadeFieldMappingResponseDTO{}
}
