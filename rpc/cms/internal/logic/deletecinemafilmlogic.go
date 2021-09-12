package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCinemaFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCinemaFilmLogic {
	return &DeleteCinemaFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCinemaFilmLogic) DeleteCinemaFilm(req *pb.DeleteCinemaFilmReq) (*pb.DeleteCinemaFilmRsp, error) {
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
	err = db.DeleteCinemaFilm(req.CfId)
	if err != nil {
		l.Logger.Error("error", "DeleteCinemaFilm", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
