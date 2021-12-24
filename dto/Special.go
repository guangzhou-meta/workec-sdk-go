package dto

type SpecialSelectRequestDTO struct {
	ServiceId string `json:"serviceId"`
	Params    int    `json:"params"`
}

type SpecialSelectResponseDTO struct {
	Id            int64  `json:"id"`
	ConfigId      int    `json:"configId"`
	CorpId        int64  `json:"corpId"`
	UserId        int64  `json:"userId"`
	DepartmentId  int    `json:"departmentId"`
	UserName      string `json:"userName"`
	Remark        string `json:"remark"`
	Description   string `json:"description"`
	Memo          string `json:"memo"`
	Condition     int    `json:"condition"`
	SubCondition  string `json:"subCondition"`
	Type          int    `json:"type"`
	IsRead        int    `json:"isRead"`
	DeadlineTime  string `json:"deadlineTime"`
	FinishTime    int    `json:"finishTime"`
	FinishNumber  int    `json:"finishNumber"`
	OverdueNumber int    `json:"overdueNumber"`
	Total         int    `json:"total"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

func NewSpecialSelectResponseDTO() *SpecialSelectResponseDTO {
	return &SpecialSelectResponseDTO{}
}
