package logic

import (
	"context"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type WantTicketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWantTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantTicketLogic {
	return &WantTicketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// WantTicket 想看
func (l *WantTicketLogic) WantTicket(req *pb.WantTicketReq) (*pb.WantTicketRsp, error) {
	rsp := &pb.WantTicketRsp{}
	err := db.InsertWantSeeRecord(req.FilmId, req.UserId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	return rsp, nil
}
