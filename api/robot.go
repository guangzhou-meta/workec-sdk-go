package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// RobotAddTask 新增任务
func (c *Config) RobotAddTask(reqData *dto.RobotAddTaskRequestDTO) (res *dto.RobotAddTaskResponseDTO, err error) {
	res = dto.NewRobotAddTaskResponseDTO()
	err = c.RequestByAutoResolve(common.RobotAddTask, reqData, res)
	return
}

// RobotAddTaskRecord 新增机器人通话记录
func (c *Config) RobotAddTaskRecord(reqData *dto.RobotAddTaskRecordRequestDTO) (res *dto.RobotAddTaskRecordResponseDTO, err error) {
	res = dto.NewRobotAddTaskRecordResponseDTO()
	err = c.RequestByAutoResolve(common.RobotAddTaskRecord, reqData, res)
	return
}

// RobotUpDateTask 更新任务
func (c *Config) RobotUpDateTask(reqData *dto.RobotUpDateTaskRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.RobotUpDateTask, reqData, res)
	return
}

// AsyncHroniZationCreate 创建任务
func (c *Config) AsynchronizationCreate(reqData *dto.AsynchronizationCreateRequestDTO) (res *dto.AsynchronizationCreateResponseDTO, err error) {
	res = dto.NewAsynchronizationCreateResponseDTO()
	err = c.RequestByAutoResolve(common.AsynchronizationCreate, reqData, res)
	return
}

// AsynchroniZationQuery 查询任务
func (c *Config) AsynchronizationQuery(reqData *dto.AsynchronizationQueryRequestDTO) (res *dto.AsynchronizationQueryResponseDTO, err error) {
	res = dto.NewAsynchronizationQueryResponseDTO()
	err = c.RequestByAutoResolve(common.AsynchronizationQuery, reqData, res)
	return
}

// ConfigGetBookFieldMapping 获取企业联系人自定义字段
func (c *Config) ConfigGetBookFieldMapping() (res *dto.ConfigGetBookFieldMappingResponseDTO, err error) {
	res = dto.NewConfigGetBookFieldMappingResponseDTO()
	err = c.RequestByAutoResolve(common.ConfigGetBookFieldMapping, nil, res)
	return
}

// ConfigGetFieldMapping 获取自定义字段
func (c *Config) ConfigGetFieldMapping() (res *dto.ConfigGetFieldMappingResponseDTO, err error) {
	res = dto.NewConfigGetFieldMappingResponseDTO()
	err = c.RequestByAutoResolve(common.ConfigGetFieldMapping, nil, res)
	return
}

// ConfigGetPubicPond 获取业务组信息
func (c *Config) ConfigGetPubicPond() (res *dto.ConfigGetPubicPondResponseDTO, err error) {
	res = dto.NewConfigGetPubicPondResponseDTO()
	err = c.RequestByAutoResolve(common.ConfigGetPubicPond, nil, res)
	return
}

// ConfigGetStages 获取客户进展信息
func (c *Config) ConfigGetStages() (res *dto.ConfigGetStagesResponseDTO, err error) {
	res = dto.NewConfigGetStagesResponseDTO()
	err = c.RequestByAutoResolve(common.ConfigGetStages, nil, res)
	return
}
