package model

import (
	"SimplePick-Mall-Server/service/ums/rpc/ums"
	"fmt"
	"gorm.io/gorm"
)

type (
	MemberModel interface {
		GetMemberByMembername(Membername string) (Member *Member, exist bool, err error)
		AddMember(Member *Member) (err error)
		UpdateMember(id int64, Member *Member) error
		DeleteByIds(ids []int64) error
		GetMemberList(in *ums.MemberListReq) ([]*Member, int64, error)
		GetMemberByID(id int64) (Member *Member, err error)
	}
	defaultMemberModel struct {
		conn *gorm.DB
	}
	Member struct {
		gorm.Model
		Username  string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Password  string `json:"password" gorm:"type:varchar(191);comment:密码;not null"`  //密码
		Nickname  string `json:"nickname" gorm:"type:varchar(191);comment:昵称;not null"`  //昵称
		Phone     string `json:"phone" gorm:"type:varchar(191);comment:手机号;not null"`    //手机号
		Status    string `json:"status" gorm:"type:varchar(255);comment:状态;"`            //状态  0--正常 1--禁用 默认正常
		Avatar    string `json:"avatar" gorm:"type:varchar(255);comment:用户头像"`           //头像
		Gender    string `json:"gender" gorm:"type:varchar(255);comment:性别;"`            //性别  0--保密 1--男  2--女
		Email     string `json:"email" gorm:"type:varchar(255);comment:邮箱"`              //邮箱
		City      string `json:"city" gorm:"type:varchar(255);comment:所做城市"`             //所做城市
		Job       string `json:"job" gorm:"type:varchar(255);comment:职业"`                //职业
		Signature string `json:"signature" gorm:"type:varchar(255);comment:个性签名"`        //个性签名
	}
)

func NewMemberModel(conn *gorm.DB) MemberModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Member{})
	return &defaultMemberModel{
		conn: conn,
	}
}
func (m *defaultMemberModel) GetMemberByID(id int64) (info *Member, err error) {
	err = m.conn.Model(&Member{}).Where("id=?", id).First(&info).Error
	return info, err
}

//GetMemberByPhone 查询用户是或存在
func (m *defaultMemberModel) GetMemberByMembername(username string) (info *Member, exist bool, err error) {
	var count int64
	err = m.conn.Model(&Member{}).Where("username=?", username).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&Member{}).Where("username=?", username).First(&info).Error
	if err != nil {
		return nil, false, err
	}
	return info, true, nil
}

//AddMember 创建用户
func (m *defaultMemberModel) AddMember(info *Member) (err error) {
	return m.conn.Model(&Member{}).Create(info).Error
}

//UpdateMember 更新用户
func (m *defaultMemberModel) UpdateMember(id int64, info *Member) error {
	err := m.conn.Model(&Member{}).Where("id=?", id).Updates(info).Error
	return err
}

//DeleteByIds 删除用户
func (m *defaultMemberModel) DeleteByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Member{}).Error
	return err
}

//GetMemberList 获取用户列表
func (m *defaultMemberModel) GetMemberList(in *ums.MemberListReq) ([]*Member, int64, error) {
	var list []*Member
	db := m.conn.Model(&list).Order("created_at DESC")
	if in.Username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", in.Phone))
	}
	if in.Status != "" {
		db = db.Where("status= ?", in.Status)
	}
	if in.Gender != "" {
		db = db.Where("gender= ?", in.Gender)
	}
	if in.Nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Nickname))
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
