package model

import "gorm.io/gorm"

type (
	AttributeValueModel interface {
		GetValueByID(id int64) (info []AttributeValue, err error)
		AddAttributeValue(info *AttributeValue) (err error)
		UpdateAttributeValue(id int64, role *AttributeValue) error
		GetAttributeValueByID(id int64) (info AttributeValue, err error)
		DeleteAttributeValueByProductID(id int64) (err error)
		GetAttributeValueBySPUID(id int64) (info []AttributeValue, err error)
		GetAttributeValueBySpuIdAndAttrId(SpuID, AttrID int64) (info []AttributeValue, err error)
		GetAttributeValueBySpuIdAndValue(SpuID int64, value string) (info AttributeValue, err error)
		DeleteAttributeValueByProductIDAndAttrID(SpuID, AttrID int64) (err error)
	}

	defaultAttributeValueModel struct {
		conn *gorm.DB
	}
	AttributeValue struct {
		gorm.Model
		ProductID   int64  `json:"product_id" gorm:"type:bigint;comment:商品id;not null"`   //商品id
		AttributeID int64  `json:"attribute_id" gorm:"type:bigint;comment:属性id;not null"` //属性id
		Value       string `json:"value" gorm:"type:varchar(191);comment:属性值名称;not null"` //属性值名称
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
func (m *defaultAttributeValueModel) GetAttributeValueBySPUID(id int64) (info []AttributeValue, err error) {
	err = m.conn.Model(&AttributeValue{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultAttributeValueModel) GetAttributeValueBySpuIdAndAttrId(SpuID, AttrID int64) (info []AttributeValue, err error) {
	err = m.conn.Model(&AttributeValue{}).Where("product_id=? && attribute_id=?", SpuID, AttrID).Find(&info).Error
	return info, err
}
func (m *defaultAttributeValueModel) GetAttributeValueBySpuIdAndValue(SpuID int64, value string) (info AttributeValue, err error) {
	err = m.conn.Model(&AttributeValue{}).Where("product_id=? && value=?", SpuID, value).Find(&info).Error
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
func (m *defaultAttributeValueModel) DeleteAttributeValueByProductID(id int64) (err error) {
	var info *AttributeValue
	return m.conn.Model(&AttributeValue{}).Where("product_id=?", id).Delete(&info).Error
}
func (m *defaultAttributeValueModel) DeleteAttributeValueByProductIDAndAttrID(SpuID, AttrID int64) (err error) {
	var info *AttributeValue
	return m.conn.Model(&AttributeValue{}).Where("product_id=? && attribute_id=?", SpuID, AttrID).Delete(&info).Error
}
