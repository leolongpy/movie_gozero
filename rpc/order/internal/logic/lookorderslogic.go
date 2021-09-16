package logic

import (
	"context"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LookOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookOrdersLogic {
	return &LookOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看所有电影票
func (l *LookOrdersLogic) LookOrders(req *pb.LookOrdersReq) (*pb.LookOrdersRsp, error) {
	rsp := &pb.LookOrdersRsp{}
	userId := req.UserId
	orders, err := db.SelectOrderNumMovieIdMHIdStartTime(userId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	movieTicketsPB := []*pb.MovieTicket{}

	for _, order := range orders {
		cinemafilm, err := db.SelectFilmNameCinemaName(order.MhId, order.MovieId)
		if err != nil {
			return rsp, errors.ErrorOrderFailed
		}
		movieTicketPB := &pb.MovieTicket{
			FilmName:  cinemafilm.FilmName,
			Cinema:    cinemafilm.CinemaName,
			StartTime: order.StartTime,
			OrderNum:  order.OrderNum,
		}
		movieTicketsPB = append(movieTicketsPB, movieTicketPB)
	}
	rsp.MovieTickets = movieTicketsPB
	return rsp, nil
}
