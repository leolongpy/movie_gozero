package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllCommentsLogic {
	return &AllCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllCommentsLogic) AllComments(in *pb.AllCommentsReq) (*pb.AllCommentsRsp, error) {
	adminID := in.AdminID
	page := in.Page
	if adminID == 0 || page == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", "SelectAdminByAUID", err)
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	rsp := pb.AllCommentsRsp{}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectCommentTotal()
		if err != nil {
			l.Logger.Error("error", "SelectCommentTotal", err)
			return nil, errors.ErrorCMSFailed

		}

		rsp.Total = total
		comments, err := db.SelectAllComment(page, 20)
		if err != nil {
			l.Logger.Error("error", "SelectAllComment", err)
			return nil, errors.ErrorCMSFailed
		}
		commentsPB := []*pb.Comment{}
		for _, comment := range comments {
			commentPB := comment.ToProtoComment()
			commentsPB = append(commentsPB, commentPB)
		}
		rsp.Comments = commentsPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		// 根据所属影院id获取影片id
		filmIDs, err := db.SelectFilmsID(admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectFilmsID", err)
			return nil, errors.ErrorCMSFailed
		}
		var total int64 = 0
		commentsPB := []*pb.Comment{}
		for _, filmID := range filmIDs {
			total_tmp, err := db.SelectCommentsTotalByCID(filmID)
			if err != nil {
				l.Logger.Error("error", "SelectCommentsTotalByCID", err)
				return nil, errors.ErrorCMSFailed
			}
			comments, err := db.SelectCommentsByCID(page, 20, filmID)
			if err != nil {
				l.Logger.Error("error", "SelectCommentsTotalByCID", err)
				return nil, errors.ErrorCMSFailed
			}
			for _, comment := range comments {
				commentPB := comment.ToProtoComment()
				commentsPB = append(commentsPB, commentPB)
			}
			total = total + total_tmp
			rsp.Comments = commentsPB
		}
		rsp.Total = total
	}

	return &rsp, nil
}
