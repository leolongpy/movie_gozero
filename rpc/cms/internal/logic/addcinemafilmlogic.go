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

type AddCinemaFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCinemaFilmLogic {
	return &AddCinemaFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCinemaFilmLogic) AddCinemaFilm(req *pb.AddCinemaFilmReq) (*pb.AddCinemaFilmRsp, error) {
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
		// 非超级管理员只能给自己影院添加影片
		if admin.CinemaID != req.CinemaID {
			return nil, errors.ErrorCMSForbiddenParam
		}
	}
	cinemaFilm := entity.CinemaFilm{
		CinemaId:         req.CinemaID,
		FilmId:           req.MovieID,
		HallId:           req.HallID,
		FilmName:         req.TitleCn,
		CinemaName:       req.CinemaName,
		ReleaseTimeYear:  req.ReleaseTimeYear,
		ReleaseTimeMonth: req.ReleaseTimeMonth,
		ReleaseTimeDay:   req.ReleaseTimeDay,
		ReleaseTime:      req.ReleaseTime,
		ReleaseType:      req.Type,
		ReleaseAdd:       req.ReleaseAdd,
		Length:           req.Length,
		ReleaseDiscount:  req.ReleaseDiscount,
	}
	err = db.AddAllCinemaFilm(&cinemaFilm)
	if err != nil {
		l.Logger.Error("error", "AddAllCinemaFilm", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
