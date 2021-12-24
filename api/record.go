package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// RecordCall 电话外呼
func (c *Config) RecordCall(reqData *dto.RecordCallRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RecordCall, reqData, res)
	return
}

// RecordGetFreeStatusUid 电话空闲用户
func (c *Config) RecordGetFreeStatusUid(reqData *dto.RecordGetFreeStatusUidRequestDTO) (res *dto.RecordGetFreeStatusUidResponseDTO, err error) {
	res = dto.NewRecordGetFreeStatusUidResponseDTO()
	err = c.RequestByAutoResolve(common.RecordGetFreeStatusUid, reqData, res)
	return
}

// RecordSmsRecord 短信记录
func (c *Config) RecordSmsRecord(reqData *dto.RecordSmsRecordRequestDTO) (res *dto.RecordSmsRecordResponseDTO, err error) {
	res = dto.NewRecordSmsRecordResponseDTO()
	err = c.RequestByAutoResolve(common.RecordSmsRecord, reqData, res)
	return
}

// RecordTelRecord 电话记录
func (c *Config) RecordTelRecord(reqData *dto.RecordTelRecordRequestDTO) (res *dto.RecordTelRecordResponseDTO, err error) {
	res = dto.NewRecordTelRecordResponseDTO()
	err = c.RequestByAutoResolve(common.RecordTelRecord, reqData, res)
	return
}
