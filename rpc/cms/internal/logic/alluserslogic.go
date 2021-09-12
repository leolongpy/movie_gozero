package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllUsersLogic {
	return &AllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllUsersLogic) AllUsers(in *pb.AllUsersReq) (*pb.AllUsersRsp, error) {
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
	total, err := db.SelectUserTotal()
	if err != nil {
		l.Logger.Error("error", "SelectTotal", err)
		return nil, errors.ErrorCMSFailed

	}
	rsp := pb.AllUsersRsp{}
	rsp.Total = total
	users, err := db.SelectAllUsers(page, 20)
	if err != nil {
		l.Logger.Error("error", "SelectAllUsers", err)
		return nil, errors.ErrorCMSFailed
	}
	usersPB := []*pb.User{}
	for _, user := range users {
		userPB := user.ToProtoUser()
		usersPB = append(usersPB, userPB)
	}
	rsp.Users = usersPB

	return &rsp, nil
}
