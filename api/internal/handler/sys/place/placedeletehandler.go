package place

import (
	"net/http"

	"SimplePick-Mall-Server/api/internal/logic/sys/place"
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlaceDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeletePlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := place.NewPlaceDeleteLogic(r.Context(), svcCtx)
		resp, err := l.PlaceDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
