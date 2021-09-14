package logic

import (
	"context"
	"movie_gozero/rpc/comment/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpNumCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpNumCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpNumCommentLogic {
	return &UpNumCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpNumComment 评论up
func (l *UpNumCommentLogic) UpNumComment(req *pb.UpNumCommentReq) (*pb.UpNumCommentRsp, error) {
	rsp := &pb.UpNumCommentRsp{}
	err := db.UpdateHotComment(req.CommentID)
	if err != nil {
		return rsp, errors.ErrorCommentFailed
	}
	upNum, err := db.SelectUpNum(req.CommentID)
	if err != nil {
		return rsp, errors.ErrorCommentFailed
	}
	rsp.UpNum = upNum
	return rsp, nil
}
