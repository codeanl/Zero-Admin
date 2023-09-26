package model

import (
	"gorm.io/gorm"
)

type (
	AttributeCategoryModel interface {
		AddAttributeCategory(AttributeCategory *AttributeCategory) (*AttributeCategory, error)
		UpdateAttributeCategory(id int64, role *AttributeCategory) error
		DeleteAttributeCategoryByIds(ids []int64) error
		GetAttributeCategoryList() ([]AttributeCategory, int64, error)
		GetAttributeCategoryByID(id int64) (info AttributeCategory, err error)
	}

	defaultAttributeCategoryModel struct {
		conn *gorm.DB
	}
	AttributeCategory struct {
		gorm.Model
		Name       string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`     //名称
		ParentId   int64  `json:"parent_id" gorm:"type:bigint;comment:上机分类的编号;not null"` //上机分类的编号
		MerchantID int64  `json:"merchant_id" gorm:"type:bigint;comment:商家id;not null"`  //商家id
	}
)

func NewAttributeCategoryModel(conn *gorm.DB) AttributeCategoryModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&AttributeCategory{})
	return &defaultAttributeCategoryModel{
		conn: conn,
	}
}
func (m *defaultAttributeCategoryModel) AddAttributeCategory(info *AttributeCategory) (*AttributeCategory, error) {
	err := m.conn.Model(&AttributeCategory{}).Create(info).Error
	return info, err
}
func (m *defaultAttributeCategoryModel) GetAttributeCategoryByID(id int64) (info AttributeCategory, err error) {
	err = m.conn.Model(&AttributeCategory{}).Where("id=?", id).Find(&info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultAttributeCategoryModel) UpdateAttributeCategory(id int64, role *AttributeCategory) error {
	err := m.conn.Model(&AttributeCategory{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultAttributeCategoryModel) DeleteAttributeCategoryByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&AttributeCategory{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultAttributeCategoryModel) GetAttributeCategoryList() ([]AttributeCategory, int64, error) {
	var list []AttributeCategory
	db := m.conn.Model(&AttributeCategory{}).Order("created_at DESC")
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Find(&list).Error
	return list, total, err
}
