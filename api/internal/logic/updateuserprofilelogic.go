package logic

import (
	"context"
	"movie_gozero/rpc/user/userservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserProfileLogic {
	return UpdateUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserProfileLogic) UpdateUserProfile(req types.UpdateUserProfileReq) (*types.UpdateUserProfileRsp, error) {
	var _, err = l.svcCtx.User.UpdateUserProfile(l.ctx, &userservice.UpdateUserProfileReq{
		UserImage: req.UserImage,
		UserName:  req.UserName,
		UserEmail: req.UserEmail,
		UserPhone: req.UserPhone,
		UserID:    req.UserID,
	})
	if err != nil {
		return &types.UpdateUserProfileRsp{}, err
	}
	return &types.UpdateUserProfileRsp{}, nil
}
