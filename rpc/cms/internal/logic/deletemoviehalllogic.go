package logic

import (
	"context"
	"go.uber.org/zap"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteMovieHallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMovieHallLogic {
	return &DeleteMovieHallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMovieHallLogic) DeleteMovieHall(req *pb.DeleteMovieHallReq) (*pb.DeleteMovieHallRsp, error) {
	adminID := req.AdminID
	if adminID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		movieHall, err := db.SelectAllMovieHallByMHID(adminID)
		if err != nil {
			l.Logger.Error("error", "SelectAllMovieHallByMHID", err)
			return nil, errors.ErrorCMSFailed
		}
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID == movieHall.CinemaId {
			return nil, errors.ErrorCMSForbiddenParam
		}
	}
	err = db.DeleteMovieHall(req.MhId)
	if err != nil {
		l.Logger.Error("error", "DeleteMovieHall", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
