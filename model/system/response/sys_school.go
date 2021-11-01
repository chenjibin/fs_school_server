package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysSchoolResponse struct {
	School system.SysSchool `json:"school"`
}
