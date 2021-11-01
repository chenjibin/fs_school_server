package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SchoolRouter struct {
}

func (s SysRouter) InitSchoolRouter(Router *gin.RouterGroup) {
	schoolRouter := Router.Group("school").Use(middleware.OperationRecord())
	var schoolApi = v1.ApiGroupApp.SystemApiGroup.SchoolApi
	{
		schoolRouter.POST("create", schoolApi.CreateSchool)
		schoolRouter.PUT("update", schoolApi.UpdateSchool)
		schoolRouter.POST("list", schoolApi.GetSchoolList)
		schoolRouter.DELETE("delete", schoolApi.DeleteSchool)
	}
}

