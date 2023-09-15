package model

import (
	"gorm.io/gorm"
)

type (
	ProductIntroduceImgModel interface {
		AddProductIntroduceImg(info *ProductIntroduceImg) (*ProductIntroduceImg, error)
		UpdateProductIntroduceImg(id int64, role *ProductIntroduceImg) error
		DeleteProductIntroduceImgBySpuID(id int64) error
		GetImgtByProducID(id int64) ([]ProductIntroduceImg, error)
	}

	defaultProductIntroduceImgModel struct {
		conn *gorm.DB
	}
	ProductIntroduceImg struct {
		gorm.Model
		ProductID int64  `json:"product_id" gorm:"type:bigint;comment:商品id;not null"` //商品id
		Url       string `json:"url" gorm:"type:varchar(225);comment:图片地址;not null"`  //图片地址
	}
)

func NewProductIntroduceImgModel(conn *gorm.DB) ProductIntroduceImgModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&ProductIntroduceImg{})
	return &defaultProductIntroduceImgModel{
		conn: conn,
	}
}
func (m *defaultProductIntroduceImgModel) AddProductIntroduceImg(info *ProductIntroduceImg) (*ProductIntroduceImg, error) {
	err := m.conn.Model(&ProductIntroduceImg{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultProductIntroduceImgModel) UpdateProductIntroduceImg(id int64, role *ProductIntroduceImg) error {
	err := m.conn.Model(&ProductIntroduceImg{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultProductIntroduceImgModel) DeleteProductIntroduceImgBySpuID(id int64) error {
	err := m.conn.Model(&ProductIntroduceImg{}).Where("product_id=?", id).Delete(&ProductIntroduceImg{}).Error
	return err
}
func (m *defaultProductIntroduceImgModel) GetAttributeValueByProductID(id int64) ([]ProductIntroduceImg, error) {
	var info []ProductIntroduceImg
	err := m.conn.Model(&ProductIntroduceImg{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultProductIntroduceImgModel) GetImgtByProducID(id int64) ([]ProductIntroduceImg, error) {
	var info []ProductIntroduceImg
	err := m.conn.Model(&ProductIntroduceImg{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
