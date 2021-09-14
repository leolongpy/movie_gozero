package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteAdminUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteAdminUserLogic {
	return DeleteAdminUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAdminUserLogic) DeleteAdminUser(req types.DeleteAdminUserReq) (*types.DeleteAdminUserRsp, error) {
	_, err := l.svcCtx.Cms.DeleteAdminUser(l.ctx, &cmsservice.DeleteAdminUserReq{
		AuID:    req.AuID,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteAdminUserRsp{}, err
	}

	return &types.DeleteAdminUserRsp{}, nil
}
