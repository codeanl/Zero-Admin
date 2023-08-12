package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"gorm.io/gorm"
)

type (
	ProductModel interface {
		AddProduct(product *Product) (*Product, error)
		UpdateProduct(id int64, role *Product) error
		DeleteProductByIds(ids []int64) error
		GetProductList(in *pms.ProductListReq) ([]*Product, int64, error)
		GetProductById(id int64) (info *Product, err error)
	}

	defaultProductModel struct {
		conn *gorm.DB
	}
	Product struct {
		gorm.Model
		CategoryID    int64   `json:"category_id" gorm:"type:bigint;comment:商品分类id;not null"`         //商品分类id
		Name          string  `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`            //商品名称
		Pic           string  `json:"pic" gorm:"type:varchar(191);comment:图片;not null"`               //图片
		ProductSn     string  `json:"product_sn" gorm:"type:varchar(191);comment:货号;not null"`        //货号
		SubTitle      string  `json:"sub_title" gorm:"type:varchar(191);comment:副标题;not null"`        //副标题
		Description   string  `json:"description" gorm:"type:varchar(191);comment:商品描述;not null"`     //商品描述
		Price         float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`          //价格
		OriginalPrice float64 `json:"original_price" gorm:"type:decimal(10, 2);comment:市场价;not null"` //市场价
		Stock         int64   `json:"stock" gorm:"type:bigint;comment:库存;not null;default:0"`         //库存
		Unit          string  `json:"unit" gorm:"type:varchar(191);comment:单位;not null"`              //单位
		Sale          int64   `json:"sale" gorm:"type:bigint;comment:销量;not null"`                    //销量
	}
)

func NewProductModel(conn *gorm.DB) ProductModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Product{})
	return &defaultProductModel{
		conn: conn,
	}
}
func (m *defaultProductModel) AddProduct(product *Product) (*Product, error) {
	err := m.conn.Model(&Product{}).Create(product).Error
	return product, err
}

func (m *defaultProductModel) UpdateProduct(id int64, role *Product) error {
	err := m.conn.Model(&Product{}).Where("id=?", id).Updates(role).Error
	return err
}
func (m *defaultProductModel) GetProductById(id int64) (info *Product, err error) {
	err = m.conn.Model(&Product{}).Where("id=?", id).Find(&info).Error
	return info, err
}

//DeleteRoleByIds 删除角色
func (m *defaultProductModel) DeleteProductByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Product{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultProductModel) GetProductList(in *pms.ProductListReq) ([]*Product, int64, error) {
	var list []*Product
	db := m.conn.Model(&Product{}).Order("created_at DESC")

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
