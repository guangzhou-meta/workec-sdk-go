package dto

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
)

type StatisticsDigitalMapPhoneRequestDTO struct {
	Sequential     int               `json:"sequential"`
	BusinessIndexs []int             `json:"businessIndexs"`
	CorpId         int               `json:"corpId"`
	UserId         int               `json:"userId"`
	ApiParams      map[string]string `json:"apiParams"`
}

type StatisticsDigitalMapPhoneResponseDTO struct {
	CommonDTO
	Data *StatisticsDigitalMapPhoneResponseDTOData `json:"data"`
}

type StatisticsDigitalMapPhoneResponseDTOData struct {
	CharData *StatisticsDigitalMapPhoneResponseDTODataCharData `json:"charData"`
}

type StatisticsDigitalMapPhoneResponseDTODataCharData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	PreValue   int     `json:"preValue"`
	NowValue   int     `json:"nowValue"`
	Trend      int     `json:"trend"`
	Percentage float64 `json:"percentage"`
}

func NewStatisticsDigitalMapPhoneRequestDTO() *StatisticsDigitalMapPhoneResponseDTO {
	return &StatisticsDigitalMapPhoneResponseDTO{}
}

type StatisticsLineGraphPhoneRequestDTO struct {
	BusinessIndexs []int             `json:"businessIndexs"`
	CorpId         int               `json:"corpId"`
	UserId         int               `json:"userId"`
	ApiParams      map[string]string `json:"apiParams"`
}

type StatisticsLineGraphPhoneResponseDTO struct {
	CommonDTO
	Data *StatisticsLineGraphPhoneResponseDTOData `json:"data"`
}

type StatisticsLineGraphPhoneResponseDTOData struct {
	ChartData *StatisticsLineGraphPhoneResponseDTODataChartData `json:"chartData"`
	Index     []int                                             `json:"index"`
}
type StatisticsLineGraphPhoneResponseDTODataChartData struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value []int  `json:"value"`
}

func NewStatisticsLineGraphPhoneRequestDTO() *StatisticsLineGraphPhoneResponseDTO {
	return &StatisticsLineGraphPhoneResponseDTO{}
}

type StatisticsDigitalMapWorkefficRequestDTO struct {
	CorpId         int64                        `json:"corpId"`
	UserId         int64                        `json:"userId"`
	BusinessIndexs []int                        `json:"businessIndexs"`
	Sequential     *common.StatisticsSequential `json:"sequential"`
	ApiParams      map[string]interface{}       `json:"apiParams"`
}

type StatisticsDigitalMapWorkefficResponseDTO struct {
	CommonDTO
	Data *StatisticsDigitalMapWorkefficResponseDTOData `json:"data"`
}

type StatisticsDigitalMapWorkefficResponseDTOData struct {
	ChartData *StatisticsDigitalMapWorkefficResponseDTOData `json:"chartData"`
}

type StatisticsDigitalMapWorkefficResponseDTODataChartData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	PreValue   int     `json:"preValue"`
	NowValue   int     `json:"nowValue"`
	Trend      int     `json:"trend"`
	Percentage float64 `json:"percentage"`
}

func NewStatisticsDigitalMapWorkefficRequestDTO() *StatisticsDigitalMapWorkefficResponseDTO {
	return &StatisticsDigitalMapWorkefficResponseDTO{}
}

type StatisticsHistogramWorkefficRequestDTO struct {
	CorpId         int64                  `json:"corpId"`
	UserId         int64                  `json:"userId"`
	BusinessIndexs []int                  `json:"businessIndexs"`
	ApiParams      map[string]interface{} `json:"apiParams"`
}

type StatisticsHistogramWorkefficResponseDTO struct {
	CommonDTO
	Data *StatisticsHistogramWorkefficResponseDTOData `json:"data"`
}

type StatisticsHistogramWorkefficResponseDTOData struct {
	Index     []string   `json:"index"`
	CharyData []struct{} `json:"charyData"`
}

func NewStatisticsHistogramWorkefficRequestDTO() *StatisticsHistogramWorkefficResponseDTO {
	return &StatisticsHistogramWorkefficResponseDTO{}
}

type StatisticsDigitalMapTagRequestDTO struct {
	Sequential     int                    `json:"sequential"`
	BusinessIndexs []int                  `json:"businessIndexs"`
	CorpId         int                    `json:"corpId"`
	UserId         int                    `json:"userId"`
	ApiParams      map[string]interface{} `json:"apiParams"`
}

type StatisticsDigitalMapTagResponseDTO struct {
	CommonDTO
	Data *StatisticsDigitalMapTagResponseDTOData `json:"data"`
}

type StatisticsDigitalMapTagResponseDTOData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	PreValue   int     `json:"preValue"`
	NowValue   int     `json:"nowValue"`
	Trend      int     `json:"trend"`
	Percentage float64 `json:"percentage"`
}

func NewStatisticsDigitalMapTagRequestDTO() *StatisticsDigitalMapTagResponseDTO {
	return &StatisticsDigitalMapTagResponseDTO{}
}

type StatisticsHistogramTagRequestDTO struct {
	BusinessIndexs []int                  `json:"businessIndexs"`
	CorpId         int                    `json:"corpId"`
	UserId         int                    `json:"userId"`
	ApiParams      map[string]interface{} `json:"apiParams"`
}

type StatisticsHistogramTagResponseDTO struct {
	CommonDTO
	Data *StatisticsHistogramTagResponseDTOData `json:"data"`
}

type StatisticsHistogramTagResponseDTOData struct {
	ChartData *StatisticsHistogramTagResponseDTODataChartData `json:"chartData"`
	Index     []int                                           `json:"index"`
}

type StatisticsHistogramTagResponseDTODataChartData struct {
	Value []int `json:"value"`
}

func NewStatisticsHistogramTagRequestDTO() *StatisticsHistogramTagResponseDTO {
	return &StatisticsHistogramTagResponseDTO{}
}

type StatisticsDigitalMapCrmQuantityRequestDTO struct {
	Sequential     int                    `json:"sequential"`
	BusinessIndexs []int                  `json:"businessIndexs"`
	CorpId         int                    `json:"corpId"`
	UserId         int                    `json:"userId"`
	ApiParams      map[string]interface{} `json:"apiParams"`
}

type StatisticsDigitalMapCrmQuantityResponseDTO struct {
	CommonDTO
	Data *StatisticsDigitalMapCrmQuantityResponseDTOData `json:"data"`
}

type StatisticsDigitalMapCrmQuantityResponseDTOData struct {
	ChartData *StatisticsDigitalMapCrmQuantityResponseDTODataChartData `json:"chartData	"`
}

type StatisticsDigitalMapCrmQuantityResponseDTODataChartData struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	PreValue   int     `json:"preValue"`
	NowValue   int     `json:"nowValue"`
	Trend      int     `json:"trend"`
	Percentage float64 `json:"percentage"`
}

func NewStatisticsDigitalMapCrmQuantityRequestDTO() *StatisticsDigitalMapCrmQuantityResponseDTO {
	return &StatisticsDigitalMapCrmQuantityResponseDTO{}
}

type StatisticsCrmStatsWueryStepCountByChannelRequestDTO struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type StatisticsCrmStatsWueryStepCountByChannelResponseDTO struct {
	CommonDTO
	Len  int                                                       `json:"len"`
	Data *StatisticsCrmStatsWueryStepCountByChannelResponseDTOData `json:"data"`
}

type StatisticsCrmStatsWueryStepCountByChannelResponseDTOData struct {
	CorpId     int64  `json:"corpId"`
	CreateDate string `json:"createDate"`
	Step       int    `json:"step"`
	Count      int    `json:"count"`
	UserId     int64  `json:"userId"`
	Channelld  int64  `json:"channelld"`
}

func NewStatisticsCrmStatsWueryStepCountByChannelRequestDTO() *StatisticsCrmStatsWueryStepCountByChannelResponseDTO {
	return &StatisticsCrmStatsWueryStepCountByChannelResponseDTO{}
}

type StatisticsCrmStatsGetTopStepCountByChannelRequestDTO struct {
	UserIdList  []int                                                     `json:"userIdList"`
	TopNum      int                                                       `json:"topNum"`
	StartStepId int                                                       `json:"startStepId"`
	EndStepId   int                                                       `json:"endStepId"`
	Date        *StatisticsCrmStatsGetTopStepCountByChannelRequestDTODate `json:"date"`
}

type StatisticsCrmStatsGetTopStepCountByChannelRequestDTODate struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type StatisticsCrmStatsGetTopStepCountByChannelResponseDTO struct {
	CommonDTO
	Len  int                                                        `json:"len"`
	Data *StatisticsCrmStatsGetTopStepCountByChannelResponseDTOData `json:"data"`
}

type StatisticsCrmStatsGetTopStepCountByChannelResponseDTOData struct {
	Channelld         int     `json:"channelld"`
	StartStepCount    int     `json:"startStepCount"`
	EndStepCount      int     `json:"endStepCount"`
	ConversionPercent float64 `json:"conversionPercent"`
}

func NewStatisticsCrmStatsGetTopStepCountByChannelRequestDTO() *StatisticsCrmStatsGetTopStepCountByChannelResponseDTO {
	return &StatisticsCrmStatsGetTopStepCountByChannelResponseDTO{}
}

type StatisticsCrmStatsGetStepCountRequestDTO struct {
	UserIdList []int                                         `json:"userIdList"`
	Date       *StatisticsCrmStatsGetStepCountRequestDTODate `json:"date"`
}

type StatisticsCrmStatsGetStepCountRequestDTODate struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type StatisticsCrmStatsGetStepCountResponseDTO struct {
	CommonDTO
	Len  int                                            `json:"len"`
	Data *StatisticsCrmStatsGetStepCountResponseDTOData `json:"data	"`
}

type StatisticsCrmStatsGetStepCountResponseDTOData struct {
	Step  int `json:"step"`
	Count int `json:"count"`
}

func NewStatisticsCrmStatsGetStepCountRequestDTO() *StatisticsCrmStatsGetStepCountResponseDTO {
	return &StatisticsCrmStatsGetStepCountResponseDTO{}
}

type StatisticsCrmStatsGroupbyUserIdsRequestDTO struct {
	Ids       []int `json:"ids"`
	QueryType int   `json:"queryType"`
}

type StatisticsCrmStatsGroupbyUserIdsResponseDTO struct {
	CommonDTO
	Len  int                                              `json:"len"`
	Data *StatisticsCrmStatsGroupbyUserIdsResponseDTOData `json:"data"`
}

type StatisticsCrmStatsGroupbyUserIdsResponseDTOData struct {
	UserId    int `json:"userId"`
	UserName  int `json:"userName"`
	GroupName int `json:"groupName"`
	Num       int `json:"num"`
}

func NewStatisticsCrmStatsGroupbyUserIdsRequestDTO() *StatisticsCrmStatsGroupbyUserIdsResponseDTO {
	return &StatisticsCrmStatsGroupbyUserIdsResponseDTO{}
}
