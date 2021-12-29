package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// StatisticsDigitalMapPhone 电话统计-电话-数字图接口
func (c *Config) StatisticsDigitalMapPhone(reqData *dto.StatisticsDigitalMapPhoneRequestDTO) (res *dto.StatisticsDigitalMapPhoneResponseDTO, err error) {
	res = dto.NewStatisticsDigitalMapPhoneResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsDigitalMapPhone, reqData, res)
	return
}

// StatisticsLineGraphPhone 电话统计-电话-折线图接口
func (c *Config) StatisticsLineGraphPhone(reqData *dto.StatisticsLineGraphPhoneRequestDTO) (res *dto.StatisticsLineGraphPhoneResponseDTO, err error) {
	res = dto.NewStatisticsLineGraphPhoneResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsLineGraphPhone, reqData, res)
	return
}

// StatisticsDigitalMapWorkeffic 工作效率-数字图接口
func (c *Config) StatisticsDigitalMapWorkeffic(reqData *dto.StatisticsDigitalMapWorkefficRequestDTO) (res *dto.StatisticsDigitalMapWorkefficResponseDTO, err error) {
	res = dto.NewStatisticsDigitalMapWorkefficResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsDigitalMapWorkeffic, reqData, res)
	return
}

// StatisticsHistogramWorkeffic 工作效率-柱状图接口
func (c *Config) StatisticsHistogramWorkeffic(reqData *dto.StatisticsHistogramWorkefficRequestDTO) (res *dto.StatisticsHistogramWorkefficResponseDTO, err error) {
	res = dto.NewStatisticsHistogramWorkefficResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsHistogramWorkeffic, reqData, res)
	return
}

// StatisticsDigitalMapTag 标签-数字图接口
func (c *Config) StatisticsDigitalMapTag(reqData *dto.StatisticsDigitalMapTagRequestDTO) (res *dto.StatisticsDigitalMapTagResponseDTO, err error) {
	res = dto.NewStatisticsDigitalMapTagResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsDigitalMapTag, reqData, res)
	return
}

// StatisticsHistogramTag 标签-柱状图接口
func (c *Config) StatisticsHistogramTag(reqData *dto.StatisticsHistogramTagRequestDTO) (res *dto.StatisticsHistogramTagResponseDTO, err error) {
	res = dto.NewStatisticsHistogramTagResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsHistogramTag, reqData, res)
	return
}

// StatisticsDigitalMapCrmQuantity 客户数量-数字图接口
func (c *Config) StatisticsDigitalMapCrmQuantity(reqData *dto.StatisticsDigitalMapCrmQuantityRequestDTO) (res *dto.StatisticsDigitalMapCrmQuantityResponseDTO, err error) {
	res = dto.NewStatisticsDigitalMapCrmQuantityResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsDigitalMapCrmQuantity, reqData, res)
	return
}

// StatisticsCrmStatsWueryStepCountByChannel 按照渠道统计每个阶段的客户数
func (c *Config) StatisticsCrmStatsWueryStepCountByChannel(reqData *dto.StatisticsCrmStatsWueryStepCountByChannelRequestDTO) (res *dto.StatisticsCrmStatsWueryStepCountByChannelResponseDTO, err error) {
	res = dto.NewStatisticsCrmStatsWueryStepCountByChannelResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsCrmStatsWueryStepCountByChannel, reqData, res)
	return
}

// StatisticsCrmStatsGetTopStepCountByChannel TOP N客户渠道统计
func (c *Config) StatisticsCrmStatsGetTopStepCountByChannel(reqData *dto.StatisticsCrmStatsGetTopStepCountByChannelRequestDTO) (res *dto.StatisticsCrmStatsGetTopStepCountByChannelResponseDTO, err error) {
	res = dto.NewStatisticsCrmStatsGetTopStepCountByChannelResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsCrmStatsGetTopStepCountByChannel, reqData, res)
	return
}

// StatisticsCrmStatsGetStepCount 阶段客户数统计
func (c *Config) StatisticsCrmStatsGetStepCount(reqData *dto.StatisticsCrmStatsGetStepCountRequestDTO) (res *dto.StatisticsCrmStatsGetStepCountResponseDTO, err error) {
	res = dto.NewStatisticsCrmStatsGetStepCountResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsCrmStatsGetStepCount, reqData, res)
	return
}

// StatisticsCrmStatsGroupbyUserIds 客户数量-统计客户数量
func (c *Config) StatisticsCrmStatsGroupbyUserIds(reqData *dto.StatisticsCrmStatsGroupbyUserIdsRequestDTO) (res *dto.StatisticsCrmStatsGroupbyUserIdsResponseDTO, err error) {
	res = dto.NewStatisticsCrmStatsGroupbyUserIdsResponseDTO()
	err = c.RequestByAutoResolve(common.StatisticsCrmStatsGroupbyUserIds, reqData, res)
	return
}
