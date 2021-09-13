package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsupdateMovieHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsupdateMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsupdateMovieHallLogic {
	return CmsupdateMovieHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsupdateMovieHallLogic) CmsupdateMovieHall(req types.UpdateMovieHallReq) (*types.UpdateMovieHallRsp, error) {
	_, err := l.svcCtx.Cms.UpdateMovieHall(l.ctx, &cmsservice.UpdateMovieHallReq{
		AdminID:   req.AdminID,
		MhName:    req.MhName,
		MhAddress: req.MhAddress,
		CinemaId:  req.CinemaId,
		MhId:      req.MhId,
	})
	if err != nil {
		return &types.UpdateMovieHallRsp{}, err
	}

	return &types.UpdateMovieHallRsp{}, nil
}
