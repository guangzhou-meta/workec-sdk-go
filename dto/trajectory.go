package dto

import "github.com/guangzhou-meta/workec-sdk-go/common"

type TrajectoryFindUserTrajectoryRequestDTO struct {
	DeptIds   string `json:"deptIds"`
	UserIds   string `json:"userIds"`
	CrmIds    string `json:"crmIds"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	PageNo    int    `json:"pageNo"`
}

type TrajectoryFindUserTrajectoryResponseDTO struct {
	CommonDTO
	TrajectoryDTO
}

func NewTrajectoryFindUserTrajectoryResponseDTO() *TrajectoryFindUserTrajectoryResponseDTO {
	return &TrajectoryFindUserTrajectoryResponseDTO{}
}

type TrajectorySaveUserTrajectoryRequestDTO struct {
	List []*TrajectorySaveUserTrajectoryRequestDTOLIst `json:"list"`
}
type TrajectorySaveUserTrajectoryRequestDTOLIst struct {
	Id          string `json:"id"`
	CrmId       int64  `json:"crmId"`
	ContactTime string `json:"contactTime"`
	UserId      string `json:"userId"`
	Content     string `json:"content"`
}

type TrajectorySaveUserTrajectoryResponseDTO struct {
	CommonDTO
	TrajectoryDTO
}

func NewTrajectorySaveUserTrajectoryResponseDTO() *TrajectorySaveUserTrajectoryResponseDTO {
	return &TrajectorySaveUserTrajectoryResponseDTO{}
}

type TrajectoryFindHistoryUserTrajectoryRequestDTO struct {
	DeptIds         string `json:"deptIds"`
	UserId          string `json:"userId"`
	CrmId           string `json:"crmId"`
	Year            int    `json:"year"`
	Month           int    `json:"month"`
	StartDayOfMonth int    `json:"startDayOfMonth"`
	EndDayOfMonth   int    `json:"endDayOfMonth"`
	PageNo          int    `json:"pageNo"`
}
type TrajectoryFindHistoryUserTrajectoryResponseDTO struct {
	CommonDTO
	TrajectoryDTO
}

func NewTrajectoryFindHistoryUserTrajectoryResponseDTO() *TrajectoryFindHistoryUserTrajectoryResponseDTO {
	return &TrajectoryFindHistoryUserTrajectoryResponseDTO{}
}

type RecordSendSmsHistoryRequestDTO struct {
	DeptIds         string `json:"deptIds"`
	UserIds         string `json:"userIds"`
	CrmIds          string `json:"crmIds"`
	Year            int    `json:"year"`
	Month           int    `json:"month"`
	StartDayOfMonth int    `json:"startDayOfMonth"`
	EndDayOfMonth   int    `json:"endDayOfMonth"`
	PageNo          int    `json:"pageNo"`
}

type RecordSendSmsHistoryResponseDTO struct {
	CommonDTO
	Data *RecordSendSmsHistoryResponseDTOData
}

type RecordSendSmsHistoryResponseDTOData struct {
	PageSize  int                                        `json:"pageSize"`
	PageNo    int                                        `json:"pageNo"`
	Total     int                                        `json:"total"`
	ManPageNo int                                        `json:"manPageNo"`
	StartRow  int                                        `json:"startRow"`
	Result    *RecordSendSmsHistoryResponseDTODataResult `json:"result"`
}

type RecordSendSmsHistoryResponseDTODataResult struct {
	ContactTime     string `json:"contactTime"`
	CrmId           int64  `json:"crmId"`
	Customer        string `json:"customer"`
	CustomerCompany string `json:"customerCompany"`
	UserId          int64  `json:"userId"`
	Content         string `json:"content"`
	Md5             string `json:"md5"`
}

func NewRecordSendSmsHistoryResponseDTO() *RecordSendSmsHistoryResponseDTO {
	return &RecordSendSmsHistoryResponseDTO{}
}

type RecordTelRecordHistoryQueryRequestDTO struct {
	UserIds string                                     `json:"userIds"`
	CrmIds  string                                     `json:"crmIds"`
	PhoneNo string                                     `json:"phoneNo"`
	Date    *RecordTelRecordHistoryQueryRequestDTODate `json:"date"`
	Page    *RecordTelRecordHistoryQueryRequestDTOPage `json:"page"`
}

type RecordTelRecordHistoryQueryRequestDTODate struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
type RecordTelRecordHistoryQueryRequestDTOPage struct {
	PageSize         int    `json:"pageSize"`
	PreLastWasteId   int64  `json:"preLastWasteId"`
	PreLastStartTime string `json:"preLastStartTime"`
}

type RecordTelRecordHistoryQueryResponseDTO struct {
	StartTime       string `json:"starttime"`
	CallToNo        string `json:"calltono"`
	CrmId           int64  `json:"crmId"`
	CallTime        string `json:"calltime"`
	Type            int    `json:"type"`
	UserId          string `json:"userId"`
	Path            string `json:"path"`
	CustomerName    string `json:"customerName"`
	CustomerCompany string `json:"customerCompany"`
	WasteId         int64  `json:"wasteId"`
	Memo            string `json:"memo"`
}

func NewRecordTelRecordHistoryQueryResponseDTO() *RecordTelRecordHistoryQueryResponseDTO {
	return &RecordTelRecordHistoryQueryResponseDTO{}
}

type RecordAddTelRecordRequestDTO struct {
	List *[]RecordAddTelRecordRequestDTOList `json:"list"`
}

type RecordAddTelRecordRequestDTOList struct {
	FUserId    string                      `json:"f_user_id"`
	FCrmId     string                      `json:"f_crm_id"`
	FStartTime string                      `json:"f_starttime"`
	FEndTime   string                      `json:"f_endtime"`
	FInOutType *common.TrajectoryinOutType `json:"f_in_out_type"`
	FCallTime  int                         `json:"f_calltime"`
	FCallNo    string                      `json:"f_callno"`
	FCallToNo  string                      `json:"f_calltono"`
	FPath      string                      `json:"f_path"`
}

type RecordAddTelRecordResponseDTO struct {
	CommonDTO
	Data *RecordAddTelRecordResponseDTOCdata `json:"data"`
}

type RecordAddTelRecordResponseDTOCdata struct {
	Index string `json:"index"`
	Mes   string `json:"mes"`
}

func NewRecordAddTelRecordResponseDTO() *RecordAddTelRecordResponseDTO {
	return &RecordAddTelRecordResponseDTO{}
}
