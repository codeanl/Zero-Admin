package model

import (
	"SimplePick-Mall-Server/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	UserModel interface {
		GetUserByUsername(username string) (user *User, exist bool, err error)
		AddUser(user *User) (err error)
		UpdateUser(id int64, user *User) error
		DeleteByIds(ids []int64) error
		GetUserList(in *sys.UserListReq) ([]*User, int64, error)
		GetUserByID(id int64) (user *User, err error)
	}
	defaultUserModel struct {
		conn *gorm.DB
	}
	User struct {
		gorm.Model
		Username string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Phone    string `json:"phone" gorm:"type:varchar(191);comment:手机号;not null"`    //手机号
		Nickname string `json:"nickname" gorm:"type:varchar(191);comment:昵称;not null"`  //昵称
		Password string `json:"password" gorm:"type:varchar(191);comment:密码;not null"`  //密码
		Gender   string `json:"gender" gorm:"type:varchar(255);comment:性别;"`            //性别  0--保密 1--男  2--女
		Avatar   string `json:"avatar" gorm:"type:varchar(255);comment:用户头像"`           //头像
		Email    string `json:"email" gorm:"type:varchar(255);comment:邮箱"`              //邮箱
		Status   string `json:"status" gorm:"type:varchar(255);comment:状态;"`            //状态  1--正常 0--禁用 默认正常
		CreateBy string `json:"create_by" gorm:"type:varchar(255);comment:创建人;"`        //创建人
		UpdateBy string `json:"update_by" gorm:"type:varchar(255);comment:更新人;"`        //更新人
	}
)

func NewUserModel(conn *gorm.DB) UserModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&User{})
	return &defaultUserModel{
		conn: conn,
	}
}

// GetUserByID 根据用户名查询用户信息
func (m *defaultUserModel) GetUserByID(id int64) (user *User, err error) {
	err = m.conn.Model(&User{}).Where("id=?", id).First(&user).Error
	return user, err
}

// GetUserByUsername  根据用户名查询用户信息｜查询用户是或存在
func (m *defaultUserModel) GetUserByUsername(username string) (user *User, exist bool, err error) {
	var count int64
	err = m.conn.Model(&User{}).Where("username=?", username).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = m.conn.Model(&User{}).Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, false, err
	}
	return user, true, nil
}

//AddUser 创建用户
func (m *defaultUserModel) AddUser(user *User) (err error) {
	return m.conn.Model(&User{}).Create(user).Error
}

//UpdateUser 更新用户
func (m *defaultUserModel) UpdateUser(id int64, user *User) error {
	err := m.conn.Model(&User{}).Where("id=?", id).Updates(user).Error
	return err
}

//DeleteByIds 删除用户
func (m *defaultUserModel) DeleteByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&User{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultUserModel) GetUserList(in *sys.UserListReq) ([]*User, int64, error) {
	var list []*User
	db := m.conn.Model(&list).Order("created_at DESC")
	if in.Nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Nickname))
	}
	if in.Phone != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", in.Phone))
	}
	if in.Username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Status != "" {
		db = db.Where("status LIKE ?", fmt.Sprintf("%%%s%%", in.Status))
	}
	if in.Gender != "" {
		db = db.Where("gender LIKE ?", fmt.Sprintf("%%%s%%", in.Gender))
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
