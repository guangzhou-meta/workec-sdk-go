package common

const (
	ECRequestContentType      = "Content-Type"
	ECRequestContentTypeValue = "application/json"
	ECRequestCorpId           = "X-Ec-Cid"
	ECRequestSign             = "X-Ec-Sign"
	ECRequestTimeStamp        = "X-Ec-TimeStamp"
)

const (
	ServerBaseUrlV1      = "https://open.workec.com/"
	ServerBaseUrlV2      = "https://open.workec.com/v2/"
	DefaultServerBaseUrl = ServerBaseUrlV2
)

type RoleChangeUserModifyType int

const (
	RoleChangeUserAdd    RoleChangeUserModifyType = 1
	RoleChangeUserDelete RoleChangeUserModifyType = 2
)

type CustomerCustomerCrmType int

const (
	CustomerSolo CustomerCustomerCrmType = 0
	CustomerCorp CustomerCustomerCrmType = 1
)

type CustomerCustomerGender int

const (
	CustomerNot  CustomerCustomerGender = 0
	CustomerGirl CustomerCustomerGender = 1
	CustomerMan  CustomerCustomerGender = 2
)

type CustomerCustomerLunar int

const (
	CustomercaLunarLendar        CustomerCustomerLunar = 0
	CustomercaLunarSolarCalendar CustomerCustomerLunar = 1
)

type CustomerQueryExistCrmType int

const (
	Personage  CustomerQueryExistCrmType = 0
	Enterprise CustomerQueryExistCrmType = 1
)
