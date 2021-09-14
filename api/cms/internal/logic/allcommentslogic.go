package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) AllCommentsLogic {
	return AllCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllCommentsLogic) AllComments(req types.AllCommentsReq) (*types.AllCommentsRsp, error) {
	resp, err := l.svcCtx.Cms.AllComments(l.ctx, &cmsservice.AllCommentsReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllCommentsRsp{}, err
	}
	items := []*types.Comment{}
	for _, v := range resp.Comments {
		item := &types.Comment{
			CommentID: v.CommentID,
			FilmID:    v.FilmID,
			Title:     v.Title,
			Content:   v.Content,
			HeadImg:   v.HeadImg,
			NickName:  v.NickName,
			CreateAt:  v.CreateAt,
			UpNum:     v.UpNum,
		}
		items = append(items, item)
	}
	return &types.AllCommentsRsp{
		Comments: items,
		Total:    resp.Total,
	}, nil
}
