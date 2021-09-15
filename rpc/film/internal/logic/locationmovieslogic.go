package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LocationMoviesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLocationMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LocationMoviesLogic {
	return &LocationMoviesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  正在热映的影片
func (l *LocationMoviesLogic) LocationMovies(in *pb.LocationMoviesReq) (*pb.LocationMoviesRsp, error) {
	rsp := &pb.LocationMoviesRsp{}
	films, err := db.SelectTickingFilims(1)
	if err != nil {
		l.Logger.Error("error", err)
		return rsp, errors.ErrorFilmFailed
	}
	MoviesPB := []*pb.Movie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors, err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			l.Logger.Error("error", err)
			return rsp, errors.ErrorFilmFailed
		}
		for _, filmActor := range filmActors {
			film.ActorName = append(film.ActorName, filmActor.ActorName)
		}
		// 处理影片种类信息

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return rsp, nil
}
