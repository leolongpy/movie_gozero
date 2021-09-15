package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotPlayMoviesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotPlayMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotPlayMoviesLogic {
	return &HotPlayMoviesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取正在售票的影片
func (l *HotPlayMoviesLogic) HotPlayMovies(in *pb.HotPlayMoviesReq) (*pb.HotPlayMoviesRsp, error) {
	rsp := &pb.HotPlayMoviesRsp{}
	films, err := db.SelectTickingFilims(1)
	if err != nil {
		l.Logger.Error("err", "films", err)
		return rsp, err
	}
	MoviesPB := []*pb.Movie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return rsp, err
		}

		filmPB := film.ToProtoDBMovies()
		tmp := ""
		if film.Is3D == 1 {
			tmp = tmp + "3D" + "|"
		}
		if film.IsDMAX == 1 {
			tmp = tmp + "DMAX" + "|"
		}
		if film.IsIMAX == 1 {
			tmp = tmp + "IMAX" + "|"
		}
		if film.IsIMAX3D == 1 {
			tmp = tmp + "IMAX3D" + "|"
		}
		if tmp != "" {
			tmp = tmp[0 : len(tmp)-1]
		}
		filmPB.MoviesSupportType = tmp
		actors := ""
		for _, filmActor := range filmActors {
			actors = actors + filmActor.ActorName + " "
		}
		filmPB.Actors = actors
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return rsp, nil
}
