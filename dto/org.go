package dto

type OrgStructInfoResponseDTO struct {
	CommonDTO
	Data *OrgStructInfoResponseData `json:"data"`
}

func NewOrgStructInfoResponseDTO() *OrgStructInfoResponseDTO {
	return &OrgStructInfoResponseDTO{}
}

type OrgStructInfoResponseData struct {
	Depts []*OrgStructInfoResponseDept  `json:"depts"`
	Users []*OrgStructInfoResponseUsers `json:"users"`
}

type OrgStructInfoResponseDept struct {
	DeptId       int64  `json:"deptId"`
	DeptName     string `json:"deptName"`
	ParentDeptId int64  `json:"parentDeptId"`
}

type OrgStructInfoResponseUsers struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	DeptId   int64  `json:"deptId"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
}

type OrgDeptCreateRequestDTO struct {
	DeptName     string `json:"deptName"`
	ParentDeptId int64  `json:"parentDeptId"`
	OptUserId    int64  `json:"optUserId"`
}

type OrgDeptCreateResponseDTO struct {
	CommonDTO
	Data *OrgDeptCreateResponseData `json:"data"`
}

type OrgDeptCreateResponseData struct {
	DeptId int64 `json:"deptId"`
}

type OrgDeptEditRequestDTO struct {
	DeptId       int64  `json:"deptId"`
	DeptName     string `json:"deptName"`
	ParentDeptId int64  `json:"parentDeptId"`
	OptUserId    int64  `json:"optUserId"`
}

type OrgDeptDeleteRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	DeptId    int64 `json:"deptId"`
}

type OrgDeptDeleteResponseDTO struct {
	CommonDTO
	Data *OrgDeptDeleteResponseData `json:"data"`
}

func NewOrgDeptDeleteResponseDTO() *OrgDeptDeleteResponseDTO {
	return &OrgDeptDeleteResponseDTO{}
}

type OrgDeptDeleteResponseData struct {
	OrgDeptDeleteRequestDTO
}

type OrgUserDeleteRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	UserId    int64 `json:"userId"`
}

type OrgUserDeleteResponseDTO struct {
	CommonDTO
	Data *OrgUserDeleteResponseData `json:"data"`
}

func NewOrgUserDeleteResponseDTO() *OrgUserDeleteResponseDTO {
	return &OrgUserDeleteResponseDTO{}
}

type OrgUserDeleteResponseData struct {
	OrgUserDeleteRequestDTO
}
