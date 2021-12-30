package dto

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
)

type LabelAddLabelGroupRequestDTO struct {
	Name   string `json:"name"`
	Type   int    `json:"type"`
	Color  string `json:"color"`
	UserId int64  `json:"userId"`
}

type LabelGetLabelInfoRequestDTO struct {
	GroupValue string `json:"groupValue"`
}

type LabelGetLabelInfoResponseDTO struct {
	CommonDTO
	Data *LabelGetLabelInfoResponseDTOData `json:"data"`
}

type LabelGetLabelInfoResponseDTOData struct {
	LabelGroupId   int                                          `json:"labelGroupId"`
	LabelGroupName string                                       `json:"labelGroupName"`
	LabelGroupSort int                                          `json:"labelGroupSort"`
	LabelGroupType int                                          `json:"labelGroupType"`
	LabelList      []*LabelGetLabelInfoResponseDTODataLabelList `json:"labelList"`
}

type LabelGetLabelInfoResponseDTODataLabelList struct {
	LabelId   int64  `json:"labelId"`
	LabelName string `json:"labelName"`
	LabelSort int    `json:"labelSort"`
}

func NewLabelGetLabelInfoResponseDTO() *LabelGetLabelInfoResponseDTO {
	return &LabelGetLabelInfoResponseDTO{}
}

type LabelAddLabelRequestDTO struct {
	Name       string `json:"name"`
	GroupValue string `json:"groupValue"`
	UserId     int64  `json:"userId"`
}

type LabelUpdateRequestDTO struct {
	OptUserId int64  `json:"optUserId"`
	CrmId     string `json:"crmId"`
	Lables    string `json:"lables"`
	Type      int    `json:"type"`
}

type LabelUpdateResponseDTO struct {
	CommonDTO
	Data bool `json:"data"`
}

func NewLabelUpdateResponseDTO() *LabelUpdateResponseDTO {
	return &LabelUpdateResponseDTO{}
}

type LabelDeleteCrmLabelsRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	CrmIds    []int `json:"crmIds"`
	LabelIds  []int `json:"labelIds"`
}

type LabelDeleteCrmLabelsResponseDTO struct {
	CommonDTO
	Data *LabelDeleteCrmLabelsResponseDTOData `json:"data"`
}

type LabelDeleteCrmLabelsResponseDTOData struct {
	SuccessCrmIds []int64 `json:"successCrmIds"`
	FailureCrmIds []int64 `json:"failureCrmIds"`
}

func NewLabelDeleteCrmLabelsResponseDTO() *LabelDeleteCrmLabelsResponseDTO {
	return &LabelDeleteCrmLabelsResponseDTO{}
}

type LabelUpdateLabelGroupNameRequestDTO struct {
	OptUserId    int64 `json:"optUserId"`
	LabelGroupId int64 `json:"labelGroupId"`
	Name         int64 `json:"name"`
}

type ContactBookAddRequestDTO struct {
	OptUserId        int64                    `json:"optUserId"`
	CrmId            int64                    `json:"crmId"`
	CrmContactBookVO *ContactCrmContactBookVO `json:"crmContactBookVO"`
}

type ContactCrmContactBookVO struct {
	Name           string                `json:"name"`
	CallName       string                `json:"callName"`
	Title          string                `json:"title"`
	Mobile         string                `json:"mobile"`
	MobileCode     string                `json:"mobileCode"`
	Phone          string                `json:"phone"`
	PhoneCode      string                `json:"phoneCode"`
	PhoneExtension string                `json:"phoneExtension"`
	Qq             string                `json:"qq"`
	Email          string                `json:"email"`
	Wechat         string                `json:"wechat"`
	Birthday       string                `json:"birthday"`
	LunarBirthday  string                `json:"lunarBirthday"`
	LunarLeap      *common.ContactLunar  `json:"lunarLeap"`
	Memo           string                `json:"memo"`
	Gender         *common.ContactGender `json:"gender"`
	Fields         map[string]string     `json:"fields"`
}

type ContactBookAddResponseDTO struct {
	CommonDTO
	Data int64 `json:"data"`
}

func NewContactBookAddResponseDTO() *ContactBookAddResponseDTO {
	return &ContactBookAddResponseDTO{}
}

type ContactBookDeleteRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	CrmId     int64 `json:"crmId"`
	Id        int64 `json:"id"`
}

type ContactBookUpdateRequestDTO struct {
	OptUserId        int64                    `json:"optUserId"`
	CrmId            int64                    `json:"crmId"`
	Id               int64                    `json:"id"`
	CrmContactBookVO *ContactCrmContactBookVO `json:"crmContactBookVO"`
}

type ContactBookListRequestDTO struct {
	OptUserId int64 `json:"optUserId"`
	CrmId     int64 `json:"crmId"`
}

type ContactBookListResponseDTO struct {
	CommonDTO
	Data *ContactCrmContactBookVO `json:"data"`
}

func NewContactBookListResponseDTO() *ContactBookListResponseDTO {
	return &ContactBookListResponseDTO{}
}

type ImMessageSendRequestDTO struct {
	ChannelId        int64                                   `json:"channelId"`
	ReceiverType     int64                                   `json:"receiverType"`
	Receiver         []int64                                 `json:"receiver"`
	BodyTitle        string                                  `json:"bodyTitle"`
	BodyContent      string                                  `json:"bodyContent"`
	Images           []string                                `json:"images"`
	OriginLink       string                                  `json:"originLink"`
	OriginMobileLink string                                  `json:"originMobileLink"`
	OriginLinkType   string                                  `json:"originLinkType"`
	Style            *ImMessageSendRequestDTOCompositeObject `json:"style"`
	Click            *ImMessageSendRequestDTOCompositeObject `json:"click"`
	File             *ImMessageSendRequestDTOFile            `json:"file"`
	AddMsg           *ImMessageSendRequestDTOAddMsg          `json:"addMsg"`
}

type ImMessageSendRequestDTOCompositeObject struct {
	HideSubmit  bool   `json:"hideSubmit"`
	HideCancel  bool   `json:"hideCancel"`
	Submit      string `json:"submit"`
	SubmitAfter string `json:"submitAfter"`
	SubmitUrl   string `json:"submitUrl"`
	Cancel      string `json:"cancel"`
	CancelAfter string `json:"cancelAfter"`
	CancelUrl   string `json:"cancelUrl"`
}

type ImMessageSendRequestDTOAddMsg struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ImMessageSendRequestDTOFile struct {
	Icon     string   `json:"icon"`
	Filename string   `json:"filename"`
	Target   []string `json:"target"`
}
