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
	CustomerAddCustomer         = NewApiModel("customer/addCustomer", POST)
	CustomerUpdateCustomer      = NewApiModel("customer/updateCustomer", POST)
	CustomerBatchUpdateCustomer = NewApiModel("customer/batchUpdateCustomer", POST)
	CustomerCombine             = NewApiModel("customer/combine", POST)
	CustomerFileUpload          = NewApiModel("customer/file/upload", POST)
	CustomerShare               = NewApiModel("customer/share", POST)
	CustomerChangeAbandon       = NewApiModel("customer/change/abandon", POST)
	CustomerChangeUser          = NewApiModel("customer/change/user", POST)
	CustomerFace                = NewApiModel("customer/face", POST)
	CustomerQueryLabel          = NewApiModel("customer/queryLabel", POST)
	CustomerQueryList           = NewApiModel("customer/queryList", POST)
	CustomerFileList            = NewApiModel("customer/file/list", POST)
	CustomerFolderList          = NewApiModel("customer/folder/list", POST)
	CustomerQueryExist          = NewApiModel("customer/queryExist", POST)
	CustomerQuery               = NewApiModel("customer/query", POST)
	CustomerPreciseQuery        = NewApiModel("customer/preciseQuery", POST)
	CustomerQueryCustomer       = NewApiModel("customer/queryCustomer", POST)
	CustomerGetCustomerGroup    = NewApiModel("customer/getCustomerGroup", POST)
	CustomerDelcrms             = NewApiModel("customer/delcrms", POST)
	CustomerGetTrajectory       = NewApiModel("customer/getTrajectory", POST)
	CustomerGetCrmVisitDetails  = NewApiModel("customer/getCrmVisitDetails", POST)
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
)
