package logic

import (
	"context"
	"movie_gozero/rpc/comment/commentserviceext"

	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteCommentLogic {
	return DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req types.DeleteCommentReq) (*types.DeleteCommentRsp, error) {
	_, err := l.svcCtx.Comment.DeleteComment(l.ctx, &commentserviceext.DeleteCommentReq{
		CommentID: req.CommentID,
	})
	if err != nil {
		return &types.DeleteCommentRsp{}, err
	}

	return &types.DeleteCommentRsp{}, nil
}
