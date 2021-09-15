package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetFilmsByCidADayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilmsByCidADayLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetFilmsByCidADayLogic {
	return GetFilmsByCidADayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilmsByCidADayLogic) GetFilmsByCidADay(req types.GetFilmsByCidADayReq) (*types.GetFilmsByCidADayRsp, error) {
	resp, err := l.svcCtx.Film.GetFilmsByCidADay(l.ctx, &filmserviceext.GetFilmsByCidADayReq{
		CinemaId: req.CinemaId,
		FilmId:   req.FilmId,
		DayNum:   req.DayNum,
	})

	if err != nil {
		return &types.GetFilmsByCidADayRsp{}, err
	}

	items := []*types.DayMovie{}
	for _, v := range resp.DayMovie {
		item := &types.DayMovie{
			ReleaseTime:     v.ReleaseTime,
			Length:          v.Length,
			ReleaseType:     v.ReleaseType,
			MhName:          v.MhName,
			ReleaseDiscount: float64(v.ReleaseDiscount),
			MovieId:         v.MovieId,
			MhId:            v.MhId,
			CinemaId:        v.CinemaId,
			FilmName:        v.FilmName,
		}
		items = append(items, item)
	}

	return &types.GetFilmsByCidADayRsp{
		DayMovie: items,
	}, nil
}
