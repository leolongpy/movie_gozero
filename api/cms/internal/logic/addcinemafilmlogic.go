package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddCinemaFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddCinemaFilmLogic {
	return AddCinemaFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCinemaFilmLogic) AddCinemaFilm(req types.AddCinemaFilmReq) (*types.AddCinemaFilmRsp, error) {
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
