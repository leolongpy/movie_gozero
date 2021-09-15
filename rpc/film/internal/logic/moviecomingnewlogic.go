package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieComingNewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMovieComingNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieComingNewLogic {
	return &MovieComingNewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  即将上映的影片
func (l *MovieComingNewLogic) MovieComingNew(in *pb.MovieComingNewReq) (*pb.MovieComingNewRsp, error) {
	rsp := &pb.MovieComingNewRsp{}
	films, err := db.SelectTickingFilims(2)
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
		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return rsp, nil
}
