package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCinemaFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteCinemaFilmLogic {
	return DeleteCinemaFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCinemaFilmLogic) DeleteCinemaFilm(req types.DeleteCinemaFilmReq) (*types.DeleteCinemaFilmRsp, error) {
	_, err := l.svcCtx.Cms.DeleteCinemaFilm(l.ctx, &cmsservice.DeleteCinemaFilmReq{
		AdminID: req.AdminID,
		CfId:    req.CfId,
	})
	if err != nil {
		return &types.DeleteCinemaFilmRsp{}, err
	}

	return &types.DeleteCinemaFilmRsp{}, nil
}
