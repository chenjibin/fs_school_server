package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SchoolApi struct {

}

// @Tags SysSchool
// @Summary 创建机构
// @Produce  application/json
// @Param data body systemReq.Register true "机构名称，机构地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /school/create [post]
func (s *SchoolApi) CreateSchool(c *gin.Context) {
	var sch systemReq.SchoolAdd
	_ = c.ShouldBindJSON(&sch)
	if err := utils.Verify(sch, utils.SchoolCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	school := &system.SysSchool{
		Name:            sch.Name,
		Address:         sch.Address,
		Principal:       sch.Principal,
		PrincipalMobile: sch.PrincipalMobile,
		Type:            sch.Type,
		Logo:            sch.Logo,
	}
	err, schoolReturn := schoolService.CreateSchool(*school)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Any("err", err))
		response.FailWithDetailed(systemRes.SysSchoolResponse{School: schoolReturn}, "创建失败!", c)
	} else {
		response.OkWithDetailed(systemRes.SysSchoolResponse{School: schoolReturn}, "创建成功!", c)
	}
}

// @Tags SysSchool
// @Summary 更新机构
// @Produce  application/json
// @Param data body systemReq.Register true "机构名称，机构地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /school/update [post]
func (s *SchoolApi) UpdateSchool(c *gin.Context) {
	var school system.SysSchool
	_ = c.ShouldBindJSON(&school)
	if err := utils.Verify(school, utils.SchoolCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := schoolService.UpdateSchool(school); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysSchool
// @Summary 删除机构
// @Produce  application/json
// @Param data body systemReq.Register true "机构名称，机构地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /school/delete [delete]
func (s *SchoolApi) DeleteSchool(c *gin.Context) {
	var school system.SysSchool
	_ = c.ShouldBindJSON(&school)
	if err := utils.Verify(school, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	if err := schoolService.DeleteSchool(school); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysSchool
// @Summary 获取机构列表
// @Produce  application/json
// @Param data body systemReq.Register true "机构名称，机构地址"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功！"}"
// @Router /school/list [post]
func (s *SchoolApi) GetSchoolList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := schoolService.GetSchoolList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败!", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
