package dto

type ClueImportRequestDTO struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Source     int64  `json:"source"`
	UserId     int64  `json:"userId"`
	Company    string `json:"company"`
	IndustryId int64  `json:"industryId"`
	Country    int64  `json:"country"`
	Province   int64  `json:"province"`
	City       int64  `json:"city"`
	Region     int64  `json:"region"`
	Qq         string `json:"qq"`
	Email      string `json:"email"`
	Wx         string `json:"wx"`
	Remark     string `json:"remark"`
}
