package logic

import (
	"context"
	"movie_gozero/rpc/cinema/cinemaserviceext"

	"movie_gozero/api/cinema/internal/svc"
	"movie_gozero/api/cinema/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LocationCinemaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLocationCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) LocationCinemaLogic {
	return LocationCinemaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LocationCinemaLogic) LocationCinema(req types.LocationCinemaReq) (*types.LocationCinemaRsp, error) {
	resp, err := l.svcCtx.Cinema.LocationCinema(l.ctx, &cinemaserviceext.LocationCinemaReq{
		LocationId: req.LocationId,
	})
	if err != nil {
		return &types.LocationCinemaRsp{}, err
	}
	items := []*types.Cinema{}
	for _, v := range resp.Cinemas {
		item := &types.Cinema{
			CinemaName:     v.CinemaName,
			CinemaAddress:  v.CinemaAddress,
			CinemaSupport:  v.CinemaSupport,
			CinemaCard:     v.CinemaCard,
			CinemaMinPrice: v.CinemaMinPrice,
			CinemaDiscount: v.CinemaDiscount,
			CinemaId:       v.CinemaId,
		}
		items = append(items, item)
	}

	return &types.LocationCinemaRsp{
		Cinemas: items,
	}, nil
}
