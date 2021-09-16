package logic

import (
	"context"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/utils/errors"
	"strconv"
	"time"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type TicketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketLogic {
	return &TicketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Ticket 下单
func (l *TicketLogic) Ticket(req *pb.TicketReq) (*pb.TicketRsp, error) {
	rsp := &pb.TicketRsp{}
	price, err := db.SelectFilmPrice(req.FilmId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	orderNum := time.Now().Unix()
	err = db.InsertOrder(strconv.Itoa(int(orderNum)), price, req.MhId, req.UserId, req.FilmId, req.X, req.Y, req.StartTime, req.EndTime)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	rsp.OrderNumD = orderNum
	return rsp, nil
}
