package logic

import (
	"context"
	"movie_gozero/rpc/user/userservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginAccountLogic {
	return LoginAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginAccountLogic) LoginAccount(req types.LoginAccountReq) (*types.LoginAccountRsp, error) {

	resp, err := l.svcCtx.User.LoginAccount(l.ctx, &userservice.LoginAccountReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &types.LoginAccountRsp{}, err
	}
	return &types.LoginAccountRsp{
		UserID:   resp.UserID,
		UserName: resp.UserName,
		Email:    resp.Email,
		Phone:    resp.Phone,
	}, nil

}
