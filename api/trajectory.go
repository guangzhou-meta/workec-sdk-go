package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// TrajectoryFindUserTrajectory 联系记录 - 导出跟进记录
func (c *Config) TrajectoryFindUserTrajectory(reqData *dto.TrajectoryFindUserTrajectoryRequestDTO) (res *dto.TrajectoryFindUserTrajectoryResponseDTO, err error) {
	res = dto.NewTrajectoryFindUserTrajectoryResponseDTO()
	err = c.RequestByAutoResolve(common.TrajectoryFindUserTrajectory, reqData, res)
	return
}

// TrajectorySaveUserTrajectory 联系记录 - 批量添加跟进记录
func (c *Config) TrajectorySaveUserTrajectory(reqData *dto.TrajectorySaveUserTrajectoryRequestDTO) (res *dto.TrajectorySaveUserTrajectoryResponseDTO, err error) {
	res = dto.NewTrajectorySaveUserTrajectoryResponseDTO()
	err = c.RequestByAutoResolve(common.TrajectorySaveUserTrajectory, reqData, res)
	return
}

// TrajectoryFindHistoryUserTrajectory 联系记录 - 导出历史跟进记录
func (c *Config) TrajectoryFindHistoryUserTrajectory(reqData *dto.TrajectoryFindHistoryUserTrajectoryRequestDTO) (res *dto.TrajectoryFindHistoryUserTrajectoryResponseDTO, err error) {
	res = dto.NewTrajectoryFindHistoryUserTrajectoryResponseDTO()
	err = c.RequestByAutoResolve(common.TrajectoryFindHistoryUserTrajectory, reqData, res)
	return
}

// RecordSendSmsHistory 联系记录 - 导出历史短信记录
func (c *Config) RecordSendSmsHistory(reqData *dto.RecordSendSmsHistoryRequestDTO) (res *dto.RecordSendSmsHistoryResponseDTO, err error) {
	res = dto.NewRecordSendSmsHistoryResponseDTO()
	err = c.RequestByAutoResolve(common.RecordSendSmsHistory, reqData, res)
	return
}

// RecordTelRecordHistoryQuery 导出电话历史记录
func (c *Config) RecordTelRecordHistoryQuery(reqData *dto.RecordTelRecordHistoryQueryRequestDTO) (res *dto.RecordTelRecordHistoryQueryResponseDTO, err error) {
	res = dto.NewRecordTelRecordHistoryQueryResponseDTO()
	err = c.RequestByAutoResolve(common.RecordTelRecordHistoryQuery, reqData, res)
	return
}

// RecordAddTelRecord 添加电话记录
func (c *Config) RecordAddTelRecord(reqData *dto.RecordAddTelRecordRequestDTO) (res *dto.RecordAddTelRecordResponseDTO, err error) {
	res = dto.NewRecordAddTelRecordResponseDTO()
	err = c.RequestByAutoResolve(common.RecordAddTelRecord, reqData, res)
	return
}
