package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CourseApi struct {
}

func (course *CourseApi) CreateCourse(c *gin.Context) {
	var cou systemReq.CourseAdd
	_ = c.ShouldBindJSON(&cou)
	if err := utils.Verify(cou, utils.CourseCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	courseInfo := &system.SysCourse{
		Name:       cou.Name,
		CourseHour: cou.CourseHour,
		CourseImg:  cou.CourseImg,
		CourseDesc: cou.CourseDesc,
		SchoolId:   cou.SchoolId,
	}
	err, courseReturn := courseService.CreateCourse(*courseInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithDetailed(systemRes.SysCourseResponse{Course: courseReturn}, "创建失败!", c)
	} else {
		response.OkWithDetailed(systemRes.SysCourseResponse{Course: courseReturn}, "创建成功!", c)
	}
}
