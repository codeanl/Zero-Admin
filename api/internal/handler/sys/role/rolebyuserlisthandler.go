package role

import (
	"net/http"

	"SimplePick-Mall-Server/api/internal/logic/sys/role"
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleByUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleByUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleByUserListLogic(r.Context(), svcCtx)
		resp, err := l.RoleByUserList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
