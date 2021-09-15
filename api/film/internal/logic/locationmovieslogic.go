package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LocationMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLocationMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) LocationMoviesLogic {
	return LocationMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LocationMoviesLogic) LocationMovies(req types.LocationMoviesReq) (*types.LocationMoviesRsp, error) {
	resp, err := l.svcCtx.Film.LocationMovies(l.ctx, &filmserviceext.LocationMoviesReq{})

	if err != nil {
		return &types.LocationMoviesRsp{}, err
	}

	items := []*types.Movie{}
	for _, v := range resp.Movies {
		item := &types.Movie{
			Img:               v.Img,
			MovieId:           v.MovieId,
			TitleCn:           v.TitleCn,
			MoviesSupportType: v.MoviesSupportType,
			FilmDirector:      v.FilmDirector,
			Actors:            v.Actors,
			FilmDrama:         v.FilmDrama,
			RatingFinal:       float64(v.RatingFinal),
		}
		items = append(items, item)
	}

	return &types.LocationMoviesRsp{
		Movies: items,
	}, nil
}
