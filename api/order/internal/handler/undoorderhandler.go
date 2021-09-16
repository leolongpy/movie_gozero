package handler

import (
	"net/http"

	"movie_gozero/api/order/internal/logic"
	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func undoOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UndoOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUndoOrderLogic(r.Context(), ctx)
		resp, err := l.UndoOrder(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
