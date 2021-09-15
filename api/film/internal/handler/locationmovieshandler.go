package handler

import (
	"net/http"

	"movie_gozero/api/film/internal/logic"
	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func locationMoviesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LocationMoviesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLocationMoviesLogic(r.Context(), ctx)
		resp, err := l.LocationMovies(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
