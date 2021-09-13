package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsregisterCinemaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsregisterCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsregisterCinemaLogic {
	return CmsregisterCinemaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsregisterCinemaLogic) CmsregisterCinema(req types.RegisterCinemaReq) (*types.RegisterCinemaRsp, error) {
	resp, err := l.svcCtx.Cms.RegisterCinema(l.ctx, &cmsservice.RegisterCinemaReq{
		AdminID:        req.AdminID,
		CinemaName:     req.CinemaName,
		CinemaAddress:  req.CinemaAddress,
		LocationID:     req.LocationID,
		CinemaTypes:    req.CinemaTypes,
		CinemaCard:     req.CinemaCard,
		CinemaMinPrice: req.CinemaMinPrice,
		CinemaSupport:  req.CinemaSupport,
		CinemaDiscount: req.CinemaDiscount,
		CinemaPhone:    req.CinemaPhone,
	})
	if err != nil {
		return &types.RegisterCinemaRsp{}, err
	}

	return &types.RegisterCinemaRsp{
		CinemaID: resp.CinemaID,
	}, nil
}
