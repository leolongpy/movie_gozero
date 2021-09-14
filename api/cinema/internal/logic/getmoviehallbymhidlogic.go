package logic

import (
	"context"
	"movie_gozero/rpc/cinema/cinemaserviceext"

	"movie_gozero/api/cinema/internal/svc"
	"movie_gozero/api/cinema/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMovieHallByMHIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMovieHallByMHIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMovieHallByMHIdLogic {
	return GetMovieHallByMHIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMovieHallByMHIdLogic) GetMovieHallByMHId(req types.GetMovieHallByMHIdReq) (*types.GetMovieHallByMHIdRsp, error) {
	resp, err := l.svcCtx.Cinema.GetMovieHallByMHId(l.ctx, &cinemaserviceext.GetMovieHallByMHIdReq{
		MhId: req.MhId,
	})
	if err != nil {
		return &types.GetMovieHallByMHIdRsp{}, err
	}

	return &types.GetMovieHallByMHIdRsp{
		MhAddress: resp.MhAddress,
	}, nil
}
