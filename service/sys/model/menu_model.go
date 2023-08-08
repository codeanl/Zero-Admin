package model

import (
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"gorm.io/gorm"
)

type (
	MenuModel interface {
		AddMenu(menu *Menu) (err error)
		GetMenuList(in *sys.MenuListReq) ([]Menu, int64, error)
		UpdateMenu(id int64, menu *Menu) error
		DeleteMenuByIds(ids []int64) error
		//GetMenusByUserID(userID int64) ([]Menu, error)
	}

	defaultMenuModel struct {
		conn *gorm.DB
	}
	Menu struct {
		gorm.Model
		Name     string `json:"name" gorm:"type:varchar(191);comment:菜单名称;not null"`     //菜单名称
		ParentID int64  `json:"parent_id" gorm:"type:bigint;comment:父菜单ID;not null;"`    //父菜单ID
		Url      string `json:"url" gorm:"type:varchar(191);comment:路径;not null"`        //路径
		Type     string `json:"type" gorm:"type:varchar(191);comment:类型;not null"`       //类型  1：目录  2：菜单   3：按钮',
		Icon     string `json:"icon" gorm:"type:varchar(191);comment:菜单图标;not null"`     //菜单图标
		OrderNum int64  `json:"order_num" gorm:"type:bigint;comment:排序;not null"`        //排序
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
		UpdateBy string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"` //更新人
		Remark   string `json:"remark" gorm:"type:varchar(191);comment:备注;not null"`     //备注
		TAG      string `json:"tag" gorm:"type:varchar(191);comment:标识;not null"`        //标识
	}
)

func NewMenuModel(conn *gorm.DB) MenuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Menu{})
	return &defaultMenuModel{
		conn: conn,
	}
}
func (m *defaultMenuModel) AddMenu(menu *Menu) (err error) {
	return m.conn.Model(&Menu{}).Create(menu).Error
}

func (m *defaultMenuModel) GetMenuList(in *sys.MenuListReq) ([]Menu, int64, error) {
	var list []Menu
	db := m.conn.Model(&Menu{}).Order("order_num ASC")
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Find(&list).Error
	return list, total, err
}
func (m *defaultMenuModel) UpdateMenu(id int64, menu *Menu) error {
	err := m.conn.Model(&Menu{}).Where("id=?", id).Updates(menu).Error
	return err
}
func (m *defaultMenuModel) DeleteMenuByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Menu{}).Error
	return err
}

//func (m *defaultMenuModel) GetMenusByUserID(userID int64) ([]Menu, error) {
//	var userrole AdminRole
//	if err := m.conn.Model(&AdminRole{}).Where("user_id=?", userID).First(&userrole).Error; err != nil {
//		return nil, err
//	}
//	var role Role
//	if err := m.conn.Model(&Role{}).First(&role, userrole.RoleID).Error; err != nil {
//		return nil, err
//	}
//	var roleMenus []RoleMenu
//	if err := m.conn.Model(&RoleMenu{}).Where("role_id = ?", role.ID).Find(&roleMenus).Error; err != nil {
//		return nil, err
//	}
//	var menuIDs []uint
//	for _, rm := range roleMenus {
//		menuIDs = append(menuIDs, uint(rm.MenuID))
//	}
//	var menus []Menu
//	if err := m.conn.Model(&Menu{}).Where("id IN (?)", menuIDs).Find(&menus).Error; err != nil {
//		return nil, err
//	}
//	return menus, nil
//}
