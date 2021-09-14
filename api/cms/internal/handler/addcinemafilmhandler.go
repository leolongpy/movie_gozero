package handler

import (
	"net/http"

	"movie_gozero/api/cms/internal/logic"
	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func addCinemaFilmHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCinemaFilmReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddCinemaFilmLogic(r.Context(), ctx)
		resp, err := l.AddCinemaFilm(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
