package model

import (
	"SimplePick-Mall-Server/service/oms/rpc/oms"
	"fmt"
	"gorm.io/gorm"
)

type (
	OrderModel interface {
		AddOrder(order *Order) (err error)
		UpdateOrder(id int64, order *Order) error
		DeleteOrderByIds(ids []int64) error
		GetOrderList(in *oms.OrderListReq) ([]*Order, int64, error)
		GetOrderById(id int64) (order *Order, err error)
	}

	defaultOrderModel struct {
		conn *gorm.DB
	}
	Order struct {
		gorm.Model
		MemberId              int64   `json:"member_id" gorm:"type:bigint;comment:用户ID;not null"`                 //用户ID
		CouponId              int64   `json:"coupon_id" gorm:"type:bigint;comment:优惠券ID;not null"`                //优惠券ID
		PlaceId               int64   `json:"place_id" gorm:"type:bigint;comment:自提点ID;not null"`                 //自提点ID
		OrderSn               string  `json:"order_sn" gorm:"type:varchar(191);comment:订单编号;not null"`            //订单编号
		MemberUsername        string  `json:"member_username" gorm:"type:varchar(191);comment:备注;not null"`       //用户帐号
		TotalAmount           float64 `json:"total_amount" gorm:"type:decimal(10, 2) ;comment:订单总金额;not null"`    //订单总金额
		PayAmount             float64 `json:"pay_amount" gorm:"type:decimal(10, 2) ;comment:应付金额;not null"`       //应付金额
		FreightAmount         float64 `json:"freight_amount" gorm:"type:decimal(10, 2) ;comment:运费金额;not null"`   //运费金额
		CouponAmount          float64 `json:"coupon_amount" gorm:"type:decimal(10, 2) ;comment:优惠券抵扣金额;not null"` //优惠券抵扣金额
		PayType               string  `json:"pay_type" gorm:"type:varchar(191);comment:支付方式;not null"`            //支付方式：1->支付宝；2->微信',
		Status                string  `json:"status" gorm:"type:varchar(191);comment:订单状态;not null"`              //订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单',
		OrderType             string  `json:"order_type" gorm:"type:varchar(191);comment:订单类型;not null"`          //订单类型：1->正常订单；2->秒杀订单',
		ReceiverName          string  `json:"receiver_name" gorm:"type:varchar(191);comment:收货人姓名;not null"`      //收货人姓名
		ReceiverPhone         string  `json:"receiver_phone" gorm:"type:varchar(191);comment:收货人电话;not null"`
		ReceiverProvince      string  `json:"receiver_province" gorm:"type:varchar(191);comment:省份;not null"`         //省份
		ReceiverCity          string  `json:"receiver_city" gorm:"type:varchar(191);comment:城市;not null"`             //城市
		ReceiverRegion        string  `json:"receiver_region" gorm:"type:varchar(191);comment:区;not null"`            //区
		ReceiverDetailAddress string  `json:"receiver_detail_address" gorm:"type:varchar(191);comment:详细地址;not null"` //详细地址
		Note                  string  `json:"note" gorm:"type:varchar(191);comment:订单备注;not null"`                    //订单备注
		ConfirmStatus         string  `json:"confirm_status" gorm:"type:varchar(191);comment:确认收货状态;not null"`        //确认收货状态：0->未确认；1->已确认,
		DeleteStatus          string  `json:"delete_status" gorm:"type:varchar(191);comment:删除状态;not null"`           //删除状态：0->未删除；1->已删除
		PaymentTime           string  `json:"payment_time" gorm:"type:datetime;comment:支付时间;not null"`                //支付时间
		DeliveryTime          string  `json:"delivery_time" gorm:"type:datetime;comment:发货时间;not null"`               //发货时间
		ReceiveTime           string  `json:"receive_time" gorm:"type:datetime;comment:确认收货时间;not null"`              //确认收货时间
		CommentTime           string  `json:"comment_time" gorm:"type:datetime;comment:评价时间;not null"`                //评价时间
	}
)

func NewOrderModel(conn *gorm.DB) OrderModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Order{})
	return &defaultOrderModel{
		conn: conn,
	}
}
func (m *defaultOrderModel) AddOrder(order *Order) (err error) {
	return m.conn.Model(&Order{}).Create(order).Error
}

//UpdateOrder 修改角色
func (m *defaultOrderModel) UpdateOrder(id int64, order *Order) error {
	err := m.conn.Model(&Order{}).Where("id=?", id).Updates(order).Error
	return err
}
func (m *defaultOrderModel) GetOrderById(id int64) (order *Order, err error) {
	err = m.conn.Model(&Order{}).Where("id=?", id).Find(order).Error
	return order, err
}

func (m *defaultOrderModel) DeleteOrderByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Order{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultOrderModel) GetOrderList(in *oms.OrderListReq) ([]*Order, int64, error) {
	var list []*Order
	db := m.conn.Model(&Order{}).Order("created_at DESC")
	if in.OrderSn != "" {
		db = db.Where("order_sn LIKE ?", fmt.Sprintf("%%%s%%", in.OrderSn))
	}
	if in.MemberUsername != "" {
		db = db.Where("order_sn LIKE ?", fmt.Sprintf("%%%s%%", in.MemberUsername))
	}
	if in.Status != "" {
		db = db.Where("order_sn LIKE ?", fmt.Sprintf("%%%s%%", in.Status))
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
