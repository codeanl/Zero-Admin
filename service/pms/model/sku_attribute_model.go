package model

import (
	"gorm.io/gorm"
)

type (
	SkuAttributeModel interface {
		AddSkuAttribute(info *SkuAttribute) (*SkuAttribute, error)
		UpdateSkuAttribute(id int64, role *SkuAttribute) error
		DeleteSkuAttributeBySpuID(id int64) error
		GetAttributeValueByProductID(id int64) ([]SkuAttribute, error)
	}

	defaultSkuAttributeModel struct {
		conn *gorm.DB
	}
	SkuAttribute struct {
		gorm.Model
		SkuID          int64 `json:"sku_id" gorm:"type:bigint;comment:名称;not null"`                //商品id
		SkuSizeValueID int64 `json:"sku_size_value_id" gorm:"type:bigint;comment:属性分类id;not null"` //属性值id
	}
)

func NewSkuAttributeModel(conn *gorm.DB) SkuAttributeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&SkuAttribute{})
	return &defaultSkuAttributeModel{
		conn: conn,
	}
}
func (m *defaultSkuAttributeModel) AddSkuAttribute(info *SkuAttribute) (*SkuAttribute, error) {
	err := m.conn.Model(&SkuAttribute{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultSkuAttributeModel) UpdateSkuAttribute(id int64, role *SkuAttribute) error {
	err := m.conn.Model(&SkuAttribute{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultSkuAttributeModel) DeleteSkuAttributeBySpuID(id int64) error {
	err := m.conn.Model(&SkuAttribute{}).Where("sku_id=?", id).Delete(&SkuAttribute{}).Error
	return err
}
func (m *defaultSkuAttributeModel) GetAttributeValueByProductID(id int64) ([]SkuAttribute, error) {
	var info []SkuAttribute
	err := m.conn.Model(&SkuAttribute{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
