package logic

import (
	"context"
	"movie_gozero/rpc/place/placeserviceext"

	"movie_gozero/api/place/internal/svc"
	"movie_gozero/api/place/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotCitiesByCinemaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotCitiesByCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) HotCitiesByCinemaLogic {
	return HotCitiesByCinemaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotCitiesByCinemaLogic) HotCitiesByCinema(req types.HotCitiesByCinemaReq) (*types.HotCitiesByCinemaRep, error) {
	resp, err := l.svcCtx.Place.HotCitiesByCinema(l.ctx, &placeserviceext.HotCitiesByCinemaReq{})

	if err != nil {
		return &types.HotCitiesByCinemaRep{}, err
	}

	items := []*types.Place{}
	for _, v := range resp.P {
		item := &types.Place{
			Count:       v.Count,
			Id:          v.Id,
			N:           v.N,
			PinyinFull:  v.PinyinFull,
			PinyinShort: v.PinyinShort,
		}
		items = append(items, item)
	}

	return &types.HotCitiesByCinemaRep{
		P: nil,
	}, nil
}
