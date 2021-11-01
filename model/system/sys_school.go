package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysSchool struct {
	global.GVA_MODEL
	Name string `json:"name" gorm:"comment:机构名称"`
	Address string `json:"address" gorm:"comment:机构地址"`
	Principal string `json:"principal" gorm:"comment:机构负责人"`
	PrincipalMobile string `json:"principalMobile" gorm:"comment:机构负责人电话"`
	Type string `json:"type" gorm:"comment:机构类型"`
	Logo string `json:"logo" gorm:"comment:机构LOGO"`
}
