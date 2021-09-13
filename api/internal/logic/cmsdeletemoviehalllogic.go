package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsdeleteMovieHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsdeleteMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsdeleteMovieHallLogic {
	return CmsdeleteMovieHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsdeleteMovieHallLogic) CmsdeleteMovieHall(req types.DeleteMovieHallReq) (*types.DeleteMovieHallRsp, error) {
	_, err := l.svcCtx.Cms.DeleteMovieHall(l.ctx, &cmsservice.DeleteMovieHallReq{
		AdminID: req.AdminID,
		MhId:    req.MhId,
	})
	if err != nil {
		return &types.DeleteMovieHallRsp{}, err
	}

	return &types.DeleteMovieHallRsp{}, nil
}
