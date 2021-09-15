package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchRsp, error) {
	resp, err := l.svcCtx.Film.Search(l.ctx, &filmserviceext.SearchReq{
		Q: req.Q,
	})

	if err != nil {
		return &types.SearchRsp{}, err
	}

	items := []*types.SearchMovie{}
	for _, v := range resp.Subjects {
		pers := []*types.Genres{}
		for _, v2 := range v.Genres {
			per := &types.Genres{
				Type: v2.Type,
			}
			pers = append(pers, per)
		}
		item := &types.SearchMovie{
			Title: v.Title,
			Id:    v.Id,
			Image: types.Images{
				Small: v.Image.Small,
			},
			Genres: pers,
			Year:   v.Year,
			Rating: types.Rating{
				Average: float64(v.Rating.Average),
			},
		}
		items = append(items, item)
	}

	return &types.SearchRsp{
		Subjects: items,
		Total:    resp.Total,
	}, nil
}
