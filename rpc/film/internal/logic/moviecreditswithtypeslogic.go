package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieCreditsWithTypesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMovieCreditsWithTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieCreditsWithTypesLogic {
	return &MovieCreditsWithTypesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取导演与演员信息
func (l *MovieCreditsWithTypesLogic) MovieCreditsWithTypes(req *pb.MovieCreditsWithTypesReq) (*pb.MovieCreditsWithTypesRsp, error) {
	rsp := &pb.MovieCreditsWithTypesRsp{}
	movieId := req.MovieId
	persons := []*pb.Persons{}
	directors, err := db.SelectActors(movieId, 2)
	if err != nil {
		l.Logger.Error("error", err)
		return rsp, errors.ErrorFilmFailed
	}

	for _, director := range directors {
		directorPB := pb.Persons{
			Name:   director.NameCN,
			NameEn: director.NameEN,
			Image:  director.ActorPhoto,
		}
		persons = append(persons, &directorPB)
	}

	typE := pb.Type{
		Persons:    persons,
		TypeName:   "导演",
		TypeNameEc: "Director",
	}

	actors, err := db.SelectActors(movieId, 1)
	persons1 := []*pb.Persons{}

	for _, actor := range actors {
		actor := pb.Persons{
			Name:   actor.NameCN,
			NameEn: actor.NameEN,
			Image:  actor.ActorPhoto,
		}
		persons1 = append(persons1, &actor)
	}

	typE1 := pb.Type{
		Persons:    persons1,
		TypeName:   "演员",
		TypeNameEc: "Director",
	}
	types := []*pb.Type{}
	types = append(types, &typE)
	types = append(types, &typE1)
	rsp.Types = types
	return rsp, nil
}
