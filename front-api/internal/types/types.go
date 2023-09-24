// Code generated by goctl. DO NOT EDIT.
package types

type MemberLoginReq struct {
}

type Info struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Nickname  string `json:"nickname"`  // 昵称
	Phone     string `json:"phone"`     // 手机号码
	Status    string `json:"status"`    // 帐号启用状态:0->禁用；1->启用
	Avatar    string `json:"avatar"`    // 头像
	Gender    string `json:"gender"`    // 性别：0->未知；1->男；2->女
	Email     string `json:"email"`     //
	City      string `json:"city"`      // 所做城市
	Job       string `json:"job"`       // 职业
	Signature string `json:"signature"` // 个性签名
	CreatTIme string `json:"creatTIme"` //
}

type LoginData struct {
	Token string `json:"token"`
	Info  Info   `json:"info"`
}

type MemberLoginResp struct {
	Code    int64     `json:"code"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}

type InfoReq struct {
	ID int64 `json:"id"`
}

type InfoResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    Info   `json:"data"`
}

type UpdateMemberReq struct {
	Id        int64  `json:"id"`
	Username  string `json:"username,optional"`  // 用户名
	Nickname  string `json:"nickname,optional"`  // 昵称
	Phone     string `json:"phone,optional"`     // 手机号码
	Status    string `json:"status,optional"`    // 帐号启用状态:0->禁用；1->启用
	Avatar    string `json:"avatar,optional"`    // 头像
	Gender    string `json:"gender,optional"`    // 性别：0->未知；1->男；2->女
	Email     string `json:"email,optional"`     //
	City      string `json:"city,optional"`      // 所做城市
	Job       string `json:"job,optional"`       // 职业
	Signature string `json:"signature,optional"` // 个性签名
}

type UpdateMemberResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListHomeAdvertiseReq struct {
}

type ListHomeAdvertiseData struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`       // 名称
	Pic        string `json:"pic"`        // 图片地址
	Status     string `json:"status"`     // 上下线状态：0->下线；1->上线
	ClickCount int64  `json:"clickCount"` // 点击数
	Url        string `json:"url"`        // 链接地址
	Note       string `json:"note"`       // 备注
}

type ListHomeAdvertiseResp struct {
	Code    int64                   `json:"code"`
	Message string                  `json:"message"`
	Data    []ListHomeAdvertiseData `json:"data"`
	Total   int64                   `json:"total"`
}

type ListCategoryReq struct {
}

type ListCategoryData struct {
	Id           int64              `json:"id"`
	ParentId     int64              `json:"parentId"` // 上级分类的编号：0表示一级分类
	Name         string             `json:"name"`
	Level        string             `json:"level"` // 分类级别：0->1级；1->2级
	ProductCount int64              `json:"productCount"`
	ProductUnit  string             `json:"productUnit"`
	NavStatus    string             `json:"navStatus"`  // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   string             `json:"showStatus"` // 显示状态：0->不显示；1->显示
	Sort         int64              `json:"sort"`
	Icon         string             `json:"icon"` // 图标
	Keywords     string             `json:"keywords"`
	Description  string             `json:"description"` // 描述
	ProductList  []ListProductData  `json:"productList"`
	Children     []ListCategoryData `json:"children"`
}

type ListCategoryResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    []ListCategoryData `json:"data"`
	Total   int64              `json:"total"`
}

type ListProductReq struct {
	Current    int64  `form:"current,optional"`
	PageSize   int64  `form:"pageSize,optional"`
	Name       string `form:"name,optional"`
	CategoryId int64  `form:"categoryId,optional"`
}

type ListProductByCateIDReq struct {
	CategoryId int64 `json:"categoryId,optional"`
}

type ListProductData struct {
	Id            int64   `json:"id"`
	CategoryID    int64   `json:"categoryId"`
	Name          string  `json:"name"`
	Pic           string  `json:"pic,optional"`
	ProductSn     string  `json:"productSn"`
	Desc          string  `json:"desc,optional"`
	OriginalPrice float64 `json:"originalPrice,optional"`
	Unit          string  `json:"unit,optional"`
	Price         float64 `json:"price,optional"`
}

type ListProductResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    []ListProductData `json:"data"`
	Total   int64             `json:"total"`
}

type ListSubjectReq struct {
	Current  int64  `form:"current,optional"`
	PageSize int64  `form:"pageSize,optional"`
	Name     string `form:"name,optional "`
	Status   string `form:"status,optional"`
}

type ListSubjectData struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Status string `json:"status"`
	Sort   int64  `json:"sort"`
}

type ListSubjectResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Total   int64             `json:"total"`
	Data    []ListSubjectData `json:"data"`
}

type ListSubjectProductReq struct {
	Current   int64  `form:"current,default=1"`
	PageSize  int64  `form:"pageSize,default=20"`
	SubjectID int64  `form:"subjectId"`
	Status    string `form:"status,optional"`
}

type ListSubjectProductData struct {
	Id          int64           `json:"id"`
	SubjectID   int64           `json:"subject_id"`
	ProductID   int64           `json:"product_id"`
	Status      string          `json:"status"`
	Sort        int64           `json:"sort"`
	ProductInfo ProductInfoData `json:"productInfo"`
}

type ProductInfoData struct {
	Id            int64   `json:"id"`
	CategoryID    int64   `json:"categoryId"`
	Name          string  `json:"name"`
	Pic           string  `json:"pic,optional"`
	ProductSn     string  `json:"productSn"`
	Desc          string  `json:"desc,optional"`
	OriginalPrice float64 `json:"originalPrice,optional"`
	Unit          string  `json:"unit,optional"`
	Price         float64 `json:"price,optional"`
}

type ListSubjectProductResp struct {
	Code    int64                    `json:"code"`
	Message string                   `json:"message"`
	Data    []ListSubjectProductData `json:"data"`
	Total   int64                    `json:"total"`
}

type ProductInfoReq struct {
	ID int64 `json:"id"`
}

type InfoData struct {
	ProductInfo     ListProductData  `json:"productInfo"`
	SkuList         []SkuList        `json:"skuList"`
	SizeList        []SizeList       `json:"sizeList"`
	AttributeList   []AttributeLists `json:"attributeList"`
	ImgUrl          []string         `json:"imgUrl"`
	IntroduceImgUrl []string         `json:"introduceImgUrl"`
}

type ProductInfoResp struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Data    InfoData `json:"data"`
}

type SkuList struct {
	ID          int64   `json:"id"`
	ProductID   int64   `json:"productId"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	SkuSn       string  `json:"skuSn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Tag         string  `json:"tag"`
	Specs       []Specs `json:"specs"`
}

type Specs struct {
	Name      string `json:"name"`
	ValueName string `json:"valueName"`
}

type SizeList struct {
	ID     int64    `json:"id,optional"`
	Name   string   `json:"name"`
	Values []Values `json:"values"`
}

type AttributeLists struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Values struct {
	Name string `json:"name"`
}

type AddCartReq struct {
	UserID int64 `json:"userID"`
	SkuID  int64 `json:"skuID"`
	Count  int64 `json:"count"`
}

type AddCartResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UpdateCartReq struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"userID"`
	SkuID  int64 `json:"skuID"`
	Count  int64 `json:"count"`
}

type UpdateCartResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteCartReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteCartResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListCartReq struct {
	Current  int64 `json:"current,optional"`
	PageSize int64 `json:"pageSize,optional"`
	UserID   int64 `json:"userID"`
}

type ListCartData struct {
	ID          int64       `json:"id"`
	UserID      int64       `json:"userID"`
	SkuID       int64       `json:"skuID"`
	Count       int64       `json:"count"`
	ListSkuData ListSkuData `json:"listSkuData"`
}

type ListCartResp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Total   int64          `json:"total"`
	Data    []ListCartData `json:"data"`
}

type ListSkuData struct {
	ID          int64       `json:"id"`
	ProductID   int64       `json:"productId"`
	Name        string      `json:"name"`
	Pic         string      `json:"pic"`
	SkuSn       string      `json:"skuSn"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Stock       int64       `json:"stock"`
	Tag         string      `json:"tag"`
	Space       []SpaceInfo `json:"space"`
	TagText     string      `json:"tagText"`
}

type SpaceInfo struct {
	Name      string `json:"name"`
	ValueName string `json:"valueName"`
}

type AddOrderReq struct {
	MemberId              int64   `json:"memberId,optional"`
	PlaceId               int64   `json:"placeId,optional"`
	CouponId              int64   `json:"couponId,optional"`
	OrderSn               string  `json:"orderSn,optional"`               // 订单编号
	MemberUsername        string  `json:"memberUserName,optional"`        // 用户帐号
	TotalAmount           float64 `json:"totalAmount,optional"`           // 订单总金额
	PayAmount             float64 `json:"payAmount,optional"`             // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freightAmount,optional"`         // 运费金额
	CouponAmount          float64 `json:"couponAmount,optional"`          // 优惠券抵扣金额
	PayType               string  `json:"payType,optional"`               // 支付方式：0->未支付；1->支付宝；2->微信
	Status                string  `json:"status,optional"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             string  `json:"orderType,optional"`             // 订单类型：0->正常订单；1->秒杀订单
	ReceiverProvince      string  `json:"receiverProvince,optional"`      // 省份/直辖市
	ReceiverName          string  `json:"receiverName,optional"`          //
	ReceiverPhone         string  `json:"receiverPhone,optional"`         //
	ReceiverCity          string  `json:"receiverCity,optional"`          // 城市
	ReceiverRegion        string  `json:"receiverRegion,optional"`        // 区
	ReceiverDetailAddress string  `json:"receiverDetailAddress,optional"` // 详细地址
	Note                  string  `json:"note,optional"`                  // 订单备注
	PaymentTime           string  `json:"paymentTime,optional,optional"`  // 支付时间
	DeliveryTime          string  `json:"deliveryTime,optional,optional"` // 发货时间
	ReceiveTime           string  `json:"receiveTime,optional,optional"`  // 确认收货时间
	CommentTime           string  `json:"commentTime,optional,optional"`  // 修改时间
	Skus                  []Skus  `json:"skus"`
}

type Skus struct {
	SkuID int64 `json:"skuID"`
	Count int64 `json:"count"`
}

type AddOrderResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	OrderID int64 `json:"orderID"`
}

type ListOrderReq struct {
	Current  int64  `json:"current,optional"`
	PageSize int64  `json:"pageSize,optional"`
	Status   string `json:"status,optional"`
	UserID   int64  `json:"userID"`
}

type ListOrderData struct {
	ID                    int64         `json:"id"`
	MemberId              int64         `json:"memberId"`
	PlaceId               int64         `json:"placeId"`
	CouponId              int64         `json:"couponId"`
	OrderSn               string        `json:"orderSn"`               // 订单编号
	MemberUsername        string        `json:"memberUserName"`        // 用户帐号
	TotalAmount           float64       `json:"totalAmount"`           // 订单总金额
	PayAmount             float64       `json:"payAmount"`             // 应付金额（实际支付金额）
	FreightAmount         float64       `json:"freightAmount"`         // 运费金额
	CouponAmount          float64       `json:"couponAmount"`          // 优惠券抵扣金额
	PayType               string        `json:"payType"`               // 支付方式：0->未支付；1->支付宝；2->微信
	Status                string        `json:"status"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             string        `json:"orderType"`             // 订单类型：0->正常订单；1->秒杀订单
	ReceiverProvince      string        `json:"receiverProvince"`      // 省份/直辖市
	ReceiverName          string        `json:"receiverName"`          //
	ReceiverPhone         string        `json:"receiverPhone"`         //
	ReceiverCity          string        `json:"receiverCity"`          // 城市
	ReceiverRegion        string        `json:"receiverRegion"`        // 区
	ReceiverDetailAddress string        `json:"receiverDetailAddress"` // 详细地址
	Note                  string        `json:"note"`                  // 订单备注
	ConfirmStatus         string        `json:"confirmStatus"`         // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          string        `json:"deleteStatus"`          // 删除状态：0->未删除；1->已删除
	PaymentTime           string        `json:"paymentTime,optional"`  // 支付时间
	CreateTime            string        `json:"createTime,optional"`   // 支付时间
	DeliveryTime          string        `json:"deliveryTime,optional"` // 发货时间
	ReceiveTime           string        `json:"receiveTime,optional"`  // 确认收货时间
	CommentTime           string        `json:"commentTime,optional"`  // 修改时间
	SkuList               []Sku         `json:"skuList"`
	Place                 PlaceInfoData `json:"placeInfo"`
}

type Sku struct {
	ID          int64   `json:"id"`
	ProductID   int64   `json:"productId"`
	ProductName string  `json:"productName"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	SkuSn       string  `json:"skuSn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Tag         string  `json:"tag"`
	Count       int64   `json:"count"`
}

type PlaceInfo struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Place     string `json:"place"`
	Status    string `json:"status"`
	Pic       string `json:"pic"`
	Phone     string `json:"phone"`
	Principal string `json:"principal"`
}

type ListOrderResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    []ListOrderData `json:"data"`
	Total   int64           `json:"total"`
}

type UpdateOrderReq struct {
	ID                    int64   `json:"id"`
	MemberId              int64   `json:"memberId,optional"`
	PlaceId               int64   `json:"placeId,optional"`
	CouponId              int64   `json:"couponId,optional"`
	OrderSn               string  `json:"orderSn,optional"`               // 订单编号
	MemberUsername        string  `json:"memberUserName,optional"`        // 用户帐号
	TotalAmount           float64 `json:"totalAmount,optional"`           // 订单总金额
	PayAmount             float64 `json:"payAmount,optional"`             // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freightAmount,optional"`         // 运费金额
	CouponAmount          float64 `json:"couponAmount,optional"`          // 优惠券抵扣金额
	PayType               string  `json:"payType,optional"`               // 支付方式：0->未支付；1->支付宝；2->微信
	Status                string  `json:"status,optional"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             string  `json:"orderType,optional"`             // 订单类型：0->正常订单；1->秒杀订单
	ReceiverProvince      string  `json:"receiverProvince,optional"`      // 省份/直辖市
	ReceiverName          string  `json:"receiverName,optional"`          //
	ReceiverPhone         string  `json:"receiverPhone,optional"`         //
	ReceiverCity          string  `json:"receiverCity,optional"`          // 城市
	ReceiverRegion        string  `json:"receiverRegion,optional"`        // 区
	ReceiverDetailAddress string  `json:"receiverDetailAddress,optional"` // 详细地址
	Note                  string  `json:"note,optional"`                  // 订单备注
	ConfirmStatus         string  `json:"confirmStatus,optional"`         // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          string  `json:"deleteStatus,optional"`          // 删除状态：0->未删除；1->已删除
	PaymentTime           string  `json:"paymentTime,optional"`           // 支付时间
	DeliveryTime          string  `json:"deliveryTime,optional"`          // 发货时间
	ReceiveTime           string  `json:"receiveTime,optional"`           // 确认收货时间
	CommentTime           string  `json:"commentTime,optional"`           // 评价时间
}

type UpdateOrderResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteOrderReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteOrderResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type OrderInfoReq struct {
	ID int64 `json:"id"`
}

type SkuListData struct {
	ID          int64   `json:"id"`
	ProductID   int64   `json:"productId"`
	ProductName string  `json:"productName"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	SkuSn       string  `json:"skuSn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Tag         string  `json:"tag"`
	Count       int64   `json:"count"`
}

type OrderInfo struct {
	OrderInfo ListOrderData  `json:"orderInfo"`
	SkuList   []*SkuListData `json:"skuList"`
	PlaceInfo PlaceInfoData  `json:"placeInfo"`
}

type OrderInfoResp struct {
	Code    int64     `json:"code"`
	Message string    `json:"message"`
	Data    OrderInfo `json:"data"`
}

type PlaceInfoData struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Place     string `json:"place"`
	Status    string `json:"status"`
	Pic       string `json:"pic"`
	Phone     string `json:"phone"`
	Principal string `json:"principal"`
}

type AddPlaceReq struct {
	Name      string `json:"name"`
	Place     string `json:"place"`
	Status    string `json:"status"`
	Pic       string `json:"pic"`
	Phone     string `json:"phone"`
	Principal string `json:"principal"`
}

type AddPlaceResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UpdatePlaceReq struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,optional"`
	Place     string `json:"place,optional"`
	Status    string `json:"status,optional"`
	Pic       string `json:"pic,optional"`
	Phone     string `json:"phone,optional"`
	Principal string `json:"principal,optional"`
}

type UpdatePlaceResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeletePlaceReq struct {
	Ids []int64 `json:"ids"`
}

type DeletePlaceResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListPlaceReq struct {
}

type ListPlaceData struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Place     string `json:"place"`
	Status    string `json:"status"`
	Pic       string `json:"pic"`
	Phone     string `json:"phone"`
	Principal string `json:"principal"`
}

type ListPlaceResp struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Total   int64            `json:"total"`
	Data    []*ListPlaceData `json:"data"`
}

type UploadResp struct {
	Data string `json:"data"`
	Code int64  `json:"code"`
}

type AddReturnApplyReq struct {
	UserID         int64   `json:"userID"`
	OrderID        int64   `json:"orderID"`
	ReturnReasonID int64   `json:"returnReasonID"`
	Status         string  `json:"status"`
	Description    string  `json:"description"`
	ProofPics      string  `json:"proofPics"`
	ReturnAmount   float64 `json:"returnAmount"`
}

type AddReturnApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UpdateReturnApplyReq struct {
	ID             int64   `json:"id"`
	UserID         int64   `json:"userID,optional"`
	OrderID        int64   `json:"orderID,optional"`
	ReturnReasonID int64   `json:"returnReasonID,optional"`
	Status         string  `json:"status,optional"`
	Description    string  `json:"description,optional"`
	ProofPics      string  `json:"proofPics,optional"`
	ReturnAmount   float64 `json:"returnAmount,optional"`
}

type UpdateReturnApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteReturnApplyReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteReturnApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListReturnApplyReq struct {
}

type ListReturnApplyData struct {
	ID               int64    `json:"id"`
	UserID           int64    `json:"userID"`
	OrderID          int64    `json:"orderID"`
	ReturnReasonID   int64    `json:"returnReasonID"`
	ReturnReasonName string   `json:"returnReasonName"`
	Status           string   `json:"status"`
	Description      string   `json:"description"`
	ProofPics        string   `json:"proofPics"`
	ReturnAmount     float64  `json:"returnAmount"`
	Order            Order    `json:"order"`
	User             UserData `json:"user"`
}

type ListReturnApplyResp struct {
	Code    int64                 `json:"code"`
	Message string                `json:"message"`
	Total   int64                 `json:"total"`
	Data    []ListReturnApplyData `json:"data"`
}

type OrderData struct {
	ID                    int64   `json:"id"`
	MemberId              int64   `json:"memberId"`
	PlaceId               int64   `json:"placeId"`
	CouponId              int64   `json:"couponId"`
	OrderSn               string  `json:"orderSn"`               // 订单编号
	MemberUsername        string  `json:"memberUserName"`        // 用户帐号
	TotalAmount           float64 `json:"totalAmount"`           // 订单总金额
	PayAmount             float64 `json:"payAmount"`             // 应付金额（实际支付金额）
	FreightAmount         float64 `json:"freightAmount"`         // 运费金额
	CouponAmount          float64 `json:"couponAmount"`          // 优惠券抵扣金额
	PayType               string  `json:"payType"`               // 支付方式：0->未支付；1->支付宝；2->微信
	Status                string  `json:"status"`                // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             string  `json:"orderType"`             // 订单类型：0->正常订单；1->秒杀订单
	ReceiverProvince      string  `json:"receiverProvince"`      // 省份/直辖市
	ReceiverName          string  `json:"receiverName"`          //
	ReceiverPhone         string  `json:"receiverPhone"`         //
	ReceiverCity          string  `json:"receiverCity"`          // 城市
	ReceiverRegion        string  `json:"receiverRegion"`        // 区
	ReceiverDetailAddress string  `json:"receiverDetailAddress"` // 详细地址
	Note                  string  `json:"note"`                  // 订单备注
	ConfirmStatus         string  `json:"confirmStatus"`         // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          string  `json:"deleteStatus"`          // 删除状态：0->未删除；1->已删除
	PaymentTime           string  `json:"paymentTime,optional"`  // 支付时间
	DeliveryTime          string  `json:"deliveryTime,optional"` // 发货时间
	ReceiveTime           string  `json:"receiveTime,optional"`  // 确认收货时间
	CommentTime           string  `json:"commentTime,optional"`  // 修改时间
}

type UserData struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
	Gander   string `json:"gender"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	CreateAt string `json:"creat_at"`
	UpdateAt string `json:"update_at"`
	CreateBy string `json:"creat_by"`
	UpdateBy string `json:"update_by"`
}

type SkuData struct {
	ID          int64   `json:"id"`
	ProductID   int64   `json:"productId"`
	ProductName string  `json:"productName"`
	Name        string  `json:"name"`
	Pic         string  `json:"pic"`
	SkuSn       string  `json:"skuSn"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Tag         string  `json:"tag"`
}

type Order struct {
	OrderInfo OrderData `json:"orderInfo"`
	SkuList   []SkuData `json:"skuList"`
}

type ReturnApplyInfoReq struct {
	OrderID int64 `json:"orderID"`
}

type ReturnApplyInfoData struct {
	ID               int64   `json:"id"`
	UserID           int64   `json:"userID"`
	OrderID          int64   `json:"orderID"`
	ReturnReasonID   int64   `json:"returnReasonID"`
	ReturnReasonName string  `json:"returnReasonName"`
	Status           string  `json:"status"`
	Description      string  `json:"description"`
	ProofPics        string  `json:"proofPics"`
	ReturnAmount     float64 `json:"returnAmount"`
}

type ReturnApplyInfoResp struct {
	Code    int64               `json:"code"`
	Message string              `json:"message"`
	Data    ReturnApplyInfoData `json:"data"`
}

type AddReturnReasonReq struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AddReturnReasonResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UpdateReturnReasonReq struct {
	ID     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Status string `json:"status,optional"`
}

type UpdateReturnReasonResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteReturnReasonReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteReturnReasonResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListReturnReasonReq struct {
}

type ListReturnReasonData struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ListReturnReasonResp struct {
	Code    int64                  `json:"code"`
	Message string                 `json:"message"`
	Total   int64                  `json:"total"`
	Data    []ListReturnReasonData `json:"data"`
}
