package logic

import (
	"context"
	"movie_gozero/rpc/place/internal/db"

	"movie_gozero/rpc/place/internal/svc"
	"movie_gozero/rpc/place/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotCitiesByCinemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotCitiesByCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotCitiesByCinemaLogic {
	return &HotCitiesByCinemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HotCitiesByCinema 获取所有电影院的地点
func (l *HotCitiesByCinemaLogic) HotCitiesByCinema(in *pb.HotCitiesByCinemaReq) (*pb.HotCitiesByCinemaRep, error) {
	res := &pb.HotCitiesByCinemaRep{}
	places, err := db.SelectPlaces()
	if err != nil {
		l.Logger.Error("err", "places", err)
		return res, err
	}
	placesPB := []*pb.Place{}
	for _, place := range places {
		placePB := place.ToProtoDBHotPlayMovies()
		placesPB = append(placesPB, placePB)
	}
	res.P = placesPB
	return res, nil
}
