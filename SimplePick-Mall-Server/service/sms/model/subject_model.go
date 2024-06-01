package model

import (
	"SimplePick-Mall-Server/service/sms/rpc/sms"
	"gorm.io/gorm"
)

type (
	SubjectModel interface {
		AddSubject(Subject *Subject) (*Subject, error)
		UpdateSubject(id int64, Subject *Subject) error
		DeleteSubjectByIds(ids []int64) error
		GetSubjectList(in *sms.SubjectListReq) ([]*Subject, int64, error)
		GetSubjectById(id int64) (Subject *Subject, err error)
	}

	defaultSubjectModel struct {
		conn *gorm.DB
	}
	Subject struct {
		gorm.Model
		Name   string `json:"name" gorm:"type:varchar(225);comment:名称;not null"`   //名称
		Pic    string `json:"pic" gorm:"type:varchar(225);comment:名称;not null"`    //封面
		Status string `json:"status" gorm:"type:varchar(225);comment:名称;not null"` //是否显示 0-不显示 1-显示
		Sort   int64  `json:"sort" gorm:"type:varchar(225);comment:名称;not null"`
	}
)

func NewSubjectModel(conn *gorm.DB) SubjectModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Subject{})
	return &defaultSubjectModel{
		conn: conn,
	}
}
func (m *defaultSubjectModel) AddSubject(info *Subject) (*Subject, error) {
	err := m.conn.Model(&Subject{}).Create(&info).Error
	return info, err
}

//UpdateSubject 修改角色
func (m *defaultSubjectModel) UpdateSubject(id int64, info *Subject) error {
	err := m.conn.Model(&Subject{}).Where("id=?", id).Updates(info).Error
	return err
}
func (m *defaultSubjectModel) GetSubjectById(id int64) (info *Subject, err error) {
	err = m.conn.Model(&Subject{}).Where("id=?", id).Find(&info).Error
	return info, err
}

func (m *defaultSubjectModel) DeleteSubjectByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Subject{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultSubjectModel) GetSubjectList(in *sms.SubjectListReq) ([]*Subject, int64, error) {
	var list []*Subject
	db := m.conn.Model(&Subject{}).Order("sort ASC")
	if in.Status != "" {
		db = db.Where("status = ?", in.Status)
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
