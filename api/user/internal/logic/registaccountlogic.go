package logic

import (
	"context"
	"movie_gozero/api/user/internal/svc"
	"movie_gozero/api/user/internal/types"
	"movie_gozero/rpc/user/userservice"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegistAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegistAccountLogic {
	return RegistAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistAccountLogic) RegistAccount(req types.RegistAccountReq) (*types.RegistAccountRsp, error) {

	_, err := l.svcCtx.User.RegistAccount(l.ctx, &userservice.RegistAccountReq{
		Email:    req.Email,
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return &types.RegistAccountRsp{}, err
	}
	return &types.RegistAccountRsp{}, nil

}
