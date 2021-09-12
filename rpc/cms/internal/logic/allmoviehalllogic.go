package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllMovieHallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllMovieHallLogic {
	return &AllMovieHallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllMovieHallLogic) AllMovieHall(req *pb.AllMovieHallReq) (*pb.AllMovieHallRsp, error) {
	rsp := pb.AllMovieHallRsp{}
	adminID := req.AdminID
	page := req.Page
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
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectMovieHallTotal()
		if err != nil {
			l.Logger.Error("error", "SelectMovieHallTotal", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		movieHalls, err := db.SelectAllMovieHall(page, 20)
		if err != nil {
			l.Logger.Error("error", "SelectAllMovieHall", err)
			return nil, errors.ErrorCMSFailed
		}
		movieHallsPB := []*pb.MovieHall{}
		for _, movieHall := range movieHalls {
			movieHallPB := movieHall.ToProtoMovieHall()
			movieHallsPB = append(movieHallsPB, movieHallPB)
		}
		rsp.MovieHalls = movieHallsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectMovieHallTotalByCinemaID(admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectMovieHallTotalByCinemaID", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		// 根据所属管理员id获取影院id
		movieHalls, err := db.SelectAllMovieHallBycinemaID(page, 20, admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectAllMovieHallBycinemaID", err)
			return nil, errors.ErrorCMSFailed
		}
		movieHallsPB := []*pb.MovieHall{}
		for _, movieHall := range movieHalls {
			movieHallPB := movieHall.ToProtoMovieHall()
			movieHallsPB = append(movieHallsPB, movieHallPB)
		}
		rsp.MovieHalls = movieHallsPB
	}

	return nil, nil
}
