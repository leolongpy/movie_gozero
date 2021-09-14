package handler

import (
	"net/http"

	"movie_gozero/api/cinema/internal/logic"
	"movie_gozero/api/cinema/internal/svc"
	"movie_gozero/api/cinema/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getCinemaMessageByCidHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCinemaMessageByCidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetCinemaMessageByCidLogic(r.Context(), ctx)
		resp, err := l.GetCinemaMessageByCid(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
