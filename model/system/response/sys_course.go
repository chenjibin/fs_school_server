package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysCourseResponse struct {
	Course system.SysCourse `json:"course"`
}
