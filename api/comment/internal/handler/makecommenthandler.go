package handler

import (
	"net/http"

	"movie_gozero/api/comment/internal/logic"
	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func makeCommentHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MakeCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMakeCommentLogic(r.Context(), ctx)
		resp, err := l.MakeComment(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
