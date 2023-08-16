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
	Children     []ListCategoryData `json:"children"`
}

type ListCategoryResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    []ListCategoryData `json:"data"`
	Total   int64              `json:"total"`
}