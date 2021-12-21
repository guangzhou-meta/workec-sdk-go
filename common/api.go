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

	OrgUserDelete = NewApiModel("org/user/delete", POST)
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
