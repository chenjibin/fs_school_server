package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type CourseService struct {
}

func (courseService *CourseService) CreateCourse(c system.SysCourse) (err error, courseInter system.SysCourse) {
	err = global.GVA_DB.Create(&c).Error
	return err, c
}
