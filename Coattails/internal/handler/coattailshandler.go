package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"Coattails/internal/logic"
	"Coattails/internal/svc"
	"Coattails/internal/types"
)

func CoattailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCoattailsLogic(r.Context(), svcCtx)
		resp, err := l.Coattails(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
