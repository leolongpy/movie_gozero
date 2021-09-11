package logic

import (
	"context"
	"movie_gozero/rpc/user/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type WantScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWantScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WantScoreLogic {
	return &WantScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// WantScore 评分
func (l *WantScoreLogic) WantScore(in *pb.WantScoreReq) (*pb.WantScoreRsp, error) {
	// todo: add your logic here and delete this line
	orderNum, err := db.SelectOrderByUidMid(in.MovieId, in.UserID)
	if err != nil {
		l.Logger.Error("error", err)
		return nil, errors.ErrorUserFailed
	}
	if orderNum == 0 {
		l.Logger.Error("error", err)
		return nil, errors.ErrorScoreForbid
	}
	err = db.UpdateOrderScore(in.MovieId, in.Score)
	if err != nil {
		l.Logger.Error("error", err)
		return nil, errors.ErrorUserFailed
	}
	return nil, nil
}
