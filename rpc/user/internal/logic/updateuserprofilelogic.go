package logic

import (
	"context"
	"movie_gozero/rpc/user/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserProfileLogic {
	return &UpdateUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserProfile 修改用户信息
func (l *UpdateUserProfileLogic) UpdateUserProfile(in *pb.UpdateUserProfileReq) (*pb.UpdateUserProfileRsp, error) {
	// todo: add your logic here and delete this line

	userEmail := in.UserEmail
	userName := in.UserName
	userPhone := in.UserPhone
	userID := in.UserID
	if userEmail != "" {
		err := db.UpdateUserEmailProfile(userEmail, userID)
		if err != nil {
			l.Logger.Error("error", err)
			return nil, errors.ErrorUserFailed
		}
	}
	if userName != "" {
		err := db.UpdateUserNameProfile(userName, userID)
		if err != nil {
			l.Logger.Error("error", err)
			return nil, errors.ErrorUserFailed
		}
	}
	if userPhone != "" {
		err := db.UpdateUserPhoneProfile(userPhone, userID)
		if err != nil {
			l.Logger.Error("error", err)
			return nil, errors.ErrorUserFailed
		}
	}
	return nil, nil
}
