package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysCourse struct {
	global.GVA_MODEL
	Name       string    `json:"name" gorm:"comment:课程名称"`
	CourseHour int       `json:"courseHour" gorm:"column:course_hour;comment:课时"`
	CourseImg  string    `json:"courseImg" gorm:"column:course_img;comment:课程封面图"`
	CourseDesc string    `json:"courseDesc" gorm:"column:course_desc;comment:课程简介"`
	SchoolId   int       `json:"schoolId" gorm:"column:school_id;comment:所属机构ID"`
	School     SysSchool `json:"school" gorm:"foreignKey:SchoolId;references:ID"`
}
