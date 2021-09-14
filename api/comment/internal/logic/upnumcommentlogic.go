package logic

import (
	"context"
	"movie_gozero/rpc/comment/commentserviceext"

	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpNumCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpNumCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpNumCommentLogic {
	return UpNumCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpNumCommentLogic) UpNumComment(req types.UpNumCommentReq) (*types.UpNumCommentRsp, error) {
	resp, err := l.svcCtx.Comment.UpNumComment(l.ctx, &commentserviceext.UpNumCommentReq{
		CommentID: req.CommentID,
	})
	if err != nil {
		return &types.UpNumCommentRsp{}, err
	}

	return &types.UpNumCommentRsp{
		UpNum: resp.UpNum,
	}, nil
}
