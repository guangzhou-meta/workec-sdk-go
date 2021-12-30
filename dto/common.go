package dto

type CommonDTO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCommonDTO() *CommonDTO {
	return &CommonDTO{}
}

type TrajectoryDTO struct {
	Data *TrajectoryDTOData `json:"data"`
}

type TrajectoryDTOData struct {
	PageSize  int                        `json:"pageSize"`
	PageNo    int                        `json:"pageNo"`
	Total     int                        `json:"total"`
	ManPageNo int                        `json:"manPageNo"`
	StartRow  int                        `json:"startRow"`
	Result    []*TrajectoryDTODataResult `json:"result"`
}

type TrajectoryDTODataResult struct {
	ContactTime     string `json:"contactTime"`
	CrmId           int64  `json:"crmId"`
	Customer        string `json:"customer"`
	CustomerCompany string `json:"customerCompany"`
	UserId          int64  `json:"userId"`
	Content         string `json:"content"`
	Md5             string `json:"md5"`
}
