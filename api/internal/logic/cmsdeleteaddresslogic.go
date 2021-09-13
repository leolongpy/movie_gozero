package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsdeleteAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsdeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsdeleteAddressLogic {
	return CmsdeleteAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsdeleteAddressLogic) CmsdeleteAddress(req types.DeleteAddressReq) (*types.DeleteAddressRsp, error) {
	_, err := l.svcCtx.Cms.DeleteAddress(l.ctx, &cmsservice.DeleteAddressReq{
		Id:      req.Id,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteAddressRsp{}, err
	}

	return &types.DeleteAddressRsp{}, nil
}
