package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/errors"
	"strconv"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMovieDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovieDetailLogic {
	return &MovieDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  电影详情
func (l *MovieDetailLogic) MovieDetail(in *pb.MovieDetailReq) (*pb.MovieDetailRsp, error) {
	rsp := &pb.MovieDetailRsp{}
	movieId := in.MovieId
	film, err := db.SelectFilmDetail(movieId)
	if err != nil {
		l.Logger.Error("error", err)
		return rsp, errors.ErrorFilmFailed
	}
	rsp.TitleCn = film.TitleCn
	rsp.TitleEn = film.TitleEn
	rsp.CommonSpecial = film.FilmDrama
	rsp.Content = film.CommonSpecial
	rsp.Image = film.Img
	rsp.Rating = film.RatingFinal
	rsp.RunTime = film.Length
	rsp.Year = film.RYear
	rsp.Type = film.Type
	str1 := strconv.Itoa(int(film.RYear)) + "-" + strconv.Itoa(int(film.RMonth)) + "-" + strconv.Itoa(int(film.RDay))
	str2 := film.Country
	pbs := pb.Release{
		Date:     str1,
		Location: str2,
	}
	rsp.Release = &pbs
	return rsp, nil
}
