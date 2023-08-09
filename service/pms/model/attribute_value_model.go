package model

import "gorm.io/gorm"

type (
	AttributeValueModel interface {
	}

	defaultAttributeValueModel struct {
		conn *gorm.DB
	}
	AttributeValue struct {
		gorm.Model
		Name        string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`     //属性值名称
		AttributeID int64  `json:"attribute_id" gorm:"type:bigint;comment:属性id;not null"` //属性id
	}
)

func NewAttributeValueModel(conn *gorm.DB) AttributeValueModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&AttributeValue{})
	return &defaultAttributeValueModel{
		conn: conn,
	}
}
