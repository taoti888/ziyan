package request

import (
	"github.com/taoti888/ziyan/model/common/request"
	"github.com/taoti888/ziyan/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
