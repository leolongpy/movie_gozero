package handler

import (
	"net/http"

	"movie_gozero/api/place/internal/logic"
	"movie_gozero/api/place/internal/svc"
	"movie_gozero/api/place/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func hotCitiesByCinemaHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HotCitiesByCinemaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHotCitiesByCinemaLogic(r.Context(), ctx)
		resp, err := l.HotCitiesByCinema(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
