package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsaddMovieHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsaddMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsaddMovieHallLogic {
	return CmsaddMovieHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsaddMovieHallLogic) CmsaddMovieHall(req types.AddMovieHallReq) (*types.AddMovieHallRsp, error) {
	_, err := l.svcCtx.Cms.AddMovieHall(l.ctx, &cmsservice.AddMovieHallReq{
		AdminID:   req.AdminID,
		MhName:    req.MhName,
		MhAddress: req.MhAddress,
		CinemaId:  req.CinemaId,
	})
	if err != nil {
		return &types.AddMovieHallRsp{}, err
	}
	return &types.AddMovieHallRsp{}, nil
}
