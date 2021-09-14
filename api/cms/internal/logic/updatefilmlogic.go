package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateFilmLogic {
	return UpdateFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFilmLogic) UpdateFilm(req types.UpdateFilmReq) (*types.UpdateFilmRsp, error) {
	_, err := l.svcCtx.Cms.UpdateFilm(l.ctx, &cmsservice.UpdateFilmReq{
		MovieID:       req.MovieID,
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
		RDay:          req.RDay,
		RMonth:        req.RMonth,
		RYear:         req.RYear,
		RYMD:          req.RYMD,
		AdminID:       req.AdminID,
		IsTicking:     req.IsTicking,
	})
	if err != nil {
		return &types.UpdateFilmRsp{}, err
	}

	return &types.UpdateFilmRsp{}, nil
}
