package handler

import (
	"net/http"

	"movie_gozero/api/comment/internal/logic"
	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func hotCommentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HotCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHotCommentLogic(r.Context(), ctx)
		resp, err := l.HotComment(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
