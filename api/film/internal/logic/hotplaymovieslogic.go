package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotPlayMoviesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotPlayMoviesLogic(ctx context.Context, svcCtx *svc.ServiceContext) HotPlayMoviesLogic {
	return HotPlayMoviesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotPlayMoviesLogic) HotPlayMovies(req types.HotPlayMoviesReq) (*types.HotPlayMoviesRsp, error) {
	resp, err := l.svcCtx.Film.HotPlayMovies(l.ctx, &filmserviceext.HotPlayMoviesReq{
		Loaction: req.Loaction,
	})

	if err != nil {
		return &types.HotPlayMoviesRsp{}, err
	}

	items := []types.Movie{}
	for _, v := range resp.Movies {
		item := types.Movie{
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

	return &types.HotPlayMoviesRsp{
		Movies: items,
	}, nil
}
