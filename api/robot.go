package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// RobotAddTask 新增任务
func (c *Config) RobotAddTask(reqData *dto.RobotAddTaskRequestDTO) (res *dto.RobotAddTaskResponseDTO, err error) {
	res = dto.NewRobotAddTaskDTO()
	err = c.RequestByAutoResolve(common.RobotAddTask, reqData, res)
	return
}

// RobotAddTaskRecord 新增机器人通话记录
func (c *Config) RobotAddTaskRecord(reqData *dto.RobotAddTaskRecordRequestDTO) (res *dto.RobotAddTaskRecordResponseDTO, err error) {
	res = dto.NewRobotAddTaskRecordDTO()
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
func (c *Config) AsyncHroniZationCreate(reqData *dto.AsyncHroniZationCreateRequestDTO) (res *dto.AsyncHroniZationCreateResponseDTO, err error) {
	res = dto.NewAsyncHroniZationCreateDTO()
	err = c.RequestByAutoResolve(common.AsyncHroniZationCreate, reqData, res)
	return
}
