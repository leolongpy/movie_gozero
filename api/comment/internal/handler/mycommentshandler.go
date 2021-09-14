package handler

import (
	"net/http"

	"movie_gozero/api/comment/internal/logic"
	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func myCommentsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MyCommentsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMyCommentsLogic(r.Context(), ctx)
		resp, err := l.MyComments(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
