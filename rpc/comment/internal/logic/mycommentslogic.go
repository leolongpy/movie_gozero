package logic

import (
	"context"
	"movie_gozero/rpc/comment/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type MyCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMyCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyCommentsLogic {
	return &MyCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MyComments 我的评论
func (l *MyCommentsLogic) MyComments(req *pb.MyCommentsReq) (*pb.MyCommentsRsp, error) {
	rsp := &pb.MyCommentsRsp{}
	comments, err := db.SelectMyComment(req.UserId)
	if err != nil {
		l.Logger.Error("err", err)
		return rsp, errors.ErrorCommentFailed
	}
	commentsPB := []*pb.MyComment{}
	for _, comment := range comments {

		img, title, err := db.SelectFilmImageAndName(comment.FilmId)
		if err != nil {
			l.Logger.Error("err", err)
			return rsp, errors.ErrorCommentFailed
		}
		commentPB := pb.MyComment{
			Content:   comment.Content,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
			FilmImage: img,
			FilmName:  title,
			Score:     comment.Title,
		}
		commentsPB = append(commentsPB, &commentPB)
	}
	rsp.MyComments = commentsPB
	return rsp, nil
}
