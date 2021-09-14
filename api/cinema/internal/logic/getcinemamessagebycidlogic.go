package logic

import (
	"context"
	"movie_gozero/api/cinema/internal/svc"
	"movie_gozero/api/cinema/internal/types"
	"movie_gozero/rpc/cinema/cinemaserviceext"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCinemaMessageByCidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCinemaMessageByCidLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCinemaMessageByCidLogic {
	return GetCinemaMessageByCidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCinemaMessageByCidLogic) GetCinemaMessageByCid(req types.GetCinemaMessageByCidReq) (*types.GetCinemaMessageByCidRsp, error) {
	resp, err := l.svcCtx.Cinema.GetCinemaMessageByCid(l.ctx, &cinemaserviceext.GetCinemaMessageByCidReq{
		CinemaId: req.CinemaId,
	})
	if err != nil {
		return &types.GetCinemaMessageByCidRsp{}, err
	}

	return &types.GetCinemaMessageByCidRsp{
		Cinema: types.Cinema{
			CinemaName:     resp.Cinema.CinemaName,
			CinemaAddress:  resp.Cinema.CinemaAddress,
			CinemaSupport:  resp.Cinema.CinemaSupport,
			CinemaCard:     resp.Cinema.CinemaCard,
			CinemaMinPrice: resp.Cinema.CinemaMinPrice,
			CinemaDiscount: resp.Cinema.CinemaDiscount,
			CinemaId:       resp.Cinema.CinemaId,
		},
		FilmMessage: nil,
	}, nil
}
