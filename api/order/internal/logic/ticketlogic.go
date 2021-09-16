package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TicketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) TicketLogic {
	return TicketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TicketLogic) Ticket(req types.TicketReq) (*types.TicketRsp, error) {
	resp, err := l.svcCtx.Order.Ticket(l.ctx, &orderserviceext.TicketReq{
		UserId:    req.UserId,
		FilmId:    req.FilmId,
		MhId:      req.MhId,
		X:         req.X,
		Y:         req.Y,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})

	if err != nil {
		return &types.TicketRsp{}, err
	}

	return &types.TicketRsp{
		OrderNumD: resp.OrderNumD,
	}, nil
}
