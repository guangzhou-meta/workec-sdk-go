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

func NewRobotAddTaskResponseDTO() *RobotAddTaskResponseDTO {
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

func NewRobotAddTaskRecordResponseDTO() *RobotAddTaskRecordResponseDTO {
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

type AsynchronizationCreateRequestDTO struct {
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

type AsynchronizationCreateResponseDTO struct {
	CommonDTO
	Data AsynchronizationCreateResponseData `json:"data"`
}

type AsynchronizationCreateResponseData struct {
	TaskId int64 `json:"taskId"`
}

func NewAsynchronizationCreateResponseDTO() *AsynchronizationCreateResponseDTO {
	return &AsynchronizationCreateResponseDTO{}
}

type AsynchronizationQueryRequestDTO struct {
	TaskId     int `json:"taskId"`
	TaskStatus int `json:"taskStatus"`
}

type AsynchronizationQueryResponseDTO struct {
	CommonDTO
	Data *AsynchroniZationQueryResponseDTOData `json:"data"`
}

type AsynchroniZationQueryResponseDTOData struct {
	TaskId      int64  `json:"taskId"`
	TaskName    int64  `json:"taskName"`
	StackStatus int64  `json:"stackStatus"`
	File        string `json:"file"`
}

func NewAsynchronizationQueryResponseDTO() *AsynchronizationQueryResponseDTO {
	return &AsynchronizationQueryResponseDTO{}
}

type ConfigGetBookFieldMappingResponseDTO struct {
	CommonDTO
	Data *ConfigGetBookFieldMappingResponseDTOData `json:"data"`
}

type ConfigGetBookFieldMappingResponseDTOData struct {
	Id     int        `json:"id"`
	Text   string     `json:"text"`
	Type   int        `json:"type"`
	IsMust int        `json:"isMust"`
	Regex  string     `json:"regex"`
	Level  int        `json:"level"`
	Params []struct{} `json:"params"`
}

func NewConfigGetBookFieldMappingResponseDTO() *ConfigGetBookFieldMappingResponseDTO {
	return &ConfigGetBookFieldMappingResponseDTO{}
}

type ConfigGetFieldMappingResponseDTO struct {
	CommonDTO
	Data *ConfigGetFieldMappingResponseDTOData `json:"data"`
}

type ConfigGetFieldMappingResponseDTOData struct {
	FieldGroupId   int                                               `json:"fieldGroupId"`
	FieldGroupName string                                            `json:"fieldGroupName"`
	FieldGroupSort int                                               `json:"fieldGroupSort"`
	FieldId        int                                               `json:"fieldId"`
	FieldName      string                                            `json:"fieldName"`
	Category       int                                               `json:"category"`
	FieldParam     []*ConfigGetFieldMappingResponseDTODataFieldParam `json:"fieldParam"`
}
type ConfigGetFieldMappingResponseDTODataFieldParam struct {
	ParamId   int    `json:"paramId"`
	ParamName string `json:"paramName"`
	ParamSort int    `json:"paramSort"`
}

func NewConfigGetFieldMappingResponseDTO() *ConfigGetFieldMappingResponseDTO {
	return &ConfigGetFieldMappingResponseDTO{}
}

type ConfigGetPubicPondResponseDTO struct {
	CommonDTO
	Data *ConfigGetPubicPondResponseDTOData `json:"data"`
}
type ConfigGetPubicPondResponseDTOData struct {
	PublicPondId   int64  `json:"publicPondId"`
	PublicPondName string `json:"publicPondName"`
}

func NewConfigGetPubicPondResponseDTO() *ConfigGetPubicPondResponseDTO {
	return &ConfigGetPubicPondResponseDTO{}
}

type ConfigGetStagesResponseDTO struct {
	CommonDTO
	Data *ConfigGetStagesResponseDTOData `json:"data"`
}

type ConfigGetStagesResponseDTOData struct {
	Name   string `json:"name"`
	Stage  int    `json:"stage"`
	Status int    `json:"status"`
}

func NewConfigGetStagesResponseDTO() *ConfigGetStagesResponseDTO {
	return &ConfigGetStagesResponseDTO{}
}
