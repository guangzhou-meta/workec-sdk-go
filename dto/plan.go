package dto

type PlanGetPlanTemplateRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
}

type PlanGetPlanTemplateResponseDTO struct {
	CommonDTO
	Data []*PlanGetPlanTemplateResponseDTOData `json:"data"`
}

type PlanGetPlanTemplateResponseDTOData struct {
	Modelld     int    `json:"modelld"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	GroupId     int64  `json:"groupId"`
	CreateTime  string `json:"createTime"`
	PlanType    int    `json:"planType"`
	DateTYpe    int    `json:"dateTYpe"`
	DateValue   string `json:"dateValue"`
	Call        int    `json:"call"`
	Sigin       int    `json:"sigin"`
	LetterPaper int    `json:"letterPaper"`
	SubTitle    string `json:"subTitle"`
	SubType     string `json:"subType"`
}

func NewPlanGetPlanTemplateResponseDTO() *PlanGetPlanTemplateResponseDTO {
	return &PlanGetPlanTemplateResponseDTO{}
}
