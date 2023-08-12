package model

import (
	"gorm.io/gorm"
)

type (
	SpuSizeModel interface {
		AddSpuSize(SpuSize *SpuSize) (*SpuSize, error)
		UpdateSpuSize(id int64, role *SpuSize) error
		DeleteSpuSizeBySpuID(id int64) error
		GetSizeBySpuID(id int64) ([]SpuSize, error)
		GetSizeByID(id int64) (info SpuSize, err error)
	}

	defaultSpuSizeModel struct {
		conn *gorm.DB
	}
	SpuSize struct {
		gorm.Model
		ProductID int64  `json:"product_id" gorm:"type:bigint;comment:名称;not null"`     //商品id
		Name      string `json:"name" gorm:"type:varchar(255);comment:属性分类id;not null"` //规格名
	}
)

func NewSpuSizeModel(conn *gorm.DB) SpuSizeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&SpuSize{})
	return &defaultSpuSizeModel{
		conn: conn,
	}
}
func (m *defaultSpuSizeModel) AddSpuSize(info *SpuSize) (*SpuSize, error) {
	err := m.conn.Model(&SpuSize{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultSpuSizeModel) UpdateSpuSize(id int64, role *SpuSize) error {
	err := m.conn.Model(&SpuSize{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultSpuSizeModel) DeleteSpuSizeBySpuID(id int64) error {
	err := m.conn.Model(&SpuSize{}).Where("product_id=?", id).Delete(&SpuSize{}).Error
	return err
}
func (m *defaultSpuSizeModel) GetSizeBySpuID(id int64) ([]SpuSize, error) {
	var size []SpuSize
	err := m.conn.Model(&SpuSize{}).Where("product_id=?", id).Find(&size).Error
	return size, err
}
func (m *defaultSpuSizeModel) GetSizeByID(id int64) (info SpuSize, err error) {
	err = m.conn.Model(&SpuSize{}).Where("id=?", id).Find(&info).Error
	return info, err
}
