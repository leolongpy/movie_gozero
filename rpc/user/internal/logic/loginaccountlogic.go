package logic

import (
	"context"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/user/internal/db"
	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginAccountLogic {
	return &LoginAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginAccount 用户登录
func (l *LoginAccountLogic) LoginAccount(in *pb.LoginAccountReq) (*pb.LoginAccountRsp, error) {
	// todo: add your logic here and delete this line
	email := in.Email
	password := in.Password
	user, err := db.SelectUserByPasswordName(email, password)
	if err != nil {
		l.Logger.Error("error:", err)
		return nil, errors.ErrorUserFailed
	}
	if user == nil {
		return nil, errors.ErrorUserLoginFailed
	}
	out := &pb.LoginAccountRsp{}
	out.Email = user.Email
	out.Phone = user.Phone
	out.UserID = user.UserId
	out.UserName = user.UserName
	return out, nil

}
