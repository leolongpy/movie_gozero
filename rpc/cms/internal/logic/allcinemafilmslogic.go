package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCinemaFilmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllCinemaFilmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllCinemaFilmsLogic {
	return &AllCinemaFilmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllCinemaFilmsLogic) AllCinemaFilms(req *pb.AllCinemaFilmsReq) (*pb.AllCinemaFilmsRsp, error) {
	rsp := pb.AllCinemaFilmsRsp{}
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
		total, err := db.SelecttAllCinemaFilmTotal()
		if err != nil {
			l.Logger.Error("error", "SelecttAllCinemaFilmTotal", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		cinemaFilms, err := db.SelectAllCinemaFilm(page, 20)
		if err != nil {
			l.Logger.Error("error", "SelectAllCinemaFilm", err)
			return nil, errors.ErrorCMSFailed
		}
		cinemaFilmsPB := []*pb.CinemaFilm{}
		for _, cinemaFilm := range cinemaFilms {
			cinemaFilmPB := cinemaFilm.ToProtoCinemaFilm()
			cinemaFilmsPB = append(cinemaFilmsPB, cinemaFilmPB)
		}
		rsp.CinemaFilms = cinemaFilmsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectAllCinemaFilmTotalByCinemaID(admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectAllCinemaFilmTotalByCinemaID", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		// 根据所属管理员id获取影院id
		cinemaFilms, err := db.SelectAllCinemaFilmByCinemaID(page, 20, admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectAllCinemaFilmByCinemaID", err)
			return nil, errors.ErrorCMSFailed
		}
		cinemaFilmsPB := []*pb.CinemaFilm{}
		for _, cinemaFilm := range cinemaFilms {
			cinemaFilmPB := cinemaFilm.ToProtoCinemaFilm()
			hallName, err := db.SelectMovieHallHallName(cinemaFilm.HallId)
			if err != nil {
				l.Logger.Error("error", "SelectMovieHallHallName", err)
				return nil, errors.ErrorCMSFailed
			}
			cinemaFilmPB.CinemaName = hallName
			cinemaFilmsPB = append(cinemaFilmsPB, cinemaFilmPB)
		}
		rsp.CinemaFilms = cinemaFilmsPB
	}

	return nil, nil
}
