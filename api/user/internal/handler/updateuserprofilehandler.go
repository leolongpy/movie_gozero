package handler

import (
	"movie_gozero/api/user/internal/logic"
	"movie_gozero/api/user/internal/svc"
	"movie_gozero/api/user/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func updateUserProfileHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateUserProfileLogic(r.Context(), ctx)
		resp, err := l.UpdateUserProfile(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
