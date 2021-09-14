package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteFilmLogic {
	return DeleteFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFilmLogic) DeleteFilm(req types.DeleteFilmReq) (*types.DeleteFilmRsp, error) {
	_, err := l.svcCtx.Cms.DeleteFilm(l.ctx, &cmsservice.DeleteFilmReq{
		MovieID: req.MovieID,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteFilmRsp{}, err
	}

	return &types.DeleteFilmRsp{}, nil
}
