package model

import (
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	PlaceModel interface {
		AddPlace(role *Place) (err error)
		DeletePlaceByIds(ids []int64) error
		GetPlaceList(in *sys.PlaceListReq) ([]*Place, int64, error)
		UpdatePlace(id int64, role *Place) error
		GetPlaceById(id int64) (place *Place, err error)
	}

	defaultPlaceModel struct {
		conn *gorm.DB
	}
	Place struct {
		gorm.Model
		Name      string `json:"name" gorm:"type:varchar(191);comment:自提点名称;not null"`    //自提点名称
		Place     string `json:"place" gorm:"type:varchar(191);comment:详细地址;not null"`    //详细地址
		Status    string `json:"status" gorm:"type:varchar(191);comment:开业状态;not null"`   //开业状态0--开业 1--关门
		Pic       string `json:"pic" gorm:"type:varchar(191);comment:门面图片;not null"`      //门面图片
		Phone     string `json:"phone" gorm:"type:varchar(191);comment:联系电话;not null"`    //联系电话
		Principal string `json:"principal" gorm:"type:varchar(191);comment:负责人;not null"` //负责人
		CreateBy  string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
		UpdateBy  string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"` //更新人
	}
)

func NewPlaceModel(conn *gorm.DB) PlaceModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Place{})
	return &defaultPlaceModel{
		conn: conn,
	}
}

func (m *defaultPlaceModel) AddPlace(role *Place) (err error) {
	return m.conn.Model(&Place{}).Create(role).Error
}

//
func (m *defaultPlaceModel) UpdatePlace(id int64, role *Place) error {
	err := m.conn.Model(&Place{}).Where("id=?", id).Updates(role).Error
	return err
}

func (m *defaultPlaceModel) DeletePlaceByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Place{}).Error
	return err
}
func (m *defaultPlaceModel) GetPlaceById(id int64) (place *Place, err error) {
	place = &Place{}
	err = m.conn.Model(&Place{}).Where("id=?", id).First(place).Error
	return place, err
}

func (m *defaultPlaceModel) GetPlaceList(in *sys.PlaceListReq) ([]*Place, int64, error) {
	var list []*Place
	db := m.conn.Model(&Place{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Place != "" {
		db = db.Where("place LIKE ?", fmt.Sprintf("%%%s%%", in.Place))
	}
	if in.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", in.Phone))
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
