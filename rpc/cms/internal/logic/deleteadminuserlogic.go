package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteAdminUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAdminUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAdminUserLogic {
	return &DeleteAdminUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAdminUserLogic) DeleteAdminUser(req *pb.DeleteAdminUserReq) (*pb.DeleteAdminUserRsp, error) {
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
	err = db.DeleteAdminUser(req.AuID)
	if err != nil {
		l.Logger.Error("error", "DeleteFilm", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
