package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsuserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsuserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsuserLoginLogic {
	return CmsuserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsuserLoginLogic) CmsuserLogin(req types.UserLoginReq) (*types.UserLoginRsp, error) {
	_, err := l.svcCtx.Cms.UserLogin(l.ctx, &cmsservice.UserLoginReq{
		User:     req.User,
		Password: req.Password,
	})
	if err != nil {
		return &types.UserLoginRsp{}, err
	}

	return &types.UserLoginRsp{}, nil
}
