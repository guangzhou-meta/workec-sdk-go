package common

type RequestMethod int

const (
	POST RequestMethod = iota
	GET
)

type ApiModel struct {
	Path   string
	Method RequestMethod
}

func NewApiModel(path string, method RequestMethod) *ApiModel {
	return &ApiModel{
		Path:   path,
		Method: method,
	}
}

// org
var (
	OrgDeptCreate = NewApiModel("org/dept/create", POST)
	OrgDeptEdit   = NewApiModel("org/dept/edit", POST)
	OrgDeptDelete = NewApiModel("org/dept/delete", POST)
	OrgStructInfo = NewApiModel("org/struct/info", GET)

	OrgUserCreate           = NewApiModel("org/user/create", POST)
	OrgUserDelete           = NewApiModel("org/user/delete", POST)
	OrgUserUpdateUser       = NewApiModel("org/user/updateUser", POST)
	OrgUserOnOff            = NewApiModel("org/user/onoff", POST)
	OrgUserFindUserInfoById = NewApiModel("org/user/findUserInfoById", POST)
)

// role
var (
	RoleAdd              = NewApiModel("role/add", POST)
	RoleDelete           = NewApiModel("role/delete", POST)
	RoleChangeName       = NewApiModel("role/changeName", POST)
	RoleChangePrivileges = NewApiModel("role/changePrivileges", POST)
	RoleChangeUsers      = NewApiModel("role/changeUsers", POST)
	RoleList             = NewApiModel("role/list", POST)
)

// customer
var (
	CustomerAddCustomer            = NewApiModel("customer/addCustomer", POST)
	CustomerUpdateCustomer         = NewApiModel("customer/updateCustomer", POST)
	CustomerBatchUpdateCustomer    = NewApiModel("customer/batchUpdateCustomer", POST)
	CustomerCombine                = NewApiModel("customer/combine", POST)
	CustomerFileUpload             = NewApiModel("customer/file/upload", POST)
	CustomerShare                  = NewApiModel("customer/share", POST)
	CustomerChangeAbandon          = NewApiModel("customer/change/abandon", POST)
	CustomerChangeUser             = NewApiModel("customer/change/user", POST)
	CustomerFace                   = NewApiModel("customer/face", POST)
	CustomerQueryLabel             = NewApiModel("customer/queryLabel", POST)
	CustomerQueryList              = NewApiModel("customer/queryList", POST)
	CustomerFileList               = NewApiModel("customer/file/list", POST)
	CustomerFolderList             = NewApiModel("customer/folder/list", POST)
	CustomerQueryExist             = NewApiModel("customer/queryExist", POST)
	CustomerQuery                  = NewApiModel("customer/query", POST)
	CustomerPreciseQuery           = NewApiModel("customer/preciseQuery", POST)
	CustomerQueryCustomer          = NewApiModel("customer/queryCustomer", POST)
	CustomerGetCustomerGroup       = NewApiModel("customer/getCustomerGroup", POST)
	CustomerDelcrms                = NewApiModel("customer/delcrms", POST)
	CustomerGetTrajectory          = NewApiModel("customer/getTrajectory", POST)
	CustomerGetCrmVisitDetails     = NewApiModel("customer/getCrmVisitDetails", POST)
	CustomerGetChannelSource       = NewApiModel("customer/getChannelSource", GET)
	CustomerGetCasCadeFieldMapping = NewApiModel("customer/getCasCadeFieldMapping", POST)
)

// special
var (
	SpecialSelect = NewApiModel("special/select", POST)
)

// step
var (
	StepUpdate = NewApiModel("step/update", POST)
)

// record
var (
	RecordCall                  = NewApiModel("record/call", POST)
	RecordGetFreeStatusUid      = NewApiModel("record/getFreeStatusUid", POST)
	RecordSmsRecord             = NewApiModel("record/smsRecord", POST)
	RecordTelRecord             = NewApiModel("record/telRecord", POST)
	RecordSendSmsHistory        = NewApiModel("record/sendSmsHistory", POST)
	RecordTelRecordHistoryQuery = NewApiModel("record/telRecordHistoryQuery", POST)
	RecordAddTelRecord          = NewApiModel("record/addTelRecord", POST)
)

// trajectory
var (
	TrajectoryFindUserTrajectory        = NewApiModel("trajectory/findUserTrajectory", POST)
	TrajectorySaveUserTrajectory        = NewApiModel("trajectory/saveUserTrajectory", POST)
	TrajectoryFindHistoryUserTrajectory = NewApiModel("TrajectorySaveUserTrajectory", POST)
)

// robot
var (
	RobotAddTask       = NewApiModel("robot/addtask", POST)
	RobotAddTaskRecord = NewApiModel("robot/addtaskrecord", POST)
	RobotUpDateTask    = NewApiModel("robot/updatetask", POST)
)

// asynchronization
var (
	AsyncHroniZationCreate = NewApiModel("asynchronization/create", POST)
	AsynchroniZationQuery  = NewApiModel("asynchronization/query", POST)
)

// config
var (
	ConfigGetBookFieldMapping = NewApiModel("config/getBookFieldMapping", GET)
	ConfigGetFieldMapping     = NewApiModel("config/getFieldMapping", GET)
	ConfigGetPubicPond        = NewApiModel("config/getPubicPond", GET)
	ConfigGetStages           = NewApiModel("config/getStages", POST)
)

// statistics
var (
	StatisticsDigitalMapPhone                  = NewApiModel("statistics/digitalMap/phone", POST)
	StatisticsLineGraphPhone                   = NewApiModel("statistics/lineGraph/phone", POST)
	StatisticsDigitalMapWorkeffic              = NewApiModel("statistics/digitalMap/workeffic", POST)
	StatisticsHistogramWorkeffic               = NewApiModel("statistics/histogram/workeffic", POST)
	StatisticsDigitalMapTag                    = NewApiModel("statistics/digitalMap/tag", POST)
	StatisticsHistogramTag                     = NewApiModel("statistics/histogram/tag", POST)
	StatisticsDigitalMapCrmQuantity            = NewApiModel("statistics/digitalMap/crmQuantity", POST)
	StatisticsCrmStatsWueryStepCountByChannel  = NewApiModel("statistics/crmStats/queryStepCountByChannel", POST)
	StatisticsCrmStatsGetTopStepCountByChannel = NewApiModel("statistics/crmStats/getTopStepCountByChannel", POST)
	StatisticsCrmStatsGetStepCount             = NewApiModel("/statistics/crmStats/getStepCount", POST)
	StatisticsCrmStatsGroupbyUserIds           = NewApiModel("statistics/crmStats/groupbyUserIds", POST)
)

// plan
var (
	PlanGetPlanTemplate = NewApiModel("plan/getPlanTemplate", GET)
)

// sales
var (
	SalesGetSalesFiles        = NewApiModel("sales/getSalesFiles", POST)
	SalesGetSalesFieldMapping = NewApiModel("sales/getSalesFieldMapping", GET)
	SalesUpdateSales          = NewApiModel("SalesGetSalesFieldMapping", POST)
	SalesUpdateStatus         = NewApiModel("sales/updateStatus", POST)
	SalesGetSales             = NewApiModel("sales/getSales", POST)
	SalesGetSalesDetail       = NewApiModel("sales/getSalesDetail", GET)
	SalesDeleteProduct        = NewApiModel("sales/deleteProduct", POST)
	SalesUpdateProduct        = NewApiModel("sales/updateProduct", POST)
	SalesAddProduct           = NewApiModel("sales/addProduct", POST)
	SalesUpload               = NewApiModel("sales/upload", POST)
	SalesGetProductGroupList  = NewApiModel("sales/getProductGroupList", GET)
	SalesGetProductList       = NewApiModel("sales/getProductList", POST)
	SalesAddSales             = NewApiModel("sales/addSales", POST)
	SalesGetProductsDetail    = NewApiModel("sales/getProductsDetail", POST)
)

// label
var (
	LabelAddLabelGroup        = NewApiModel("label/addLabelGroup", POST)
	LabelGetLabelInfo         = NewApiModel("LabelAddLabelGroup", POST)
	LabelAddLabel             = NewApiModel("label/addLabel", POST)
	LabelUpdate               = NewApiModel("LabelUpdate", POST)
	LabelDeleteCrmLabels      = NewApiModel("label/deleteCrmLabels", POST)
	LabelUpdateLabelGroupName = NewApiModel("label/updateLabelGroupName", POST)
)

// contact
var (
	ContactBookAdd    = NewApiModel("contactbook/add", POST)
	ContactBookDelete = NewApiModel("contactbook/delete", POST)
	ContactBookUpdate = NewApiModel("contactbook/update", POST)
	ContactBookList   = NewApiModel("contactbook/list", POST)
)

// im
var (
	ImMessageSend = NewApiModel("im/message/send", POST)
)

// apipush
var (
	ApiPushGetApiPush = NewApiModel("apipush/getApiPush", POST)
)

// clue
var (
	ClueImport = NewApiModel("clue/import", POST)
)
