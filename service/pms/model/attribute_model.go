package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	AttributeModel interface {
		AddAttribute(role *Attribute) (err error)
		UpdateAttribute(id int64, role *Attribute) error
		DeleteAttributeByIds(ids []int64) error
		GetAttributeList(in *pms.AttributeListReq) ([]*Attribute, int64, error)
	}

	defaultAttributeModel struct {
		conn *gorm.DB
	}
	Attribute struct {
		gorm.Model
		Name       string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`      //名称
		CategoryID int64  `json:"category_id" gorm:"type:bigint;comment:属性分类id;not null"` //分类id
		Type       string `json:"type" gorm:"type:varchar(191);comment:属性的类型;not null"`   //属性的类型；1->基本属性；2->销售属性'
	}
)

func NewAttributeModel(conn *gorm.DB) AttributeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Attribute{})
	return &defaultAttributeModel{
		conn: conn,
	}
}
func (m *defaultAttributeModel) AddAttribute(role *Attribute) (err error) {
	return m.conn.Model(&Attribute{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultAttributeModel) UpdateAttribute(id int64, role *Attribute) error {
	err := m.conn.Model(&Attribute{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultAttributeModel) DeleteAttributeByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Attribute{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultAttributeModel) GetAttributeList(in *pms.AttributeListReq) ([]*Attribute, int64, error) {
	var list []*Attribute
	db := m.conn.Model(&Attribute{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Type != "" {
		db = db.Where("type LIKE ?", fmt.Sprintf("%%%s%%", in.Type))
	}
	if in.ProductAttributeCategoryId != 0 {
		db = db.Where("product_attribute_category_id = ?", fmt.Sprintf("%%%s%%", in.ProductAttributeCategoryId))
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
