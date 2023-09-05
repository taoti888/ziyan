package response

import (
	"github.com/taoti888/ziyan/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
