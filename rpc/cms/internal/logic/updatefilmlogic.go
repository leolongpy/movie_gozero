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

type UpdateFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFilmLogic {
	return &UpdateFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFilmLogic) UpdateFilm(req *pb.UpdateFilmReq) (*pb.UpdateFilmRsp, error) {
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
	film := entity.Film{
		Img:           req.Img,
		Length:        req.Length,
		FilmPrice:     req.FilmPrice,
		FilmDirector:  req.FilmDirector,
		TitleCn:       req.TitleCn,
		TitleEn:       req.TitleEn,
		Type:          req.Type,
		FilmDrama:     req.FilmDrama,
		CommonSpecial: req.CommonSpecial,
		CompanyIssued: req.CompanyIssued,
		Country:       req.Country,
		MovieId:       req.MovieID,
		RDay:          req.RDay,
		RMonth:        req.RMonth,
		RYear:         req.RYear,
		IsTicking:     req.IsTicking,
	}
	err = db.UpdateFilm(&film)
	if err != nil {
		l.Logger.Error("error", "UpdateFilm", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
