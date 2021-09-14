package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllAddressLogic {
	return AllAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllAddressLogic) AllAddress(req types.AllAddressReq) (*types.AllAddressRsp, error) {
	resp, err := l.svcCtx.Cms.AllAddress(l.ctx, &cmsservice.AllAddressReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllAddressRsp{}, err
	}
	places := []*types.PlaceAll{}
	for _, v := range resp.Places {
		placeAll := &types.PlaceAll{
			Id:          v.Id,
			Count:       v.Count,
			Name:        v.Name,
			PinyinFull:  v.PinyinFull,
			PinyinShort: v.PinyinShort,
		}
		places = append(places, placeAll)
	}
	return &types.AllAddressRsp{
		Places: places,
		Total:  resp.Total,
	}, nil
}
