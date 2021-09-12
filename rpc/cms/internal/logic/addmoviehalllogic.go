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

type AddMovieHallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMovieHallLogic {
	return &AddMovieHallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMovieHallLogic) AddMovieHall(req *pb.AddMovieHallReq) (*pb.AddMovieHallRsp, error) {
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
		// 非超级管理员只能给自己影院添加影厅
		if admin.CinemaID != req.CinemaId {
			return nil, errors.ErrorCMSForbiddenParam
		}
	}
	movieHall := entity.MovieHall{
		MhName:    req.MhName,
		MhAddress: req.MhAddress,
		CinemaId:  req.CinemaId,
	}
	err = db.AddMovieHall(&movieHall)
	if err != nil {
		l.Logger.Error("error", "AddMovieHall", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
