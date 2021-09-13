package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsallCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsallCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsallCommentsLogic {
	return CmsallCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsallCommentsLogic) CmsallComments(req types.AllCommentsReq) (*types.AllCommentsRsp, error) {
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
