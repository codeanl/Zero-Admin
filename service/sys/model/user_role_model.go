package model

import (
	"gorm.io/gorm"
)

type (
	UserRoleModel interface {
		AddUserRole(userrole *UserRole) (err error)
		UpdateUserRole(userrole *UserRole) error
		DeleteByUserID(userid int64) error
		GetUserRoleByUserID(userid int64) ([]UserRole, error)
		IsSuperAdmin(userid int64) (exist bool, err error)
	}
	defaultUserRoleModel struct {
		conn *gorm.DB
	}
	UserRole struct {
		gorm.Model
		UserID   int64  `json:"user_id" gorm:"type:bigint;comment:用户ID;not null"`        //用户id
		RoleID   int64  `json:"role_id" gorm:"type:bigint;comment:角色ID;not null"`        //角色id
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建者;not null"` //创建者
	}
)

func NewUserRoleModel(conn *gorm.DB) UserRoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&UserRole{})
	return &defaultUserRoleModel{
		conn: conn,
	}
}

//AddUserRole 添加用户角色
func (m *defaultUserRoleModel) AddUserRole(userrole *UserRole) (err error) {
	return m.conn.Model(&UserRole{}).Create(userrole).Error
}

//UpdateUserRole 更新用户角色
func (m *defaultUserRoleModel) UpdateUserRole(userrole *UserRole) error {
	err := m.conn.Model(&UserRole{}).Where("user_id = ?", userrole.UserID).Updates(userrole).Error
	return err
}
func (m *defaultUserRoleModel) DeleteByUserID(userid int64) error {
	var userrole *UserRole
	err := m.conn.Model(&UserRole{}).Where("user_id = ?", userid).Delete(&userrole).Error
	return err
}
func (m *defaultUserRoleModel) GetUserRoleByUserID(userid int64) ([]UserRole, error) {
	var userrole []UserRole
	err := m.conn.Model(&UserRole{}).Where("user_id = ?", userid).Find(&userrole).Error
	return userrole, err
}
func (m *defaultUserRoleModel) IsSuperAdmin(userid int64) (exist bool, err error) {
	var count int64
	err = m.conn.Model(&UserRole{}).Where("user_id=? && role_id=?", userid, 1).Count(&count).Error
	if count == 0 {
		return false, err
	}
	return true, nil
}
