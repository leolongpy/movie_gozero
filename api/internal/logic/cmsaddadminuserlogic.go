package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsaddAdminUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsaddAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsaddAdminUserLogic {
	return CmsaddAdminUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsaddAdminUserLogic) CmsaddAdminUser(req types.AddAdminUserReq) (*types.AddAdminUserRsp, error) {
	_, err := l.svcCtx.Cms.AddAdminUser(l.ctx, &cmsservice.AddAdminUserReq{
		AdminID:       req.AdminID,
		AdminName:     req.AdminName,
		AdminPassword: req.AdminPassword,
		AdminCinemaID: req.AdminCinemaID,
		AdminNum:      req.AdminNum,
	})
	if err != nil {
		return &types.AddAdminUserRsp{}, err
	}
	return &types.AddAdminUserRsp{}, nil
}
