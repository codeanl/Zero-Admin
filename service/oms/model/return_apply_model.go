package model

import (
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"gorm.io/gorm"
)

type (
	ReturnApplyModel interface {
		AddReturnApply(ReturnApply *ReturnApply) (*ReturnApply, error)
		UpdateReturnApply(id int64, ReturnApply *ReturnApply) error
		DeleteReturnApplyByIds(ids []int64) error
		GetReturnApplyList(in *oms.ReturnApplyListReq) ([]*ReturnApply, int64, error)
		GetReturnApplyById(id int64) (ReturnApply *ReturnApply, err error)
		GetReturnApplyByOrderID(id int64) (info *ReturnApply, err error)
	}
	defaultReturnApplyModel struct {
		conn *gorm.DB
	}
	ReturnApply struct {
		gorm.Model
		UserID         int64   `json:"user_id" gorm:"type:bigint;comment:订单id;not null"`                //用户id
		OrderID        int64   `json:"order_id" gorm:"type:bigint;comment:订单id;not null"`               //订单id
		ReturnReasonID int64   `json:"return_reason_id" gorm:"type:bigint;comment:订单id;not null"`       //订单id
		PlaceId        int64   `json:"place_id" gorm:"type:bigint;comment:自提点ID;not null"`              //自提点ID
		MerchantID     int64   `json:"merchant_id" gorm:"type:bigint;comment:商家ID;not null"`            //商家ID
		Status         string  `json:"status" gorm:"type:varchar(191);comment:状态;not null"`             //'申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝'
		Description    string  `json:"description" gorm:"type:varchar(191);comment:状态;not null"`        //描述
		ProofPics      string  `json:"proof_pics" gorm:"type:varchar(191);comment:状态;not null"`         //凭证图片
		ReturnAmount   float64 `json:"return_amount" gorm:"type: decimal(10, 2);comment:退款金额;not null"` //退款金额      //价格
	}
)

func NewReturnApplyModel(conn *gorm.DB) ReturnApplyModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&ReturnApply{})
	return &defaultReturnApplyModel{
		conn: conn,
	}
}
func (m *defaultReturnApplyModel) AddReturnApply(info *ReturnApply) (*ReturnApply, error) {
	err := m.conn.Model(&ReturnApply{}).Create(&info).Error
	return info, err
}

//UpdateReturnApply 修改角色
func (m *defaultReturnApplyModel) UpdateReturnApply(id int64, info *ReturnApply) error {
	err := m.conn.Model(&ReturnApply{}).Where("id=?", id).Updates(info).Error
	return err
}
func (m *defaultReturnApplyModel) GetReturnApplyById(id int64) (info *ReturnApply, err error) {
	err = m.conn.Model(&ReturnApply{}).Where("id=?", id).Find(info).Error
	return info, err
}

func (m *defaultReturnApplyModel) DeleteReturnApplyByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&ReturnApply{}).Error
	return err
}

func (m *defaultReturnApplyModel) GetReturnApplyByOrderID(id int64) (info *ReturnApply, err error) {
	err = m.conn.Model(&ReturnApply{}).Where("order_id=?", id).Find(&info).Error
	return info, err
}

//GetUserList 获取用户列表
func (m *defaultReturnApplyModel) GetReturnApplyList(in *oms.ReturnApplyListReq) ([]*ReturnApply, int64, error) {
	var list []*ReturnApply
	db := m.conn.Model(&ReturnApply{}).Order("created_at DESC")
	if in.Status != "" {
		db = db.Where("status = ?", in.Status)
	}
	if in.PlaceId != 0 {
		db = db.Where("place_id = ?", in.PlaceId)
	}
	if in.MerchantID != 0 {
		db = db.Where("merchant_id = ?", in.MerchantID)
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
