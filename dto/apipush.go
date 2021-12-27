package dto

type ApiPushGetApiPushRequestDTO struct {
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	PageSize  int    `json:"pageSize"`
	OnlyError bool   `json:"onlyError"`
}

type ApiPushGetApiPushResponseDTO struct {
	CommonDTO
	Data []*ApiPushGetApiPushResponseDTOData `json:"data"`
}
type ApiPushGetApiPushResponseDTOData struct {
	F_Id          int64  `json:"f_id"`
	F_Cost_Time   int64  `json:"f_cost_time"`
	F_Create_Time string `json:"f_create_time"`
	F_Corp_Id     int64  `json:"f_corp_id"`
	F_Push_Paras  string `json:"f_push_paras"`
	F_Push_Url    string `json:"f_push_url"`
	F_Return_Code int64  `json:"f_return_code"`
}

func NewApiPushGetApiPushRequestDTO() *ApiPushGetApiPushResponseDTO {
	return &ApiPushGetApiPushResponseDTO{}
}
