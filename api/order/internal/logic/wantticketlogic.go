package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WantTicketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) WantTicketLogic {
	return WantTicketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantTicketLogic) WantTicket(req types.WantTicketReq) (*types.WantTicketRsp, error) {
	_, err := l.svcCtx.Order.WantTicket(l.ctx, &orderserviceext.WantTicketReq{
		UserId: req.UserId,
		FilmId: req.FilmId,
		Type:   req.Type,
	})

	if err != nil {
		return &types.WantTicketRsp{}, err
	}

	return &types.WantTicketRsp{}, nil
}
