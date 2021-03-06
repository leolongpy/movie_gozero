package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCinemaFilmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllCinemaFilmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllCinemaFilmsLogic {
	return AllCinemaFilmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllCinemaFilmsLogic) AllCinemaFilms(req types.AllCinemaFilmsReq) (*types.AllCinemaFilmsRsp, error) {
	resp, err := l.svcCtx.Cms.AllCinemaFilms(l.ctx, &cmsservice.AllCinemaFilmsReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllCinemaFilmsRsp{}, err
	}
	film := []*types.CinemaFilm{}
	for _, v := range resp.CinemaFilms {
		films := &types.CinemaFilm{
			CinemaID:         v.CinemaID,
			FilmID:           v.FilmID,
			HallID:           v.HallID,
			FilmName:         v.FilmName,
			CinemaName:       v.CinemaName,
			ReleaseTimeYear:  v.ReleaseTimeYear,
			ReleaseTimeMonth: v.ReleaseTimeMonth,
			ReleaseTimeDay:   v.ReleaseTimeDay,
			ReleaseTime:      v.ReleaseTime,
			ReleaseType:      v.ReleaseType,
			ReleaseAdd:       v.ReleaseAdd,
			CfID:             v.CfID,
			Length:           v.Length,
			ReleaseDiscount:  float64(v.ReleaseDiscount),
			HallName:         v.HallName,
		}
		film = append(film, films)
	}
	return &types.AllCinemaFilmsRsp{
		CinemaFilms: nil,
		Total:       resp.Total,
	}, nil
}
