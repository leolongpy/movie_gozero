package logic

import (
	"context"
	"go.uber.org/zap"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/rpc/cms/internal/entity"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFilmLogic {
	return &AddFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFilmLogic) AddFilm(in *pb.AddFilmReq) (*pb.AddFilmRsp, error) {

	adminID := in.AdminID
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
		return nil, errors.ErrorCMSForbiddenParam
	}
	film := entity.Film{
		Img:              in.Img,
		Length:           in.Length,
		FilmPrice:        in.FilmPrice,
		FilmScreenwriter: in.FilmScreenwriter,
		FilmDirector:     in.FilmDirector,
		TitleCn:          in.TitleCn,
		TitleEn:          in.TitleEn,
		Type:             in.Type,
		FilmDrama:        in.FilmDrama,
		CommonSpecial:    in.CommonSpecial,
		CompanyIssued:    in.CompanyIssued,
		Country:          in.Country,
		Is3D:             in.Is3D,
		IsDMAX:           in.IsDMAX,
		IsIMAX:           in.IsIMAX,
		IsIMAX3D:         in.IsIMAX3D,
		RDay:             in.RDay,
		RMonth:           in.RMonth,
		RYear:            in.RYear,
	}
	filmID, err := db.InsertFilm(&film)
	if err != nil {
		l.Logger.Error("error", zap.Any("InsertFilm", err))
		return nil, errors.ErrorCMSFailed
	}
	director := entity.Actor{
		NameCN:     in.FilmDirector,
		ActorPhoto: in.FilmDirectorImg,
	}
	err, directorID := db.InsertActor(&director, 2)
	if err != nil {
		l.Logger.Error("error", "InsertActor", err)
		return nil, errors.ErrorCMSFailed
	}
	if directorID == 0 {
		l.Logger.Error("error", "directorID", directorID)
		return nil, errors.ErrorCMSFailed
	}
	err = db.InsertFilmActor(filmID, directorID, in.TitleCn, in.FilmDirector)
	if err != nil {
		l.Logger.Error("error", "InsertFilmActor", err)
		return nil, errors.ErrorCMSFailed
	}
	if in.FilmActor1 != "" {
		filmActor1 := entity.Actor{
			NameCN:     in.FilmActor1,
			ActorPhoto: in.FilmActor1Img,
		}
		err, filmActor1ID := db.InsertActor(&filmActor1, 1)
		if err != nil {
			l.Logger.Error("error", "InsertActor", err)
			return nil, errors.ErrorCMSFailed
		}
		if filmActor1ID == 0 {
			l.Logger.Error("error", "filmActor1ID", filmActor1ID)
			return nil, errors.ErrorCMSFailed
		}
		err = db.InsertFilmActor(filmID, filmActor1ID, in.TitleCn, in.FilmActor1)
		if err != nil {
			l.Logger.Error("error", "InsertFilmActor", err)
			return nil, errors.ErrorCMSFailed
		}
	}
	if in.FilmActor2 != "" {
		filmActor2 := entity.Actor{
			NameCN:     in.FilmActor2,
			ActorPhoto: in.FilmActor2Img,
		}
		err, filmActor2ID := db.InsertActor(&filmActor2, 1)
		if err != nil {
			l.Logger.Error("error", "InsertActor", err)
			return nil, errors.ErrorCMSFailed
		}
		if filmActor2ID == 0 {
			l.Logger.Error("error", "filmActor2ID", filmActor2ID)
			return nil, errors.ErrorCMSFailed
		}
		err = db.InsertFilmActor(filmID, filmActor2ID, in.TitleCn, in.FilmActor2)
		if err != nil {
			l.Logger.Error("error", "InsertFilmActor", err)
			return nil, errors.ErrorCMSFailed
		}
	}
	return nil, nil
}
