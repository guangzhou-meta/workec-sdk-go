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
	CustomerNil  CustomerCustomerGender = 0
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

type TrajectoryinOutType int

const (
	CallIn  TrajectoryinOutType = 1
	CallOut TrajectoryinOutType = 2
)

type StatisticsSequential int

const (
	NotInclude StatisticsSequential = 0
	Include    StatisticsSequential = 1
)

type ContactLunar int

const (
	ContactNotLeapMonth ContactLunar = 0
	ContactLeapMonth    ContactLunar = 1
)

type ContactGender int

const (
	ContactNil  ContactGender = 0
	ContactGirl ContactGender = 1
	ContactMan  ContactGender = 2
)
