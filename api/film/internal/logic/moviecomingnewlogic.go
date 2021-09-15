package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieComingNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieComingNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) MovieComingNewLogic {
	return MovieComingNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieComingNewLogic) MovieComingNew(req types.MovieComingNewReq) (*types.MovieComingNewRsp, error) {
	resp, err := l.svcCtx.Film.MovieComingNew(l.ctx, &filmserviceext.MovieComingNewReq{})

	if err != nil {
		return &types.MovieComingNewRsp{}, err
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

	return &types.MovieComingNewRsp{
		Movies: items,
	}, nil
}
