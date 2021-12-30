package dto

type StepUpdateRequestDTO struct {
	OptUserId int64  `json:"optUserId"`
	CrmIds    string `json:"crmIds"`
	Step      int    `json:"step"`
}

type StepUpdateResponseDTO struct {
	CommonDTO
	Data bool `json:"data"`
}

func NewStepUpdateResponseDTO() *StepUpdateResponseDTO {
	return &StepUpdateResponseDTO{}
}
