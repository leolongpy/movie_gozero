package logic

import (
	"context"
	"movie_gozero/rpc/order/internal/db"
	"movie_gozero/rpc/order/internal/entity"
	"movie_gozero/utils/errors"
	"strconv"

	"movie_gozero/rpc/order/internal/svc"
	"movie_gozero/rpc/order/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCommentLogic {
	return &OrderCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderComment 进行评分
func (l *OrderCommentLogic) OrderComment(req *pb.OrderCommentReq) (*pb.OrderCommentRsp, error) {
	rsp := &pb.OrderCommentRsp{}
	userId := req.UserId
	score := req.Score
	orderNum := req.OrderNum
	content := req.CommentContent
	order, err := db.SelectOrderScore(orderNum, userId)
	if err != nil {
		l.Logger.Error("err", "SelectOrderScore", err)
		return rsp, errors.ErrorOrderFailed
	}
	if order == nil {
		l.Logger.Error("err", "order", err)
		return rsp, errors.ErrorOrderFailed
	}
	if order.OrderScore != -1 {
		return rsp, errors.ErrorOrderAlreadyScore
	}
	err = db.UpdateOrderScore(score, userId, orderNum)
	if err != nil {
		l.Logger.Error("err", "UpdateOrderScore", err)
		return rsp, errors.ErrorOrderFailed
	}
	err = db.UpdateFilmScore(order.MovieId, score)
	if err != nil {
		l.Logger.Error("err", "UpdateFilmScore", err)
		return rsp, errors.ErrorOrderFailed
	}
	user, err := db.SelectUserNameByUserId(userId)
	if err != nil {
		l.Logger.Error("err", "SelectUserNameByUserId", err)
		return rsp, errors.ErrorOrderFailed
	}
	comment := entity.Comment{
		FilmId:   order.MovieId,
		Title:    strconv.Itoa(int(score)),
		Content:  content,
		NickName: user.UserName,
		UserId:   userId,
	}
	err = db.InsertComment(&comment)
	if err != nil {
		l.Logger.Error("err", "InsertComment", err)
		return rsp, errors.ErrorOrderFailed
	}
	return rsp, nil
}
