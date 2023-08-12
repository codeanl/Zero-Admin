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
	err := l.svcCtx.SkuModel.UpdateSku(in.ID, &model.Sku{
		ProductID:   in.ProductID,
		Name:        in.Name,
		Pic:         in.Pic,
		Description: in.Description,
		Stock:       in.Stock,
		Price:       in.Price,
		Tag:         in.Tag,
		SkuSn:       in.SkuSn,
	})
	var data []string
	_ = l.svcCtx.SkuAttributeModel.DeleteSkuAttributeBySpuID(in.ID)
	for _, i := range in.AttributeShopValueID {
		//
		_, _ = l.svcCtx.SkuAttributeModel.AddSkuAttribute(&model.SkuAttribute{SkuID: in.ID, SkuSizeValueID: i})
		//
		info1, _ := l.svcCtx.SpuSizeValueModel.GetSizeValueByID(i)
		info2, _ := l.svcCtx.SpuSizeModel.GetSizeByID(info1.SizeID)
		nn := fmt.Sprintf(`{"%s": "%s"}`, info2.Name, info1.Value)
		data = append(data, nn)
	}
	combined := strings.Join(data, ", ")
	_ = l.svcCtx.SkuModel.UpdateSku(in.ID, &model.Sku{
		Tag: combined,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &pms.SkuUpdateResp{}, nil
}
