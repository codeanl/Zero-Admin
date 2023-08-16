package model

import "gorm.io/gorm"

type (
	AttributeValueModel interface {
		GetValueByID(id int64) (info []AttributeValue, err error)
		AddAttributeValue(info *AttributeValue) (err error)
		UpdateAttributeValue(id int64, role *AttributeValue) error
		GetAttributeValueByID(id int64) (info AttributeValue, err error)
		DeleteAttributeValueByAttributeID(id int64) (err error)
	}

	defaultAttributeValueModel struct {
		conn *gorm.DB
	}
	AttributeValue struct {
		gorm.Model
		AttributeID int64  `json:"attribute_id" gorm:"type:bigint;comment:属性id;not null"` //属性id
		Name        string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`     //属性值名称
	}
)

func NewAttributeValueModel(conn *gorm.DB) AttributeValueModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&AttributeValue{})
	return &defaultAttributeValueModel{
		conn: conn,
	}
}

func (m *defaultAttributeValueModel) GetValueByID(id int64) (info []AttributeValue, err error) {
	err = m.conn.Model(&AttributeValue{}).Where("attribute_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultAttributeValueModel) GetAttributeValueByID(id int64) (info AttributeValue, err error) {
	err = m.conn.Model(&AttributeValue{}).Where("id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultAttributeValueModel) AddAttributeValue(info *AttributeValue) (err error) {
	return m.conn.Model(&AttributeValue{}).Create(info).Error
}
func (m *defaultAttributeValueModel) UpdateAttributeValue(id int64, role *AttributeValue) error {
	err := m.conn.Model(&AttributeValue{}).Where("id=?", id).Updates(role).Error
	return err
}
func (m *defaultAttributeValueModel) DeleteAttributeValueByAttributeID(id int64) (err error) {
	var info *AttributeValue
	return m.conn.Model(&AttributeValue{}).Where("attribute_id=?", id).Delete(&info).Error
}
