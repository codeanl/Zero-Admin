package model

import (
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"gorm.io/gorm"
)

type (
	HotRecommendModel interface {
		AddHotRecommend(info *HotRecommend) (err error)
		UpdateHotRecommend(id int64, info *HotRecommend) error
		DeleteHotRecommendByIds(ids []int64) error
		GetHotRecommendList(in *sms.HotRecommendListReq) ([]*HotRecommend, int64, error)
		////GetUserByID(id int64) (user *User, err error)
		ExistHotRecommendById(id int64) (info *HotRecommend, exist bool, err error)
	}
	defaultHotRecommendModel struct {
		conn *gorm.DB
	}
	HotRecommend struct {
		gorm.Model
		ProductId int64  `json:"product_id" gorm:"type:bigint;comment:数量;not null"`
		Status    string `json:"status" gorm:"type:varchar(191);comment:推荐状态;not null"` //推荐状态：0->不展示;1->展示',
		Sort      int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`           //排序
	}
)

func NewHotRecommendModel(conn *gorm.DB) HotRecommendModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&HotRecommend{})
	return &defaultHotRecommendModel{
		conn: conn,
	}
}

func (m *defaultHotRecommendModel) AddHotRecommend(info *HotRecommend) (err error) {
	return m.conn.Model(&HotRecommend{}).Create(info).Error
}
func (m *defaultHotRecommendModel) ExistHotRecommendById(id int64) (info *HotRecommend, exist bool, err error) {
	var count int64
	err = m.conn.Model(&HotRecommend{}).Where("product_id=?", id).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&HotRecommend{}).Where("product_id=?", id).First(&info).Error
	if err != nil {
		return nil, false, err
	}
	return info, true, nil
}

//
func (m *defaultHotRecommendModel) UpdateHotRecommend(id int64, info *HotRecommend) error {
	err := m.conn.Model(&HotRecommend{}).Where("id=?", id).Updates(info).Error
	return err
}

func (m *defaultHotRecommendModel) DeleteHotRecommendByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&HotRecommend{}).Error
	return err
}

//
////GetUserList 获取用户列表
func (m *defaultHotRecommendModel) GetHotRecommendList(in *sms.HotRecommendListReq) ([]*HotRecommend, int64, error) {
	var list []*HotRecommend
	db := m.conn.Model(&list).Order("sort ASC")
	if in.Status != "" {
		db = db.Where("status =?", in.Status)
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
