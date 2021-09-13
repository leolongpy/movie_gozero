package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsdeleteCinemaFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsdeleteCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsdeleteCinemaFilmLogic {
	return CmsdeleteCinemaFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsdeleteCinemaFilmLogic) CmsdeleteCinemaFilm(req types.DeleteCinemaFilmReq) (*types.DeleteCinemaFilmRsp, error) {
	_, err := l.svcCtx.Cms.DeleteCinemaFilm(l.ctx, &cmsservice.DeleteCinemaFilmReq{
		AdminID: req.AdminID,
		CfId:    req.CfId,
	})
	if err != nil {
		return &types.DeleteCinemaFilmRsp{}, err
	}

	return &types.DeleteCinemaFilmRsp{}, nil
}
