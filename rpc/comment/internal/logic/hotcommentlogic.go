package logic

import (
	"context"
	"movie_gozero/rpc/comment/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/comment/internal/svc"
	"movie_gozero/rpc/comment/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type HotCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotCommentLogic {
	return &HotCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HotComment 精彩影评
func (l *HotCommentLogic) HotComment(req *pb.HotCommentReq) (*pb.HotCommentRsp, error) {
	rsp := &pb.HotCommentRsp{}
	movieId := req.MovieId
	comments, err := db.SelectHotComment(movieId)
	if err != nil {
		l.Logger.Error("err", err)
		return rsp, errors.ErrorCommentFailed
	}
	records := []*pb.CommentRecord{}
	for _, comment := range comments {
		record := pb.CommentRecord{
			Title:     comment.Title,
			Content:   comment.Content,
			HeadImg:   comment.HeadImg,
			Nickname:  comment.NickName,
			CreateAt:  comment.CreateAt,
			UpNum:     comment.UpNum,
			CommentID: comment.CommentId,
		}
		records = append(records, &record)
	}

	plus := pb.CommentPlus{
		Total: int64(len(comments)),
		List:  records,
	}

	data := pb.CommentData{
		Plus: &plus,
	}
	rsp.Data = &data
	return rsp, nil
}
