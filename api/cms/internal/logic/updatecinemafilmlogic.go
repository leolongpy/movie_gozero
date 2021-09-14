package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateCinemaFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateCinemaFilmLogic {
	return UpdateCinemaFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCinemaFilmLogic) UpdateCinemaFilm(req types.UpdateCinemaFilmReq) (*types.UpdateCinemaFilmRsp, error) {
	_, err := l.svcCtx.Cms.UpdateCinemaFilm(l.ctx, &cmsservice.UpdateCinemaFilmReq{
		CinemaID:         req.CinemaID,
		FilmID:           req.FilmID,
		HallID:           req.HallID,
		FilmName:         req.FilmName,
		CinemaName:       req.CinemaName,
		ReleaseTimeYear:  req.ReleaseTimeYear,
		ReleaseTimeMonth: req.ReleaseTimeMonth,
		ReleaseTimeDay:   req.ReleaseTimeDay,
		ReleaseTime:      req.ReleaseTime,
		ReleaseType:      req.ReleaseType,
		ReleaseAdd:       req.ReleaseAdd,
		AdminID:          req.AdminID,
		Length:           req.Length,
		ReleaseDiscount:  req.ReleaseDiscount,
		CfID:             req.CfID,
	})
	if err != nil {
		return &types.UpdateCinemaFilmRsp{}, err
	}

	return &types.UpdateCinemaFilmRsp{}, nil
}
