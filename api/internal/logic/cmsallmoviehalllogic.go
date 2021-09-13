package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsallMovieHallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsallMovieHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsallMovieHallLogic {
	return CmsallMovieHallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsallMovieHallLogic) CmsallMovieHall(req types.AllMovieHallReq) (*types.AllMovieHallRsp, error) {
	resp, err := l.svcCtx.Cms.AllMovieHall(l.ctx, &cmsservice.AllMovieHallReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllMovieHallRsp{}, err
	}
	items := []*types.MovieHall{}
	for _, v := range resp.MovieHalls {
		item := &types.MovieHall{
			MhId:      v.MhId,
			MhName:    v.MhName,
			MhAddress: v.MhAddress,
			CinemaId:  v.CinemaId,
		}
		items = append(items, item)
	}
	return &types.AllMovieHallRsp{
		MovieHalls: items,
		Total:      resp.Total,
	}, nil
}
