package logic

import (
	"context"
	"movie_gozero/rpc/user/userservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ResetAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) ResetAccountLogic {
	return ResetAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetAccountLogic) ResetAccount(req types.ResetAccountReq) (*types.ResetAccountRsp, error) {
	var _, err = l.svcCtx.User.ResetAccount(l.ctx, &userservice.ResetAccountReq{})
	if err != nil {
		return &types.ResetAccountRsp{}, err
	}
	return &types.ResetAccountRsp{}, nil
}
