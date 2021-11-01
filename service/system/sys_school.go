package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type SchoolService struct {
}


//@author: [chenjibin](https://github.com/chenjibin)
//@function: CreateSchool
//@description: 创建机构
//@param: u model.SysSchool
//@return: err error, schoolInter model.SysSchool

func (schoolService *SchoolService) CreateSchool(s system.SysSchool) (err error,schoolInter system.SysSchool) {
	var sch system.SysSchool
	if !errors.Is(global.GVA_DB.Where("name = ?", s.Name).First(&sch).Error, gorm.ErrRecordNotFound) {
		return errors.New("机构名称已经存在"), schoolInter
	}
	err = global.GVA_DB.Create(&s).Error
	return err, s
}

func (schoolService *SchoolService) GetSchoolList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysSchool{})
	var schoolList []system.SysSchool
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&schoolList).Error
	return err, schoolList, total
}

func (schoolService *SchoolService) UpdateSchool(sch system.SysSchool) (err error)  {
	//err = global.GVA_DB.Updates(&sch).Error
	err = global.GVA_DB.Model(&sch).Select("Name", "Address", "Principal", "PrincipalMobile", "Type", "Logo").Updates(system.SysSchool{
		Name:            sch.Name,
		Address:         sch.Address,
		Principal:       sch.Principal,
		PrincipalMobile: sch.PrincipalMobile,
		Type:            sch.Type,
		Logo:            sch.Logo,
	}).Error
	return err
}

func (schoolService *SchoolService) DeleteSchool(sch system.SysSchool) (err error) {
	err = global.GVA_DB.Delete(&sch).Error
	return err
}
