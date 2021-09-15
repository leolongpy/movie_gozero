package logic

import (
	"context"
	"movie_gozero/rpc/film/filmserviceext"

	"movie_gozero/api/film/internal/svc"
	"movie_gozero/api/film/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ImageAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImageAllLogic {
	return ImageAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageAllLogic) ImageAll(req types.ImageAllReq) (*types.ImageAllRsp, error) {
	resp, err := l.svcCtx.Film.ImageAll(l.ctx, &filmserviceext.ImageAllReq{
		MovieId: req.MovieId,
	})

	if err != nil {
		return &types.ImageAllRsp{}, err
	}

	items := []*types.Image{}
	for _, v := range resp.Images {
		item := &types.Image{
			Image: v.Image,
		}
		items = append(items, item)
	}

	return &types.ImageAllRsp{
		Images: items,
	}, nil
}
