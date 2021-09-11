package logic

import (
	"context"
	"movie_gozero/rpc/user/userservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WantScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWantScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) WantScoreLogic {
	return WantScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WantScoreLogic) WantScore(req types.WantScoreReq) (*types.WantScoreRsp, error) {
	var _, err = l.svcCtx.User.WantScore(l.ctx, &userservice.WantScoreReq{
		UserID:  req.UserID,
		MovieId: req.MovieId,
		Score:   req.Score,
	})
	if err != nil {
		return &types.WantScoreRsp{}, err
	}
	return &types.WantScoreRsp{}, nil
}
