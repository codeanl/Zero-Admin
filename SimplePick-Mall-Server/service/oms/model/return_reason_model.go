package model

import (
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"fmt"
	"gorm.io/gorm"
)

type (
	ReturnReasonModel interface {
		AddReturnReason(ReturnReason *ReturnReason) (*ReturnReason, error)
		UpdateReturnReason(id int64, ReturnReason *ReturnReason) error
		DeleteReturnReasonByIds(ids []int64) error
		GetReturnReasonList(in *oms.ReturnReasonListReq) ([]*ReturnReason, int64, error)
		GetReturnReasonById(id int64) (ReturnReason *ReturnReason, err error)
	}

	defaultReturnReasonModel struct {
		conn *gorm.DB
	}
	ReturnReason struct {
		gorm.Model
		Name   string `json:"name" gorm:"type:varchar(191);comment:退货原因;not null"` //退货原因
		Status string `json:"status" gorm:"type:varchar(191);comment:状态;not null"` //状态
	}
)

func NewReturnReasonModel(conn *gorm.DB) ReturnReasonModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&ReturnReason{})
	return &defaultReturnReasonModel{
		conn: conn,
	}
}
func (m *defaultReturnReasonModel) AddReturnReason(info *ReturnReason) (*ReturnReason, error) {
	err := m.conn.Model(&ReturnReason{}).Create(&info).Error
	return info, err
}

//UpdateReturnReason 修改角色
func (m *defaultReturnReasonModel) UpdateReturnReason(id int64, info *ReturnReason) error {
	err := m.conn.Model(&ReturnReason{}).Where("id=?", id).Updates(info).Error
	return err
}
func (m *defaultReturnReasonModel) GetReturnReasonById(id int64) (info *ReturnReason, err error) {
	err = m.conn.Model(&ReturnReason{}).Where("id=?", id).Find(&info).Error
	return info, err
}

func (m *defaultReturnReasonModel) DeleteReturnReasonByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&ReturnReason{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultReturnReasonModel) GetReturnReasonList(in *oms.ReturnReasonListReq) ([]*ReturnReason, int64, error) {
	var list []*ReturnReason
	db := m.conn.Model(&ReturnReason{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Status != "" {
		db = db.Where("status LIKE ?", in.Status)
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
