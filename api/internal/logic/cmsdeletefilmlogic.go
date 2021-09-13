package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsdeleteFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsdeleteFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsdeleteFilmLogic {
	return CmsdeleteFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsdeleteFilmLogic) CmsdeleteFilm(req types.DeleteFilmReq) (*types.DeleteFilmRsp, error) {
	_, err := l.svcCtx.Cms.DeleteFilm(l.ctx, &cmsservice.DeleteFilmReq{
		MovieID: req.MovieID,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteFilmRsp{}, err
	}

	return &types.DeleteFilmRsp{}, nil
}
