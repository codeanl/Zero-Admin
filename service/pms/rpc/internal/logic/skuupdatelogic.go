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

type SkuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuUpdateLogic {
	return &SkuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新Sku
func (l *SkuUpdateLogic) SkuUpdate(in *pms.SkuUpdateReq) (*pms.SkuUpdateResp, error) {
	//sku标识
	var data []string
	for _, i := range in.SizeValueID {
		info1, _ := l.svcCtx.AttributeValueModel.GetAttributeValueByID(i)
		info2, _ := l.svcCtx.AttributeModel.GetAttributeByID(info1.AttributeID)
		nn := fmt.Sprintf(`{"%s": "%s"}`, info2.Name, info1.Value)
		data = append(data, nn)
	}
	tag := strings.Join(data, ", ")
	err := l.svcCtx.SkuModel.UpdateSku(in.ID, &model.ProductSku{
		ProductID:   in.ProductID,
		Name:        in.Name,
		Pic:         in.Pic,
		Description: in.Description,
		Stock:       in.Stock,
		Price:       in.Price,
		Tag:         tag,
		SkuSn:       in.SkuSn,
		Sale:        in.Sale,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &pms.SkuUpdateResp{}, nil
}
