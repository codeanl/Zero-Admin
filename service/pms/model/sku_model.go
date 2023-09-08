package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"gorm.io/gorm"
)

type (
	SkuModel interface {
		AddSku(Sku *Sku) (*Sku, error)
		UpdateSku(id int64, role *Sku) error
		DeleteSkuByIds(ids []int64) error
		GetSkuList(in *pms.SkuListReq) ([]*Sku, int64, error)
		GetSkuById(id int64) (info *Sku, err error)
		DeleteSkuBySpuID(id int64) error
		GetSkuByTag(tag string) (info *Sku, err error)
	}

	defaultSkuModel struct {
		conn *gorm.DB
	}
	Sku struct {
		gorm.Model
		ProductID   int64   `json:"product_id" gorm:"type:bigint;comment:商品id;not null"`        //商品id
		Name        string  `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`        //商品名称
		Pic         string  `json:"pic" gorm:"type:varchar(191);comment:图片;not null"`           //图片
		SkuSn       string  `json:"Sku_sn" gorm:"type:varchar(191);comment:货号;not null"`        //货号
		Description string  `json:"description" gorm:"type:varchar(191);comment:商品描述;not null"` //商品描述
		Price       float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`      //价格
		Stock       int64   `json:"stock" gorm:"type:bigint;comment:库存;not null;default:0"`     //库存
		Sale        int64   `json:"sale" gorm:"type:bigint;comment:销量;not null;default:0"`      //销量
		Tag         string  `json:"tag" gorm:"type:varchar(191);comment:tag;not null"`
	}
)

func NewSkuModel(conn *gorm.DB) SkuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Sku{})
	return &defaultSkuModel{
		conn: conn,
	}
}
func (m *defaultSkuModel) AddSku(info *Sku) (*Sku, error) {
	err := m.conn.Model(&Sku{}).Create(info).Error
	return info, err
}

func (m *defaultSkuModel) UpdateSku(id int64, role *Sku) error {
	err := m.conn.Model(&Sku{}).Where("id=?", id).Updates(role).Error
	return err
}
func (m *defaultSkuModel) GetSkuById(id int64) (info *Sku, err error) {
	err = m.conn.Model(&Sku{}).Where("id=?", id).Find(&info).Error
	return info, err
}
func (m *defaultSkuModel) GetSkuByTag(tag string) (info *Sku, err error) {
	err = m.conn.Model(&Sku{}).Where("tag=?", tag).Find(&info).Error
	return info, err
}

//DeleteRoleByIds 删除角色
func (m *defaultSkuModel) DeleteSkuByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Sku{}).Error
	return err
}

func (m *defaultSkuModel) DeleteSkuBySpuID(id int64) error {
	err := m.conn.Model(&Sku{}).Where("product_id=?", id).Delete(&Sku{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultSkuModel) GetSkuList(in *pms.SkuListReq) ([]*Sku, int64, error) {
	var list []*Sku
	db := m.conn.Model(&Sku{}).Order("created_at DESC")
	if in.ProductID != 0 {
		db = db.Where("product_id = ?", in.ProductID)
	}
	err := db.Find(&list).Error
	if err != nil {
		return list, 0, err
	}
	total := int64(len(list))
	return list, total, err
}
