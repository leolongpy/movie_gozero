package logic

import (
	"context"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type PayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PayOrder 支付
func (l *PayOrderLogic) PayOrder(req *pb.PayOrderReq) (*pb.PayOrderRsp, error) {
	rsp := &pb.PayOrderRsp{}
	err := db.UpdateOrderStatus(req.OrderNum, req.UserId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}

	err = db.UpdateUserPhone(req.UserId, req.Phone)
	if err != nil {
		l.Logger.Error("err", "UpdateUserPhone", err)
		return rsp, errors.ErrorOrderFailed
	}
	return rsp, nil
}
