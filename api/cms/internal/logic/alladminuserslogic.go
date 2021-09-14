package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllAdminUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllAdminUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllAdminUsersLogic {
	return AllAdminUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllAdminUsersLogic) AllAdminUsers(req types.AllAdminUsersReq) (*types.AllAdminUsersRsp, error) {
	resp, err := l.svcCtx.Cms.AllAdminUsers(l.ctx, &cmsservice.AllAdminUsersReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllAdminUsersRsp{}, err
	}
	users := []*types.AdminUser{}
	for _, v := range resp.AdminUsers {
		adminUser := &types.AdminUser{
			AuID:               v.AuID,
			AdminName:          v.AdminName,
			AdminPassword:      v.AdminPassword,
			AdminCinemaID:      v.AdminCinemaID,
			AdminLastLoginTime: v.AdminLastLoginTime,
			AdminNum:           v.AdminNum,
		}
		users = append(users, adminUser)
	}
	return &types.AllAdminUsersRsp{
		AdminUsers: users,
		Total:      resp.Total,
	}, nil
}
