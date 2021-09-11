package logic

import (
	"context"
	"fmt"
	"movie_gozero/rpc/user/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegistAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistAccountLogic {
	return &RegistAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RegistAccount 注册用户
func (l *RegistAccountLogic) RegistAccount(in *pb.RegistAccountReq) (*pb.RegistAccountRsp, error) {
	// todo: add your logic here and delete this line

	userName := in.UserName
	password := in.Password
	email := in.Email
	user, err := db.SelectUserByEmail(email)
	fmt.Printf("user:%#v", user)
	if err != nil {
		l.Logger.Error("error", err)
		return &pb.RegistAccountRsp{}, errors.ErrorUserFailed
	}
	if user != nil {
		return &pb.RegistAccountRsp{}, errors.ErrorUserAlready
	}
	err = db.InsertUser(userName, password, email)
	if err != nil {
		l.Logger.Error("error", err)
		return &pb.RegistAccountRsp{}, errors.ErrorUserFailed
	}
	return &pb.RegistAccountRsp{}, nil
}
