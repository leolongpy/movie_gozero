package logic

import (
	"context"
	"movie_gozero/rpc/comment/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteComment 删除评论
func (l *DeleteCommentLogic) DeleteComment(req *pb.DeleteCommentReq) (*pb.DeleteCommentRsp, error) {
	rsp := &pb.DeleteCommentRsp{}
	err := db.DeleteComment(req.CommentID)
	if err != nil {
		l.Logger.Error("err", err)
		return rsp, errors.ErrorCommentFailed
	}
	return rsp, nil
}
