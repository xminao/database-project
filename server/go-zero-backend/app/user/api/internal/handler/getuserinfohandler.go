package handler

import (
	"backend/util/result"
	"net/http"

	"backend/app/user/api/internal/logic"
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getuserinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetuserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Getuserinfo(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
		result.HttpResult(r, w, resp, err)
	}
}
