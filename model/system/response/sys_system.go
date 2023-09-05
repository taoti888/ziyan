package response

import "github.com/taoti888/ziyan/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
