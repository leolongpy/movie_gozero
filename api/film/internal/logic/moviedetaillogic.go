package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) MovieDetailLogic {
	return MovieDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieDetailLogic) MovieDetail(req types.MovieDetailReq) (*types.MovieDetailRsp, error) {
	resp, err := l.svcCtx.Film.MovieDetail(l.ctx, &filmserviceext.MovieDetailReq{
		LocationId: req.LocationId,
		MovieId:    req.MovieId,
	})

	if err != nil {
		return &types.MovieDetailRsp{}, err
	}

	return &types.MovieDetailRsp{
		Image:         resp.Image,
		TitleCn:       resp.TitleEn,
		TitleEn:       resp.TitleCn,
		Rating:        float64(resp.Rating),
		Year:          resp.Year,
		RunTime:       resp.RunTime,
		CommonSpecial: resp.CommonSpecial,
		Director:      resp.Director,
		Release: types.Release{
			Location: resp.Release.Location,
			Date:     resp.Release.Date,
		},
		Content: resp.Content,
		Type:    resp.Type,
	}, nil
}
