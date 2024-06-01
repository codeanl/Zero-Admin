package model

import (
	"gorm.io/gorm"
)

type (
	CategoryModel interface {
		AddCategory(role *Category) (err error)
		UpdateCategory(id int64, role *Category) error
		DeleteCategoryByIds(ids []int64) error
		GetCategoryList() ([]Category, int64, error)
		GetCategoryById(id int64) (role *Category, err error)
		GetCategoryByParentId(id int64) (info []*Category, err error)
	}

	defaultCategoryModel struct {
		conn *gorm.DB
	}
	Category struct {
		gorm.Model
		Name         string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"`           //角色名称
		ParentId     int64  `json:"parent_id" gorm:"type:bigint;comment:上机分类的编号;not null"`         //上机分类的编号
		Level        string `json:"level" gorm:"type:varchar(191);comment:分类级别;not null"`          //分类级别：1->1级；2->2级
		ProductCount int64  `json:"product_count" gorm:"type:bigint;comment:商品数量;not null"`        //商品数量
		ProductUnit  string `json:"product_unit" gorm:"type:varchar(191);comment:商品单位;not null"`   //商品单位
		NavStatus    string `json:"nav_status" gorm:"type:varchar(191);comment:是否显示在导航栏;not null"` //是否显示在导航栏0->不是；1->是'
		ShowStatus   string `json:"show_status" gorm:"type:varchar(191);comment:展示状态;not null"`    //展示状态0->不是；1->是'
		Sort         int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`                   //排序
		Icon         string `json:"icon" gorm:"type:varchar(191);comment:图标;not null"`             //图标
		Keywords     string `json:"keywords" gorm:"type:varchar(191);comment:关键字;not null"`        //关键字
		Description  string `json:"description" gorm:"type:varchar(191);comment:描述;not null"`      //描述
	}
)

func NewCategoryModel(conn *gorm.DB) CategoryModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Category{})
	return &defaultCategoryModel{
		conn: conn,
	}
}
func (m *defaultCategoryModel) AddCategory(role *Category) (err error) {
	return m.conn.Model(&Category{}).Create(role).Error
}
func (m *defaultCategoryModel) GetCategoryById(id int64) (role *Category, err error) {
	err = m.conn.Model(&Category{}).Where("id=?", id).First(&role).Error
	return role, err
}
func (m *defaultCategoryModel) GetCategoryByParentId(id int64) (info []*Category, err error) {
	err = m.conn.Model(&Category{}).Where("parent_id = ?", id).Find(&info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultCategoryModel) UpdateCategory(id int64, role *Category) error {
	err := m.conn.Model(&Category{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultCategoryModel) DeleteCategoryByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Category{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultCategoryModel) GetCategoryList() ([]Category, int64, error) {
	var list []Category
	db := m.conn.Model(&Category{}).Order("sort ASC").Find(&list)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	return list, total, err
}
