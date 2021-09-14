package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddAddressLogic {
	return AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressLogic) AddAddress(req types.AddAddressReq) (*types.AddAddressRsp, error) {
	_, err := l.svcCtx.Cms.AddAddress(l.ctx, &cmsservice.AddAddressReq{
		AdminID:     req.AdminID,
		Count:       req.Count,
		Name:        req.Name,
		PinyinFull:  req.PinyinFull,
		PinyinShort: req.PinyinShort,
	})
	if err != nil {
		return &types.AddAddressRsp{}, err
	}
	return &types.AddAddressRsp{}, nil
}
