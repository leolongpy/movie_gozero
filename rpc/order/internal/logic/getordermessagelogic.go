package logic

import (
	"context"
	"fmt"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderMessageLogic {
	return &GetOrderMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据订单号获取电影票的信息
func (l *GetOrderMessageLogic) GetOrderMessage(req *pb.GetOrderMessageReq) (*pb.GetOrderMessageRsp, error) {
	rsp := &pb.GetOrderMessageRsp{}
	orderNum := req.OrderNum
	userId := req.UserId
	order, err := db.SelectOrderMessage(orderNum, userId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	film, err := db.SelectFilmMessage(order.MovieId)
	if err != nil {
		l.Logger.Error("err", "film", err)
		return rsp, errors.ErrorOrderFailed
	}
	hall, err := db.SelectMHNameMHId(order.MhId)
	if err != nil {
		l.Logger.Error("err", "hall", err)
		return rsp, errors.ErrorOrderFailed
	}

	ciname, err := db.SelectCinemaByCid(hall.CinemaId)
	if err != nil {
		l.Logger.Error("err", "cinema", err)
		return rsp, errors.ErrorOrderFailed
	}
	user, err := db.SelectUserPhoneByUserId(userId)
	if err != nil {
		l.Logger.Error("err", "user", err)
		return rsp, errors.ErrorOrderFailed
	}
	ticketDetailPB := &pb.TicketDetail{
		FilmName:      film.TitleCn,
		FilmImg:       film.Img,
		StartTime:     order.StartTime,
		EndTime:       order.EndTime,
		CinemaName:    ciname.CinemaName,
		MhName:        hall.MhName,
		Seat:          fmt.Sprintf("%d排%d座", order.OrderX, order.OrderY),
		OrderNum:      orderNum,
		CinemaAddress: ciname.CinemaAdd,
		Price:         order.OrderPrice,
		CreateAt:      order.CreateAt,
		Phone:         user.Phone,
		CinemaPhone:   ciname.CinemaPhone,
	}
	rsp.Ticketdetail = ticketDetailPB
	return rsp, nil
}
