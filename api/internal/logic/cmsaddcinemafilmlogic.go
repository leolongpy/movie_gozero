package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsaddCinemaFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsaddCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsaddCinemaFilmLogic {
	return CmsaddCinemaFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsaddCinemaFilmLogic) CmsaddCinemaFilm(req types.AddCinemaFilmReq) (*types.AddCinemaFilmRsp, error) {
	_, err := l.svcCtx.Cms.AddCinemaFilm(l.ctx, &cmsservice.AddCinemaFilmReq{
		CinemaID:         req.CinemaID,
		MovieID:          req.MovieID,
		HallID:           req.HallID,
		TitleCn:          req.TitleCn,
		CinemaName:       req.CinemaName,
		ReleaseTimeYear:  req.ReleaseTimeYear,
		ReleaseTimeMonth: req.ReleaseTimeMonth,
		ReleaseTimeDay:   req.ReleaseTimeDay,
		ReleaseTime:      req.ReleaseTime,
		Type:             req.Type,
		ReleaseAdd:       req.ReleaseAdd,
		AdminID:          req.AdminID,
		Length:           req.Length,
		ReleaseDiscount:  req.ReleaseDiscount,
	})
	if err != nil {
		return &types.AddCinemaFilmRsp{}, err
	}
	return &types.AddCinemaFilmRsp{}, nil
}
