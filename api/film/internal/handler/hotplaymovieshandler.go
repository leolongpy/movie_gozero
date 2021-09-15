package handler

import (
	"net/http"

	"movie_gozero/api/film/internal/logic"
	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func hotPlayMoviesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HotPlayMoviesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHotPlayMoviesLogic(r.Context(), ctx)
		resp, err := l.HotPlayMovies(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
