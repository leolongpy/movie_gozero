package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddFilmLogic {
	return AddFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFilmLogic) AddFilm(req types.AddFilmReq) (*types.AddFilmRsp, error) {
	_, err := l.svcCtx.Cms.AddFilm(l.ctx, &cmsservice.AddFilmReq{
		AdminID:              req.AdminID,
		Img:                  req.Img,
		Length:               req.Length,
		IsSelectSeat:         req.IsSelectSeat,
		FilmPrice:            req.FilmPrice,
		FilmScreenwriter:     req.FilmScreenwriter,
		FilmDirector:         req.FilmDirector,
		CommentNum:           req.CommentNum,
		TitleCn:              req.TitleCn,
		TitleEn:              req.TitleEn,
		IsSupportInlineWatch: req.IsSupportInlineWatch,
		CreateAt:             req.CreateAt,
		Type:                 req.Type,
		FilmDrama:            req.FilmDrama,
		CommonSpecial:        req.CommonSpecial,
		UserAccessTimes:      req.UserAccessTimes,
		FilmBoxoffice:        req.FilmBoxoffice,
		WantedCount:          req.WantedCount,
		UserCommentTimes:     req.UserCommentTimes,
		CompanyIssued:        req.CompanyIssued,
		Country:              req.Country,
		RatingFinal:          req.RatingFinal,
		Is3D:                 req.Is3D,
		IsDMAX:               req.IsDMAX,
		IsFilter:             req.IsFilter,
		IsHot:                req.IsHot,
		IsIMAX:               req.IsIMAX,
		IsIMAX3D:             req.IsIMAX3D,
		IsNew:                req.IsNew,
		IsTicking:            req.IsTicking,
		RDay:                 req.RDay,
		RMonth:               req.RMonth,
		RYear:                req.RYear,
		FilmDirectorImg:      req.FilmDirectorImg,
		FilmActor1:           req.FilmActor1,
		FilmActor1Img:        req.FilmActor1Img,
		FilmActor2:           req.FilmActor2,
		FilmActor2Img:        req.FilmActor2Img,
	})
	if err != nil {
		return &types.AddFilmRsp{}, err
	}
	return &types.AddFilmRsp{}, nil
}
