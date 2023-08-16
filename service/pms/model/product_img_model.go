package model

import (
	"gorm.io/gorm"
)

type (
	ProductImgModel interface {
		AddProductImg(info *ProductImg) (*ProductImg, error)
		UpdateProductImg(id int64, role *ProductImg) error
		DeleteProductImgBySpuID(id int64) error
		GetImgtByProducID(id int64) ([]ProductImg, error)
	}

	defaultProductImgModel struct {
		conn *gorm.DB
	}
	ProductImg struct {
		gorm.Model
		ProductID int64  `json:"product_id" gorm:"type:bigint;comment:商品id;not null"` //商品id
		Url       string `json:"url" gorm:"type:varchar(225);comment:图片地址;not null"`  //图片地址
	}
)

func NewProductImgModel(conn *gorm.DB) ProductImgModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&ProductImg{})
	return &defaultProductImgModel{
		conn: conn,
	}
}
func (m *defaultProductImgModel) AddProductImg(info *ProductImg) (*ProductImg, error) {
	err := m.conn.Model(&ProductImg{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultProductImgModel) UpdateProductImg(id int64, role *ProductImg) error {
	err := m.conn.Model(&ProductImg{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultProductImgModel) DeleteProductImgBySpuID(id int64) error {
	err := m.conn.Model(&ProductImg{}).Where("product_id=?", id).Delete(&ProductImg{}).Error
	return err
}
func (m *defaultProductImgModel) GetAttributeValueByProductID(id int64) ([]ProductImg, error) {
	var info []ProductImg
	err := m.conn.Model(&ProductImg{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultProductImgModel) GetImgtByProducID(id int64) ([]ProductImg, error) {
	var info []ProductImg
	err := m.conn.Model(&ProductImg{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
