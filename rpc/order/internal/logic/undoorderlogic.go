package logic

import (
	"context"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UndoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUndoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UndoOrderLogic {
	return &UndoOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UndoOrder 取消订单
func (l *UndoOrderLogic) UndoOrder(in *pb.UndoOrderReq) (*pb.UndoOrderRsp, error) {

	return &pb.UndoOrderRsp{}, nil
}
