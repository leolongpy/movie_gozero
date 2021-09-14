package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.UserLoginReq) (*types.UserLoginRsp, error) {
	_, err := l.svcCtx.Cms.UserLogin(l.ctx, &cmsservice.UserLoginReq{
		User:     req.User,
		Password: req.Password,
	})
	if err != nil {
		return &types.UserLoginRsp{}, err
	}

	return &types.UserLoginRsp{}, nil
}
