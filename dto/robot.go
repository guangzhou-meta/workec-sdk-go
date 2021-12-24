package dto

type RobotAddTaskRequestDTO struct {
	Title      string `json:"title"`
	Type       int64  `json:"type"`
	UserId     int64  `json:"userId"`
	Time       string `json:"time"`
	FinishTime string `json:"finishTime"`
	Finish     string `json:"finish"`
	Total      string `json:"total"`
	Craft      string `json:"craft"`
	RobotId    int64  `json:"robotId"`
}

type RobotAddTaskResponseDTO struct {
	CommonDTO
	Data map[string]int `json:"data"`
}

func NewRobotAddTaskDTO() *RobotAddTaskResponseDTO {
	return &RobotAddTaskResponseDTO{}
}

type RobotAddTaskRecordRequestDTO struct {
	TaskId int64                               `json:"taskId"`
	Data   []*RobotAddTaskRecordRequestDTOData `json:"data"`
}
type RobotAddTaskRecordRequestDTOData struct {
	Mobile   string `json:"mobile"`
	Phone    string `json:"phone"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	CallTime string `json:"callTime"`
	TalkTime int    `json:"talkTime"`
	Record   string `json:"record"`
}

type RobotAddTaskRecordResponseDTO struct {
	CommonDTO
	Data map[string]string
}

func NewRobotAddTaskRecordDTO() *RobotAddTaskRecordResponseDTO {
	return &RobotAddTaskRecordResponseDTO{}
}

type RobotTimeRangeBaseVO struct {
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
}

type RobotUpDateTaskRequestDTO struct {
	TaskId     int64  `json:"taskId"`
	Title      string `json:"title"`
	FinishTime string `json:"finishTime"`
	Craft      string `json:"craft"`
	Finish     string `json:"finish"`
}

type AsyncHroniZationCreateRequestDTO struct {
	TaskName     string                `json:"taskName"`
	Type         int                   `json:"type"`
	Step         int                   `json:"step"`
	LabelIds     string                `json:"labelIds"`
	FollowUserId int64                 `json:"followUserId"`
	PublicPondId int64                 `json:"publicPondId"`
	ModifyTime   *RobotTimeRangeBaseVO `json:"modifyTime"`
	ContactTime  *RobotTimeRangeBaseVO `json:"contactTime"`
	CreateTime   *RobotTimeRangeBaseVO `json:"createTime"`
}
