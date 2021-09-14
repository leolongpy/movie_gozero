package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteAddressLogic {
	return DeleteAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAddressLogic) DeleteAddress(req types.DeleteAddressReq) (*types.DeleteAddressRsp, error) {
	_, err := l.svcCtx.Cms.DeleteAddress(l.ctx, &cmsservice.DeleteAddressReq{
		Id:      req.Id,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteAddressRsp{}, err
	}

	return &types.DeleteAddressRsp{}, nil
}
