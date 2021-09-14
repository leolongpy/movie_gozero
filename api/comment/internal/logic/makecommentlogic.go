package logic

import (
	"context"
	"movie_gozero/rpc/comment/commentserviceext"

	"movie_gozero/api/comment/internal/svc"
	"movie_gozero/api/comment/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MakeCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) MakeCommentLogic {
	return MakeCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeCommentLogic) MakeComment(req types.MakeCommentReq) (*types.MakeCommentRsp, error) {
	_, err := l.svcCtx.Comment.MakeComment(l.ctx, &commentserviceext.MakeCommentReq{
		MovieId:  req.MovieId,
		Title:    req.Title,
		HeadImg:  req.HeadImg,
		Nickname: req.Nickname,
		UserId:   req.UserId,
		Content:  req.Content,
	})
	if err != nil {
		return &types.MakeCommentRsp{}, err
	}

	return &types.MakeCommentRsp{}, nil
}
