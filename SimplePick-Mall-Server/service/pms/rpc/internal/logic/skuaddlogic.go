package logic

import (
	"SimplePick-Mall-Server/service/pms/model"
	"context"
	"errors"
	"fmt"
	"strings"

	"SimplePick-Mall-Server/service/pms/rpc/internal/svc"
	"SimplePick-Mall-Server/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuAddLogic {
	return &SkuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加Sku
func (l *SkuAddLogic) SkuAdd(in *pms.SkuAddReq) (*pms.SkuAddResp, error) {
	var data []string
	for _, i := range in.SizeValueID {
		info1, _ := l.svcCtx.AttributeValueModel.GetAttributeValueByID(i)
		info2, _ := l.svcCtx.AttributeModel.GetAttributeByID(info1.AttributeID)
		nn := fmt.Sprintf(`{"%s": "%s"}`, info2.Name, info1.Value)
		data = append(data, nn)
	}
	tag := strings.Join(data, ", ")
	info := model.ProductSku{
		ProductID:   in.ProductID,
		Name:        in.Name,
		Pic:         in.Pic,
		SkuSn:       in.SkuSn,
		Description: in.Description,
		Stock:       in.Stock,
		Price:       in.Price,
		Tag:         tag,
	}
	_, err := l.svcCtx.SkuModel.AddSku(&info)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &pms.SkuAddResp{}, nil
}
