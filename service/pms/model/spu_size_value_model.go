package model

import (
	"gorm.io/gorm"
)

type (
	SpuSizeValueModel interface {
		AddSpuSizeValue(SpuSizeValue *SpuSizeValue) (*SpuSizeValue, error)
		UpdateSpuSizeValue(id int64, role *SpuSizeValue) error
		DeleteSpuSizeValueBySizeID(id int64) error
		GetSizeValueBySizeID(id int64) ([]SpuSizeValue, error)
		GetSizeValueByID(id int64) (info SpuSizeValue, err error)
	}

	defaultSpuSizeValueModel struct {
		conn *gorm.DB
	}
	SpuSizeValue struct {
		gorm.Model
		SizeID int64  `json:"size_id" gorm:"type:bigint;comment:名称;not null"`     //规格id
		Value  string `json:"name" gorm:"type:varchar(255);comment:规格名;not null"` //规格值
	}
)

func NewSpuSizeValueModel(conn *gorm.DB) SpuSizeValueModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&SpuSizeValue{})
	return &defaultSpuSizeValueModel{
		conn: conn,
	}
}
func (m *defaultSpuSizeValueModel) AddSpuSizeValue(info *SpuSizeValue) (*SpuSizeValue, error) {
	err := m.conn.Model(&SpuSizeValue{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultSpuSizeValueModel) UpdateSpuSizeValue(id int64, role *SpuSizeValue) error {
	err := m.conn.Model(&SpuSizeValue{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultSpuSizeValueModel) DeleteSpuSizeValueBySizeID(id int64) error {
	err := m.conn.Model(&SpuSizeValue{}).Where("size_id=?", id).Delete(&SpuSizeValue{}).Error
	return err
}
func (m *defaultSpuSizeValueModel) GetSizeValueBySizeID(id int64) ([]SpuSizeValue, error) {
	var spuSize []SpuSizeValue
	err := m.conn.Model(&SpuSizeValue{}).Where("size_id=?", id).Find(&spuSize).Error
	return spuSize, err
}
func (m *defaultSpuSizeValueModel) GetSizeValueByID(id int64) (info SpuSizeValue, err error) {
	err = m.conn.Model(&SpuSizeValue{}).Where("id=?", id).Find(&info).Error
	return info, err
}
