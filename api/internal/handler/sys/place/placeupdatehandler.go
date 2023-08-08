package place

import (
	"net/http"

	"SimplePick-Mall-Server/api/internal/logic/sys/place"
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlaceUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := place.NewPlaceUpdateLogic(r.Context(), svcCtx)
		resp, err := l.PlaceUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
