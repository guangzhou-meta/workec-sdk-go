package common

const (
	ECRequestContentType      = "Content-Type"
	ECRequestContentTypeValue = "application/json"
	ECRequestCorpId           = "X-Ec-Cid"
	ECRequestSign             = "X-Ec-Sign"
	ECRequestTimeStamp        = "X-Ec-TimeStamp"
)

const (
	ServerBaseUrlV1      = "https://open.workec.com/"
	ServerBaseUrlV2      = "https://open.workec.com/v2/"
	DefaultServerBaseUrl = ServerBaseUrlV2
)

type RoleChangeUserModifyType int

const (
	RoleChangeUserAdd    RoleChangeUserModifyType = 1
	RoleChangeUserDelete RoleChangeUserModifyType = 2
)
