package model

import (
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		AddRole(role *Role) (err error)
		UpdateRole(id int64, role *Role) error
		DeleteRoleByIds(ids []int64) error
		GetRoleList(in *sys.RoleListReq) ([]*Role, int64, error)
		GetRoleByUserID(userid int64) ([]Role, error)
		GetRoleByUserList(in *sys.RoleByUserListReq) ([]Role, error)
	}

	defaultRoleModel struct {
		conn *gorm.DB
	}
	Role struct {
		gorm.Model
		Name     string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"`     //角色名称
		Remark   string `json:"remark" gorm:"type:varchar(191);comment:备注;not null"`     //备注
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
		UpdateBy string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"` //更新人
	}
)

func NewRoleModel(conn *gorm.DB) RoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Role{})
	return &defaultRoleModel{
		conn: conn,
	}
}
func (m *defaultRoleModel) AddRole(role *Role) (err error) {
	return m.conn.Model(&Role{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultRoleModel) UpdateRole(id int64, role *Role) error {
	err := m.conn.Model(&Role{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultRoleModel) DeleteRoleByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Role{}).Error
	// 手动删除相关的 userrole 数据
	err = m.conn.Where("role_id IN (?)", ids).Delete(&UserRole{}).Error
	return err
	return err
}

// GetRoleByUserID 通过用户id查询用户的所有角色
func (m *defaultRoleModel) GetRoleByUserID(userid int64) ([]Role, error) {
	var userrole []UserRole
	if err := m.conn.Model(&UserRole{}).Where("user_id=?", userid).Find(&userrole).Error; err != nil {
		return nil, err
	}
	var roles []Role
	for _, item := range userrole {
		var role Role
		if err := m.conn.Model(&Role{}).Where("id = ?", item.RoleID).First(&role).Error; err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

//GetUserList 获取用户列表
func (m *defaultRoleModel) GetRoleList(in *sys.RoleListReq) ([]*Role, int64, error) {
	var list []*Role
	db := m.conn.Model(&Role{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if in.Current > 0 && in.PageSize > 0 {
		err = db.Offset(int((in.Current - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
func (m *defaultRoleModel) GetRoleByUserList(in *sys.RoleByUserListReq) ([]Role, error) {
	var userRole []UserRole
	err := m.conn.Model(&UserRole{}).Where("user_id=?", in.UserID).Find(&userRole).Error

	var roles []Role
	for _, userRole := range userRole {
		var role Role
		err := m.conn.Model(&Role{}).Where("id = ?", userRole.RoleID).First(&role).Error
		if err != nil {
			// 处理查询角色信息的错误
		}
		// 将角色信息存储在 roles 数组中
		roles = append(roles, role)
	}
	return roles, err
}
