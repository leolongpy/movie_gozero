// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package server

import (
	"context"

	"movie_gozero/rpc/comment/internal/logic"
	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"
)

type CommentServiceExtServer struct {
	svcCtx *svc.ServiceContext
}

func NewCommentServiceExtServer(svcCtx *svc.ServiceContext) *CommentServiceExtServer {
	return &CommentServiceExtServer{
		svcCtx: svcCtx,
	}
}

//  精彩影评
func (s *CommentServiceExtServer) HotComment(ctx context.Context, in *pb.HotCommentReq) (*pb.HotCommentRsp, error) {
	l := logic.NewHotCommentLogic(ctx, s.svcCtx)
	return l.HotComment(in)
}

//  进行评论
func (s *CommentServiceExtServer) MakeComment(ctx context.Context, in *pb.MakeCommentReq) (*pb.MakeCommentRsp, error) {
	l := logic.NewMakeCommentLogic(ctx, s.svcCtx)
	return l.MakeComment(in)
}

//  评论up
func (s *CommentServiceExtServer) UpNumComment(ctx context.Context, in *pb.UpNumCommentReq) (*pb.UpNumCommentRsp, error) {
	l := logic.NewUpNumCommentLogic(ctx, s.svcCtx)
	return l.UpNumComment(in)
}

//  我的评论
func (s *CommentServiceExtServer) MyComments(ctx context.Context, in *pb.MyCommentsReq) (*pb.MyCommentsRsp, error) {
	l := logic.NewMyCommentsLogic(ctx, s.svcCtx)
	return l.MyComments(in)
}

//  删除评论
func (s *CommentServiceExtServer) DeleteComment(ctx context.Context, in *pb.DeleteCommentReq) (*pb.DeleteCommentRsp, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}
