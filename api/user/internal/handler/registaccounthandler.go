package handler

import (
	"movie_gozero/api/user/internal/logic"
	"movie_gozero/api/user/internal/svc"
	"movie_gozero/api/user/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func registAccountHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegistAccountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegistAccountLogic(r.Context(), ctx)
		resp, err := l.RegistAccount(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
