package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CourseRouter struct {
}

func (s SysRouter) InitCourseRouter(Router *gin.RouterGroup) {
	courseRouter := Router.Group("course").Use(middleware.OperationRecord())
	var courseApi = v1.ApiGroupApp.SystemApiGroup.CourseApi
	{
		courseRouter.POST("create", courseApi.CreateCourse)
		//courseRouter.PUT("update", courseApi.UpdateSchool)
		//courseRouter.POST("list", courseApi.GetSchoolList)
		//courseRouter.DELETE("delete", courseApi.DeleteSchool)
	}
}
