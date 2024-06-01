package model

import (
	"gorm.io/gorm"
)

type (
	RoleMenuModel interface {
		GetMenuByRoleId(roleId int64) ([]RoleMenu, error)
		DeleteByRoleId(id int64) error
		AddRoleMenu(rolemenu *RoleMenu) (err error)
	}

	defaultRoleMenuModel struct {
		conn *gorm.DB
	}
	RoleMenu struct {
		gorm.Model
		RoleID   int64  `json:"role_id" gorm:"type:bigint;comment:角色ID;not null"`        //角色ID
		MenuID   int64  `json:"menu_id" gorm:"type:bigint;comment:菜单ID;not null"`        //菜单ID
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
	}
)

func NewRoleMenuModel(conn *gorm.DB) RoleMenuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&RoleMenu{})
	return &defaultRoleMenuModel{
		conn: conn,
	}
}

func (m *defaultRoleMenuModel) GetMenuByRoleId(roleId int64) ([]RoleMenu, error) {
	var rolemenus []RoleMenu
	err := m.conn.Model(&RoleMenu{}).Where("role_id=?", roleId).Find(&rolemenus).Error
	return rolemenus, err
}
func (m *defaultRoleMenuModel) DeleteByRoleId(id int64) error {
	err := m.conn.Where("role_id=?", id).Delete(&RoleMenu{}).Error
	return err
}
func (m *defaultRoleMenuModel) AddRoleMenu(rolemenu *RoleMenu) (err error) {
	return m.conn.Model(&RoleMenu{}).Create(rolemenu).Error
}
