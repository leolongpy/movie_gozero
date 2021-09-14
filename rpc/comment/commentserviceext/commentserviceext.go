// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package commentserviceext

import (
	"context"

	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	CommentData      = pb.CommentData
	CommentPlus      = pb.CommentPlus
	MakeCommentRsp   = pb.MakeCommentRsp
	MyCommentsReq    = pb.MyCommentsReq
	CommentRecord    = pb.CommentRecord
	MakeCommentReq   = pb.MakeCommentReq
	UpNumCommentReq  = pb.UpNumCommentReq
	DeleteCommentReq = pb.DeleteCommentReq
	HotCommentReq    = pb.HotCommentReq
	UpNumCommentRsp  = pb.UpNumCommentRsp
	DeleteCommentRsp = pb.DeleteCommentRsp
	HotCommentRsp    = pb.HotCommentRsp
	CommentMini      = pb.CommentMini
	MyCommentsRsp    = pb.MyCommentsRsp
	MyComment        = pb.MyComment
	Movie            = pb.Movie
	HotMovie         = pb.HotMovie

	CommentServiceExt interface {
		//  精彩影评
		HotComment(ctx context.Context, in *HotCommentReq) (*HotCommentRsp, error)
		//  进行评论
		MakeComment(ctx context.Context, in *MakeCommentReq) (*MakeCommentRsp, error)
		//  评论up
		UpNumComment(ctx context.Context, in *UpNumCommentReq) (*UpNumCommentRsp, error)
		//  我的评论
		MyComments(ctx context.Context, in *MyCommentsReq) (*MyCommentsRsp, error)
		//  删除评论
		DeleteComment(ctx context.Context, in *DeleteCommentReq) (*DeleteCommentRsp, error)
	}

	defaultCommentServiceExt struct {
		cli zrpc.Client
	}
)

func NewCommentServiceExt(cli zrpc.Client) CommentServiceExt {
	return &defaultCommentServiceExt{
		cli: cli,
	}
}

//  精彩影评
func (m *defaultCommentServiceExt) HotComment(ctx context.Context, in *HotCommentReq) (*HotCommentRsp, error) {
	client := pb.NewCommentServiceExtClient(m.cli.Conn())
	return client.HotComment(ctx, in)
}

//  进行评论
func (m *defaultCommentServiceExt) MakeComment(ctx context.Context, in *MakeCommentReq) (*MakeCommentRsp, error) {
	client := pb.NewCommentServiceExtClient(m.cli.Conn())
	return client.MakeComment(ctx, in)
}

//  评论up
func (m *defaultCommentServiceExt) UpNumComment(ctx context.Context, in *UpNumCommentReq) (*UpNumCommentRsp, error) {
	client := pb.NewCommentServiceExtClient(m.cli.Conn())
	return client.UpNumComment(ctx, in)
}

//  我的评论
func (m *defaultCommentServiceExt) MyComments(ctx context.Context, in *MyCommentsReq) (*MyCommentsRsp, error) {
	client := pb.NewCommentServiceExtClient(m.cli.Conn())
	return client.MyComments(ctx, in)
}

//  删除评论
func (m *defaultCommentServiceExt) DeleteComment(ctx context.Context, in *DeleteCommentReq) (*DeleteCommentRsp, error) {
	client := pb.NewCommentServiceExtClient(m.cli.Conn())
	return client.DeleteComment(ctx, in)
}
