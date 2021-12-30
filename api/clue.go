package api

import (
	"github.com/guangzhou-meta/workec-sdk-go/common"
	"github.com/guangzhou-meta/workec-sdk-go/dto"
)

// ClueImport 录入线索
func (c *Config) ClueImport(reqData *dto.ClueImportRequestDTO) (res *dto.CommonDTO, err error) {
	res = dto.NewCommonDTO()
	err = c.RequestByAutoResolve(common.ClueImport, reqData, res)
	return
}
