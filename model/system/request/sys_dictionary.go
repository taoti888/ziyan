package request

import (
	"github.com/taoti888/ziyan/model/common/request"
	"github.com/taoti888/ziyan/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
