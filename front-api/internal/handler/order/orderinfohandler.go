package order

import (
	"net/http"

	"SimplePick-Mall-Server/front-api/internal/logic/order"
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderInfoLogic(r.Context(), svcCtx)
		resp, err := l.OrderInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
