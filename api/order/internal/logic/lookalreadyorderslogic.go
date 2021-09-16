package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LookAlreadyOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAlreadyOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) LookAlreadyOrdersLogic {
	return LookAlreadyOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAlreadyOrdersLogic) LookAlreadyOrders(req types.LookAlreadyOrdersReq) (*types.LookAlreadyOrdersRsp, error) {
	resp, err := l.svcCtx.Order.LookAlreadyOrders(l.ctx, &orderserviceext.LookAlreadyOrdersReq{
		UserId: req.UserId,
	})

	if err != nil {
		return &types.LookAlreadyOrdersRsp{}, err
	}

	items := []*types.AlreadyMovie{}
	for _, v := range resp.Movies {
		item := &types.AlreadyMovie{
			FilmImg:    v.FilmImg,
			FilmName:   v.FilmName,
			Time:       v.Time,
			Director:   v.Director,
			ActorNames: v.ActorNames,
			OrderNum:   v.OrderNum,
		}
		items = append(items, item)
	}

	return &types.LookAlreadyOrdersRsp{
		TotalMovie:   resp.TotalMovie,
		OneNoComment: resp.OneNoComment,
		Movies:       items,
	}, nil
}
