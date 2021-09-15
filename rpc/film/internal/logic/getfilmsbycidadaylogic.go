package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/common"
	"movie_gozero/utils/errors"
	"time"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFilmsByCidADayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilmsByCidADayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilmsByCidADayLogic {
	return &GetFilmsByCidADayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  根据影院id和时间获取正在销售的影片信息
func (l *GetFilmsByCidADayLogic) GetFilmsByCidADay(req *pb.GetFilmsByCidADayReq) (*pb.GetFilmsByCidADayRsp, error) {
	rsp := &pb.GetFilmsByCidADayRsp{}
	cinemaId := req.CinemaId
	filmId := req.FilmId
	dayNum := req.DayNum
	var year int64
	var month int64
	var day int64
	if dayNum == 0 {
		year = int64(time.Now().Year())
		month = common.SwitchMonth(time.Now().Month().String())
		day = int64(time.Now().Day())
	}
	if dayNum == 1 {
		dd, _ := time.ParseDuration("24h")
		tomTime := time.Now().Add(dd)
		year = int64(tomTime.Year())
		month = common.SwitchMonth(tomTime.Month().String())
		day = int64(tomTime.Day())
	}
	if dayNum == 2 {
		dd, _ := time.ParseDuration("48h")
		tomTime := time.Now().Add(dd)
		year = int64(tomTime.Year())
		month = common.SwitchMonth(tomTime.Month().String())
		day = int64(tomTime.Day())
	}

	films, err := db.SelectFilmMessageCidADay(cinemaId, filmId, int64(year), month, int64(day))
	if err != nil {
		l.Logger.Error("error", err)
		return rsp, errors.ErrorFilmFailed
	}
	if films != nil {
		dayMoviesPB := []*pb.DayMovie{}
		for _, film := range films {
			mhName, err := db.SelectMHName(film.HallId)
			if err != nil {
				l.Logger.Error("error", err)
				return rsp, errors.ErrorFilmFailed
			}
			dayMoviePB := pb.DayMovie{
				ReleaseTime:     film.ReleaseTime,
				ReleaseType:     film.ReleaseType,
				ReleaseDiscount: film.ReleaseDiscount,
				Length:          film.Length,
				MhName:          mhName,
				MovieId:         film.FilmId,
				MhId:            film.HallId,
				CinemaId:        film.CinemaId,
				FilmName:        film.FilmName,
			}
			dayMoviesPB = append(dayMoviesPB, &dayMoviePB)
		}
		rsp.DayMovie = dayMoviesPB
	}
	return rsp, nil
}
