package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteMovieHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMovieHallLogic {
	return DeleteMovieHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMovieHallLogic) DeleteMovieHall(req types.DeleteMovieHallReq) (*types.DeleteMovieHallRsp, error) {
	_, err := l.svcCtx.Cms.DeleteMovieHall(l.ctx, &cmsservice.DeleteMovieHallReq{
		AdminID: req.AdminID,
		MhId:    req.MhId,
	})
	if err != nil {
		return &types.DeleteMovieHallRsp{}, err
	}

	return &types.DeleteMovieHallRsp{}, nil
}
