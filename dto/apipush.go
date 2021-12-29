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
	FId         int64  `json:"f_id"`
	FCostTime   int64  `json:"f_cost_time"`
	FCreateTime string `json:"f_create_time"`
	FCorpId     int64  `json:"f_corp_id"`
	FPushParas  string `json:"f_push_paras"`
	FPushUrl    string `json:"f_push_url"`
	FReturnCode int64  `json:"f_return_code"`
}

func NewApiPushGetApiPushRequestDTO() *ApiPushGetApiPushResponseDTO {
	return &ApiPushGetApiPushResponseDTO{}
}
