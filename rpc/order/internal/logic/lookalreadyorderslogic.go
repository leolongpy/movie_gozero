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

type LookAlreadyOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAlreadyOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAlreadyOrdersLogic {
	return &LookAlreadyOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LookAlreadyOrders 查看所有看过的电影票
func (l *LookAlreadyOrdersLogic) LookAlreadyOrders(req *pb.LookAlreadyOrdersReq) (*pb.LookAlreadyOrdersRsp, error) {
	rsp := &pb.LookAlreadyOrdersRsp{}
	userId := req.UserId
	orders, err := db.SelectLookAlreadyOrders(userId)
	if err != nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	var oneNoComment int64 = 0
	movies := []*pb.AlreadyMovie{}
	for _, order := range orders {
		actorNames := []string{}
		film, err := db.SelectFilmDetail(order.MovieId)
		if err != nil {
			l.Logger.Error("err", "order", err)
			return rsp, errors.ErrorOrderFailed
		}
		actors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			l.Logger.Error("err", "order", err)
			return rsp, errors.ErrorOrderFailed
		}
		for _, actor := range actors {
			actorNames = append(actorNames, actor.ActorName)
		}
		movie := &pb.AlreadyMovie{
			FilmImg:    film.Img,
			FilmName:   film.TitleCn,
			Time:       fmt.Sprintf("%d-%d-%d", film.RYear, film.RMonth, film.RDay),
			Director:   film.FilmDirector,
			ActorNames: actorNames,
			OrderNum:   order.OrderNum,
		}
		movies = append(movies, movie)
		if order.OrderScore == -1 {
			oneNoComment = oneNoComment + 1
		}
	}
	rsp.Movies = movies
	rsp.OneNoComment = oneNoComment
	rsp.TotalMovie = int64(len(orders))
	return rsp, nil
}
