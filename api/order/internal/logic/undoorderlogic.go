package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UndoOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUndoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) UndoOrderLogic {
	return UndoOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UndoOrderLogic) UndoOrder(req types.UndoOrderReq) (*types.UndoOrderRsp, error) {
	_, err := l.svcCtx.Order.UndoOrder(l.ctx, &orderserviceext.UndoOrderReq{
		OrderId: req.OrderId,
	})

	if err != nil {
		return &types.UndoOrderRsp{}, err
	}

	return &types.UndoOrderRsp{}, nil
}
