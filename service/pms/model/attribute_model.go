package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	AttributeModel interface {
		AddAttribute(attribute *Attribute) (*Attribute, error)
		UpdateAttribute(id int64, role *Attribute) error
		DeleteAttributeByIds(ids []int64) error
		GetAttributeList(in *pms.AttributeListReq) ([]Attribute, int64, error)
		GetAttributeByID(id int64) (info Attribute, err error)
	}

	defaultAttributeModel struct {
		conn *gorm.DB
	}
	Attribute struct {
		gorm.Model
		AttributeCategoryID int64  `json:"attribute_category_id" gorm:"type:bigint;comment:属性分类id;not null"` //分类id
		Name                string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`                //名称
		Type                string `json:"type" gorm:"type:varchar(191);comment:类型;not null"`                //1->商品属性 2->规格
		Value               string `json:"value" gorm:"type:varchar(255);comment:可选值列表;not null"`            //'可选值列表，以逗号隔开'
		Sort                int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`                      //排序

	}
)

func NewAttributeModel(conn *gorm.DB) AttributeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Attribute{})
	return &defaultAttributeModel{
		conn: conn,
	}
}
func (m *defaultAttributeModel) AddAttribute(attribute *Attribute) (*Attribute, error) {
	err := m.conn.Model(&Attribute{}).Create(attribute).Error
	return attribute, err
}
func (m *defaultAttributeModel) GetAttributeByID(id int64) (info Attribute, err error) {
	err = m.conn.Model(&Attribute{}).Where("id=?", id).Find(&info).Error
	return info, err
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
func (m *defaultAttributeModel) GetAttributeList(in *pms.AttributeListReq) ([]Attribute, int64, error) {
	var list []Attribute
	db := m.conn.Model(&Attribute{}).Order("sort ASC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}

	if in.AttributeCategoryID != 0 {
		db = db.Where("attribute_category_id = ?", in.AttributeCategoryID)
	}
	if in.Type != "" {
		db = db.Where("type = ?", in.Type)
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
