package logic

import (
	"context"
	"movie_gozero/rpc/comment/commentserviceext"

	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) HotCommentLogic {
	return HotCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotCommentLogic) HotComment(req types.HotCommentReq) (*types.HotCommentRsp, error) {
	resp, err := l.svcCtx.Comment.HotComment(l.ctx, &commentserviceext.HotCommentReq{
		MovieId: req.MovieId,
	})
	if err != nil {
		return &types.HotCommentRsp{}, err
	}
	items := []*types.CommentRecord{}
	items2 := []*types.CommentRecord{}
	for _, v := range resp.Data.Mini.List {
		item := &types.CommentRecord{
			Title:     v.Title,
			Content:   v.Content,
			HeadImg:   v.HeadImg,
			Nickname:  v.Nickname,
			CreateAt:  v.CreateAt,
			UpNum:     v.UpNum,
			CommentID: v.CommentID,
		}
		items = append(items, item)
	}

	for _, v := range resp.Data.Plus.List {
		item := &types.CommentRecord{
			Title:     v.Title,
			Content:   v.Content,
			HeadImg:   v.HeadImg,
			Nickname:  v.Nickname,
			CreateAt:  v.CreateAt,
			UpNum:     v.UpNum,
			CommentID: v.CommentID,
		}
		items2 = append(items2, item)
	}

	return &types.HotCommentRsp{
		Data: types.CommentData{
			Mini: types.CommentMini{
				List:  items,
				Total: resp.Data.Mini.Total,
			},
			Plus: types.CommentPlus{
				List:  items2,
				Total: resp.Data.Plus.Total,
			},
		},
	}, nil
}
