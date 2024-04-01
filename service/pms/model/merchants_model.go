package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	MerchantsModel interface {
		AddMerchants(info *Merchants) (*Merchants, error)
		UpdateMerchants(id int64, role *Merchants) error
		DeleteMerchantsBySpuID(id int64) error
		GetImgtByProducID(id int64) ([]Merchants, error)
		DeleteMerchantsByIds(ids []int64) error
		GetMerchantsList(in *pms.MerchantsListReq) ([]*Merchants, int64, error)
		GetMerchantsByIDOrUserID(in *pms.MerchantsInfoReq) (info Merchants, err error)
		GetMerchantsByPhone(phone string) (user *Merchants, exist bool, err error)
	}
	defaultMerchantsModel struct {
		conn *gorm.DB
	}
	Merchants struct {
		gorm.Model
		Name      string `json:"name" gorm:"type:varchar(225);comment:名称;not null"`       //名称
		Principal string `json:"principal" gorm:"type:varchar(225);comment:负责人;not null"` //负责人
		Phone     string `json:"phone" gorm:"type:varchar(225);comment:联系电话;not null"`    //联系电话
		Address   string `json:"address" gorm:"type:varchar(225);comment:地址;not null"`    //地址
		Pic       string `json:"pic" gorm:"type:varchar(225);comment:图片;not null"`        //图片
		UserID    int64  `json:"user_id" gorm:"type:bigint;comment:用户id;not null"`        //用户id
	}
)

func NewMerchantsModel(conn *gorm.DB) MerchantsModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Merchants{})
	return &defaultMerchantsModel{
		conn: conn,
	}
}
func (m *defaultMerchantsModel) AddMerchants(info *Merchants) (*Merchants, error) {
	err := m.conn.Model(&Merchants{}).Create(info).Error
	return info, err
}

//UpdateRole 修改角色
func (m *defaultMerchantsModel) UpdateMerchants(id int64, role *Merchants) error {
	err := m.conn.Model(&Merchants{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultMerchantsModel) DeleteMerchantsBySpuID(id int64) error {
	err := m.conn.Model(&Merchants{}).Where("product_id=?", id).Delete(&Merchants{}).Error
	return err
}
func (m *defaultMerchantsModel) GetAttributeValueByProductID(id int64) ([]Merchants, error) {
	var info []Merchants
	err := m.conn.Model(&Merchants{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultMerchantsModel) GetImgtByProducID(id int64) ([]Merchants, error) {
	var info []Merchants
	err := m.conn.Model(&Merchants{}).Where("product_id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultMerchantsModel) DeleteMerchantsByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Merchants{}).Error
	return err
}

func (m *defaultMerchantsModel) GetMerchantsList(in *pms.MerchantsListReq) ([]*Merchants, int64, error) {
	var list []*Merchants
	db := m.conn.Model(&Merchants{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Address != "" {
		db = db.Where("address LIKE ?", fmt.Sprintf("%%%s%%", in.Address))
	}
	if in.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", in.Phone))
	}
	if in.Principal != "" {
		db = db.Where("principal LIKE ?", fmt.Sprintf("%%%s%%", in.Principal))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if in.PageNum > 0 && in.PageSize > 0 {
		err = db.Offset(int((in.PageNum - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
func (m *defaultMerchantsModel) GetMerchantsByIDOrUserID(in *pms.MerchantsInfoReq) (info Merchants, err error) {
	fmt.Println(in)
	if in.UserID != 0 {
		err = m.conn.Model(&Merchants{}).Where("user_id=?", in.UserID).Find(&info).Error
		return info, err
	}
	if in.ID != 0 {
		err = m.conn.Model(&Merchants{}).Where("id=?", in.ID).Find(&info).Error
		return info, err
	}
	return info, nil
}
func (m *defaultMerchantsModel) GetMerchantsByPhone(phone string) (user *Merchants, exist bool, err error) {
	var count int64
	err = m.conn.Model(&Merchants{}).Where("phone=?", phone).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&Merchants{}).Where("phone=?", phone).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}
