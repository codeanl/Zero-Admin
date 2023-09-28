package model

import (
	"SimplePick-Mall-Server/service/pms/rpc/pms"
	"fmt"
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
		CategoryID          int64   `json:"category_id" gorm:"type:bigint;comment:商品分类id;not null"`           //商品分类id
		AttributeCategoryID int64   `json:"attribute_category_id" gorm:"type:bigint;comment:属性分类id;not null"` //属性分类id
		Name                string  `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`              //商品名称
		Pic                 string  `json:"pic" gorm:"type:varchar(191);comment:封面图片;not null"`               //图片
		ProductSn           string  `json:"product_sn" gorm:"type:varchar(191);comment:货号;not null"`          //货号
		Desc                string  `json:"sub_title" gorm:"type:varchar(191);comment:商品描述;not null"`         //副标题
		Price               float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`            //价格
		Sale                int64   `json:"sale" gorm:"type:bigint;comment:销量;not null"`                      //销量
		OriginalPrice       float64 `json:"original_price" gorm:"type:decimal(10, 2);comment:市场价;not null"`   //市场价
		Unit                string  `json:"unit" gorm:"type:varchar(191);comment:单位;not null"`                //单位
		MerchantID          int64   `json:"merchant_id" gorm:"type:bigint;comment:商家id;not null"`             //商家id
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

func (m *defaultProductModel) GetProductList(in *pms.ProductListReq) ([]*Product, int64, error) {
	var list []*Product
	db := m.conn.Model(&Product{}).Order("created_at DESC")
	if in.CategoryID != 0 {

		var count int64
		err := m.conn.Model(&Category{}).Where("parent_id = ?", in.CategoryID).Count(&count).Error
		if err != nil {
			return list, 0, err
		}
		// 检查是否有父级分类
		if count == 0 {
			db = db.Where("category_id = ?", in.CategoryID)
		} else {
			var category []Category
			err := m.conn.Model(&Category{}).Where("parent_id = ?", in.CategoryID).Find(&category).Error
			if err != nil {
				return list, 0, err
			}
			var ids []int64
			for _, i := range category {
				ids = append(ids, int64(i.ID))
			}
			db = db.Where("category_id IN (?)", ids)
		}
	}
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
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
