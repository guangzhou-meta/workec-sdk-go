package dto

type CommonDTO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCommonDTO() *CommonDTO {
	return &CommonDTO{}
}
