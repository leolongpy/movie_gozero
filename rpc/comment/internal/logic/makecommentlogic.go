package logic

import (
	"context"
	"movie_gozero/rpc/comment/internal/db"
	"movie_gozero/rpc/comment/internal/entity"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type MakeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeCommentLogic {
	return &MakeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MakeComment 进行评论
func (l *MakeCommentLogic) MakeComment(req *pb.MakeCommentReq) (*pb.MakeCommentRsp, error) {
	rsp := &pb.MakeCommentRsp{}
	comment := entity.Comment{
		Title:    req.Title,
		Content:  req.Content,
		HeadImg:  req.HeadImg,
		FilmId:   req.MovieId,
		NickName: req.Nickname,
	}
	err := db.InsertHotComment(&comment)
	if err != nil {
		return rsp, errors.ErrorCommentFailed
	}
	return rsp, nil
}
