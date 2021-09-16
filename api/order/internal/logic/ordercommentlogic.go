package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type OrderCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCommentLogic {
	return OrderCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCommentLogic) OrderComment(req types.OrderCommentReq) (*types.OrderCommentRsp, error) {
	_, err := l.svcCtx.Order.OrderComment(l.ctx, &orderserviceext.OrderCommentReq{
		UserId:         req.UserId,
		Score:          req.Score,
		CommentContent: req.CommentContent,
		OrderNum:       req.OrderNum,
		MovieId:        req.MovieId,
	})

	if err != nil {
		return &types.OrderCommentRsp{}, err
	}

	return &types.OrderCommentRsp{}, nil
}
