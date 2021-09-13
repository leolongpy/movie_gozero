package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsallUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsallUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsallUsersLogic {
	return CmsallUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsallUsersLogic) CmsallUsers(req types.AllUsersReq) (*types.AllUsersRsp, error) {
	resp, err := l.svcCtx.Cms.AllUsers(l.ctx, &cmsservice.AllUsersReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllUsersRsp{}, err
	}
	items := []*types.User{}
	for _, v := range resp.Users {
		item := &types.User{
			UserId:   v.UserId,
			UserName: v.UserName,
			Password: v.Password,
			CreateAt: v.CreateAt,
			Email:    v.Email,
			Phone:    v.Phone,
		}
		items = append(items, item)
	}
	return &types.AllUsersRsp{
		Users: items,
		Total: resp.Total,
	}, nil
}
