package model

import (
	"gorm.io/gorm"
)

type (
	OrderSkuModel interface {
		AddOrderSku(OrderSku *OrderSku) (*OrderSku, error)
		UpdateOrderSku(id int64, role *OrderSku) error
		DeleteOrderSkuBySpuID(id int64) error
		GetOrderSkuByOrderID(id int64) ([]OrderSku, error)
	}

	defaultOrderSkuModel struct {
		conn *gorm.DB
	}
	OrderSku struct {
		gorm.Model
		OrderID int64 `json:"order_id" gorm:"type:bigint;comment:订单id;not null"` //订单id
		SkuID   int64 `json:"sku_id" gorm:"type:bigint;comment:skuId;not null"`  //skuId
	}
)

func NewOrderSkuModel(conn *gorm.DB) OrderSkuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&OrderSku{})
	return &defaultOrderSkuModel{
		conn: conn,
	}
}
func (m *defaultOrderSkuModel) AddOrderSku(info *OrderSku) (*OrderSku, error) {
	err := m.conn.Model(&OrderSku{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultOrderSkuModel) UpdateOrderSku(id int64, role *OrderSku) error {
	err := m.conn.Model(&OrderSku{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultOrderSkuModel) DeleteOrderSkuBySpuID(id int64) error {
	err := m.conn.Model(&OrderSku{}).Where("product_id=?", id).Delete(&OrderSku{}).Error
	return err
}
func (m *defaultOrderSkuModel) GetOrderSkuByOrderID(id int64) ([]OrderSku, error) {
	var info []OrderSku
	err := m.conn.Model(&OrderSku{}).Where("order_id=?", id).Find(&info).Error
	return info, err
}
