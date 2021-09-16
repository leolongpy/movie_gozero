package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LookOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) LookOrdersLogic {
	return LookOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookOrdersLogic) LookOrders(req types.LookOrdersReq) (*types.LookOrdersRsp, error) {
	resp, err := l.svcCtx.Order.LookOrders(l.ctx, &orderserviceext.LookOrdersReq{
		UserId: req.UserId,
	})

	if err != nil {
		return &types.LookOrdersRsp{}, err
	}

	items := []*types.MovieTicket{}
	for _, v := range resp.MovieTickets {
		item := &types.MovieTicket{
			FilmName:  v.FilmName,
			StartTime: v.StartTime,
			Cinema:    v.Cinema,
			OrderNum:  v.OrderNum,
		}
		items = append(items, item)
	}

	return &types.LookOrdersRsp{
		MovieTickets: nil,
	}, nil
}
