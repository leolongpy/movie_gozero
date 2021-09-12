package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/rpc/cms/internal/entity"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAdminUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminUserLogic {
	return &AddAdminUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAdminUserLogic) AddAdminUser(req *pb.AddAdminUserReq) (*pb.AddAdminUserRsp, error) {
	adminID := req.AdminID
	if adminID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", "SelectAdminByAUID", err)
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return nil, errors.ErrorCMSForbiddenParam
	}
	adminUser := entity.Admin{
		AdminName:     req.AdminName,
		AdminPassword: req.AdminPassword,
		AdminNum:      req.AdminNum,
		CinemaID:      req.AdminCinemaID,
	}
	err = db.AddAdminUser(&adminUser)
	if err != nil {
		l.Logger.Error("error", "AddAdminUser", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
