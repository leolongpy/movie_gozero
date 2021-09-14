package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCinemaHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllCinemaHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllCinemaHallLogic {
	return AllCinemaHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllCinemaHallLogic) AllCinemaHall(req types.AllCinemaHallReq) (*types.AllCinemaHallRsp, error) {
	resp, err := l.svcCtx.Cms.AllCinemaHall(l.ctx, &cmsservice.AllCinemaHallReq{
		CinemaID: req.CinemaID,
		AdminID:  req.AdminID,
	})
	if err != nil {
		return &types.AllCinemaHallRsp{}, err
	}
	items := []*types.HallAddressList{}
	for _, v := range resp.HallAddresses {
		item := &types.HallAddressList{
			MhID:   v.MhID,
			MhName: v.MhName,
		}
		items = append(items, item)
	}
	return &types.AllCinemaHallRsp{
		HallAddresses: items,
	}, nil
}
