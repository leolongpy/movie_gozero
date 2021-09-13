package handler

import (
	"net/http"

	"movie_gozero/api/internal/logic"
	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func cmsaddCinemaFilmHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCinemaFilmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCmsaddCinemaFilmLogic(r.Context(), ctx)
		resp, err := l.CmsaddCinemaFilm(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
