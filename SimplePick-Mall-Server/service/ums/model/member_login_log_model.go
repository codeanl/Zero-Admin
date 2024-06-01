package model

import (
	"SimplePick-Mall-Server/service/ums/rpc/ums"
	"gorm.io/gorm"
)

type (
	MemberLoginLogModel interface {
		AddMemberLoginLog(MemberLoginLog *MemberLoginLog) (err error)
		GetMemberLoginLogList(in *ums.MemberLoginLogListReq) ([]*MemberLoginLog, int64, error)
		DeleteMemberLoginLogByIds(ids []int64) error
	}

	defaultMemberLoginLogModel struct {
		conn *gorm.DB
	}
	MemberLoginLog struct {
		gorm.Model
		UserID int64  `json:"user_id" gorm:"type:bigint;comment:用户名;not null"`   //用户id
		IP     string `json:"ip" gorm:"type:varchar(191);comment:IP地址;not null"` //IP地址
	}
)

func NewMemberLoginLogModel(conn *gorm.DB) MemberLoginLogModel {
	conn.AutoMigrate(&MemberLoginLog{})
	return &defaultMemberLoginLogModel{
		conn: conn,
	}
}
func (m *defaultMemberLoginLogModel) AddMemberLoginLog(info *MemberLoginLog) (err error) {
	return m.conn.Model(&MemberLoginLog{}).Create(info).Error
}
func (m *defaultMemberLoginLogModel) GetMemberLoginLogList(in *ums.MemberLoginLogListReq) ([]*MemberLoginLog, int64, error) {
	var list []*MemberLoginLog
	db := m.conn.Model(&MemberLoginLog{}).Order("created_at DESC")
	if in.UserID != 0 {
		db = db.Where("user_id= ?", in.UserID)
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
func (m *defaultMemberLoginLogModel) DeleteMemberLoginLogByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&MemberLoginLog{}).Error
	return err
}
