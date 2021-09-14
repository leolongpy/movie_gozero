package logic

import (
	"context"
	"movie_gozero/rpc/comment/commentserviceext"

	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MyCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) MyCommentsLogic {
	return MyCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyCommentsLogic) MyComments(req types.MyCommentsReq) (*types.MyCommentsRsp, error) {
	resp, err := l.svcCtx.Comment.MyComments(l.ctx, &commentserviceext.MyCommentsReq{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.MyCommentsRsp{}, err
	}
	items := []*types.MyComment{}
	for _, v := range resp.MyComments {
		item := &types.MyComment{
			FilmImage: v.FilmImage,
			FilmName:  v.FilmName,
			Score:     v.Score,
			CommentID: v.CommentID,
			Content:   v.Content,
			UpNum:     v.UpNum,
		}
		items = append(items, item)
	}

	return &types.MyCommentsRsp{
		MyComments: nil,
	}, nil
}
