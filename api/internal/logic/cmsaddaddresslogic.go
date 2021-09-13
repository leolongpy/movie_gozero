package logic

import (
	"context"
	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"
	"movie_gozero/rpc/cms/cmsservice"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsaddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsaddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsaddAddressLogic {
	return CmsaddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsaddAddressLogic) CmsaddAddress(req types.AddAddressReq) (*types.AddAddressRsp, error) {
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
