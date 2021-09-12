package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllAdminUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllAdminUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllAdminUsersLogic {
	return &AllAdminUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllAdminUsersLogic) AllAdminUsers(in *pb.AllAdminUsersReq) (*pb.AllAdminUsersRsp, error) {
	adminID := in.AdminID
	page := in.Page
	if adminID == 0 || page == 0 {
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
	total, err := db.SelectAdminTotal()
	if err != nil {
		l.Logger.Error("error", "SelectAdminTotal", err)
		return nil, errors.ErrorCMSFailed

	}
	rsp := pb.AllAdminUsersRsp{}
	rsp.Total = total
	adminUsers, err := db.SelectAllAdmin(page, 20)
	if err != nil {
		l.Logger.Error("error", "SelectAllAdmin", err)
		return nil, errors.ErrorCMSFailed
	}
	adminUsersPB := []*pb.AdminUser{}
	for _, adminUser := range adminUsers {
		adminUserPB := adminUser.ToProtoAdminUser()
		adminUsersPB = append(adminUsersPB, adminUserPB)
	}
	rsp.AdminUsers = adminUsersPB
	return &rsp, nil
}
