package model

import (
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"gorm.io/gorm"
)

type (
	CartModel interface {
		AddCart(Cart *Cart) (*Cart, error)
		UpdateCart(id int64, Cart *Cart) error
		DeleteCartByIds(ids []int64) error
		GetCartList(in *oms.CartListReq) ([]*Cart, int64, error)
		GetCartById(id int64) (Cart *Cart, err error)
	}

	defaultCartModel struct {
		conn *gorm.DB
	}
	Cart struct {
		gorm.Model
		UserID int64 `json:"user_id" gorm:"type:bigint;comment:用户ID;not null"`
		SkuID  int64 `json:"sku_id" gorm:"type:bigint;comment:skuID;not null"`
		Count  int64 `json:"count" gorm:"type:bigint;comment:数量;not null"`
	}
)

func NewCartModel(conn *gorm.DB) CartModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Cart{})
	return &defaultCartModel{
		conn: conn,
	}
}
func (m *defaultCartModel) AddCart(info *Cart) (*Cart, error) {
	err := m.conn.Model(&Cart{}).Create(&info).Error
	return info, err
}

//UpdateCart 修改角色
func (m *defaultCartModel) UpdateCart(id int64, info *Cart) error {
	err := m.conn.Model(&Cart{}).Where("id=?", id).Updates(info).Error
	return err
}
func (m *defaultCartModel) GetCartById(id int64) (info *Cart, err error) {
	err = m.conn.Model(&Cart{}).Where("id=?", id).Find(&info).Error
	return info, err
}

func (m *defaultCartModel) DeleteCartByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Cart{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultCartModel) GetCartList(in *oms.CartListReq) ([]*Cart, int64, error) {
	var list []*Cart
	db := m.conn.Model(&Cart{}).Order("created_at DESC")
	if in.UserID != 0 {
		db = db.Where("user_id = ?", in.UserID)
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
