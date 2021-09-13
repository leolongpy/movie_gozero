package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsdeleteAdminUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsdeleteAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsdeleteAdminUserLogic {
	return CmsdeleteAdminUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsdeleteAdminUserLogic) CmsdeleteAdminUser(req types.DeleteAdminUserReq) (*types.DeleteAdminUserRsp, error) {
	_, err := l.svcCtx.Cms.DeleteAdminUser(l.ctx, &cmsservice.DeleteAdminUserReq{
		AuID:    req.AuID,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.DeleteAdminUserRsp{}, err
	}

	return &types.DeleteAdminUserRsp{}, nil
}
