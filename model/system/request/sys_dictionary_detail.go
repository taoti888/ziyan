package request

import (
	"github.com/taoti888/ziyan/model/common/request"
	"github.com/taoti888/ziyan/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
