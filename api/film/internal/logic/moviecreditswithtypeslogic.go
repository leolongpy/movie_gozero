package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MovieCreditsWithTypesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovieCreditsWithTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) MovieCreditsWithTypesLogic {
	return MovieCreditsWithTypesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovieCreditsWithTypesLogic) MovieCreditsWithTypes(req types.MovieCreditsWithTypesReq) (*types.MovieCreditsWithTypesRsp, error) {
	resp, err := l.svcCtx.Film.MovieCreditsWithTypes(l.ctx, &filmserviceext.MovieCreditsWithTypesReq{
		MovieId: req.MovieId,
	})

	if err != nil {
		return &types.MovieCreditsWithTypesRsp{}, err
	}

	items := []*types.Type{}
	for _, v := range resp.Types {
		pers := []*types.Persons{}
		for _, v2 := range v.Persons {
			per := &types.Persons{
				Image:     v2.Image,
				Name:      v2.Name,
				NameEn:    v2.NameEn,
				RoleCover: v2.RoleCover,
				Personate: v2.Personate,
			}
			pers = append(pers, per)
		}
		item := &types.Type{
			TypeName:   v.TypeName,
			TypeNameEc: v.TypeNameEc,
			Persons:    pers,
		}
		items = append(items, item)
	}

	return &types.MovieCreditsWithTypesRsp{
		Types: items,
	}, nil
}
