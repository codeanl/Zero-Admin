package model

import (
	"gorm.io/gorm"
)

type (
	SpuAttributeModel interface {
		AddSpuAttribute(SpuAttribute *SpuAttribute) (*SpuAttribute, error)
		UpdateSpuAttribute(id int64, role *SpuAttribute) error
		DeleteSpuAttributeBySpuID(id int64) error
		GetAttributeValueByProductID(id int64) ([]SpuAttribute, error)
	}

	defaultSpuAttributeModel struct {
		conn *gorm.DB
	}
	SpuAttribute struct {
		gorm.Model
		ProductID        int64 `json:"product_id" gorm:"type:bigint;comment:名称;not null"`             //商品id
		AttributeValueID int64 `json:"attribute_value_id" gorm:"type:bigint;comment:属性分类id;not null"` //属性值id
	}
)

func NewSpuAttributeModel(conn *gorm.DB) SpuAttributeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&SpuAttribute{})
	return &defaultSpuAttributeModel{
		conn: conn,
	}
}
func (m *defaultSpuAttributeModel) AddSpuAttribute(info *SpuAttribute) (*SpuAttribute, error) {
	err := m.conn.Model(&SpuAttribute{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultSpuAttributeModel) UpdateSpuAttribute(id int64, role *SpuAttribute) error {
	err := m.conn.Model(&SpuAttribute{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultSpuAttributeModel) DeleteSpuAttributeBySpuID(id int64) error {
	err := m.conn.Model(&SpuAttribute{}).Where("product_id=?", id).Delete(&SpuAttribute{}).Error
	return err
}
func (m *defaultSpuAttributeModel) GetAttributeValueByProductID(id int64) ([]SpuAttribute, error) {
	var info []SpuAttribute
	err := m.conn.Model(&SpuAttribute{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
