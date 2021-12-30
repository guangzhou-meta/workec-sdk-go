package dto

type RecordCallRequestDTO struct {
	Userid         string `json:"userid"`
	Phone          string `json:"phone"`
	Name           string `json:"name"`
	Memo           string `json:"memo"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Title          string `json:"title"`
	CompanyAddress string `json:"companyAddress"`
	Company        string `json:"company"`
	CompanyUrl     string `json:"companyUrl"`
	ClientType     int    `json:"clientType"`
}

type RecordCallResponseDTO struct {
	CommonDTO
}

type RecordGetFreeStatusUidRequestDTO struct {
	UserState int `json:"userState1"`
}

type RecordGetFreeStatusUidResponseDTO struct {
	CommonDTO
	Data []int64 `json:"data"`
}

func NewRecordGetFreeStatusUidResponseDTO() *RecordGetFreeStatusUidResponseDTO {
	return &RecordGetFreeStatusUidResponseDTO{}
}

type RecordSmsRecordRequestDTO struct {
	Userid       int64                          `json:"userid"`
	CrmIds       string                         `json:"crmIds"`
	PhoneNumbers string                         `json:"phoneNumbers"`
	Date         *RecordSmsRecordRequestDTODate `json:"date"`
	LastId       int64                          `json:"lastId"`
	PageSize     int                            `json:"pageSize"`
}

type RecordSmsRecordRequestDTODate struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type RecordSmsRecordResponseDTO struct {
	CommonDTO
	Data *RecordSmsRecordResponseDTOData
}

type RecordSmsRecordResponseDTOData struct {
	SmsRecordList []*RecordSmsRecordResponseDTODataSmsRecordList
	NextPageDTO   *RecordSmsRecordResponseDTODataNextPageDTO
}
type RecordSmsRecordResponseDTODataSmsRecordList struct {
	SendTime    string `json:"sendTime"`
	PhoneNumber string `json:"phoneNumber"`
	CrmId       int64  `json:"crmId"`
	Userid      int64  `json:"userid"`
	Content     string `json:"content"`
}

type RecordSmsRecordResponseDTODataNextPageDTO struct {
	PageSize    int    `json:"pageSize"`
	HasNextPage int    `json:"hasNextPage"`
	NextLastId  string `json:"nextLastId"`
}

func NewRecordSmsRecordResponseDTO() *RecordSmsRecordResponseDTO {
	return &RecordSmsRecordResponseDTO{}
}

type RecordTelRecordRequestDTO struct {
	UserId       int64  `json:"userId"`
	CrmIds       string `json:"crmIds"`
	PhoneNumbers string `json:"phoneNumbers"`
	Date         *RecordTelRecordRequestDTOData
	LastId       int64 `json:"lastId"`
	PageSize     int   `json:"pageSize"`
}

type RecordTelRecordRequestDTOData struct {
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
}

type RecordTelRecordResponseDTO struct {
	CommonDTO
	Data *RecordTelRecordResponseDTOData `json:"data"`
}

type RecordTelRecordResponseDTOData struct {
	TelRecordList *RecordTelRecordResponseDTODataTelRecordList
	NextPageDTO   *RecordTelRecordResponseDTODataNextPageDTO
}

type RecordTelRecordResponseDTODataTelRecordList struct {
	StartTime   string `json:"startTime"`
	PhoneNumber string `json:"phoneNumber"`
	CrmId       int64  `json:"crmId"`
	CallTime    int64  `json:"callTime"`
	UserId      int64  `json:"userId"`
	Path        string `json:"path"`
	CallType    int    `json:"callType"`
	InOutType   int    `json:"inOutType"`
	Memo        string `json:"memo"`
	WasteId     int64  `json:"wasteId"`
}

type RecordTelRecordResponseDTODataNextPageDTO struct {
	HasNextPage  int    `json:"hasNextPage"`
	NextLastId   int64  `json:"nextLastId"`
	NextLastTime string `json:"nextLastTime"`
	PageSize     int    `json:"pageSize"`
}

func NewRecordTelRecordResponseDTO() *RecordTelRecordResponseDTO {
	return &RecordTelRecordResponseDTO{}
}
