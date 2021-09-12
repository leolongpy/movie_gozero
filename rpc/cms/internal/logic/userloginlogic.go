package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *pb.UserLoginReq) (*pb.UserLoginRsp, error) {
	userName := in.User
	password := in.Password
	admin, err := db.SelectAdmin(userName, password)
	if err != nil {
		l.Logger.Error("err", err)
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil {
		return nil, errors.ErrorCMSLogin
	}
	var cinemaName = "超级管理员"
	if admin.CinemaID == 0 {
		cinemaName = "待注册影院"
	}
	if admin.CinemaID > 0 {
		cinemaName, err = db.SelectCinemaName(admin.CinemaID)
		if err != nil {
			l.Logger.Error("err", err)
			return nil, errors.ErrorCinemaFailed
		}
	}
	rsp := pb.UserLoginRsp{}
	rsp.CinemaName = cinemaName
	rsp.AdminID = admin.AuID
	rsp.CinemaID = admin.CinemaID
	rsp.AdminNum = admin.AdminNum
	return &rsp, nil
}
