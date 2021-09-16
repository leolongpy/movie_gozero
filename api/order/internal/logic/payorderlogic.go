package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) PayOrderLogic {
	return PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req types.PayOrderReq) (*types.PayOrderRsp, error) {
	_, err := l.svcCtx.Order.PayOrder(l.ctx, &orderserviceext.PayOrderReq{
		OrderNum: req.OrderNum,
		UserId:   req.UserId,
		Phone:    req.Phone,
	})

	if err != nil {
		return &types.PayOrderRsp{}, err
	}

	return &types.PayOrderRsp{}, nil
}
