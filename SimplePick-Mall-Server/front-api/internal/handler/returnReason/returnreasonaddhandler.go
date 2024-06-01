package returnReason

import (
	"net/http"

	"SimplePick-Mall-Server/front-api/internal/logic/returnReason"
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReturnReasonAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReturnReasonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := returnReason.NewReturnReasonAddLogic(r.Context(), svcCtx)
		resp, err := l.ReturnReasonAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
