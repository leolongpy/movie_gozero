package logic

import (
	"context"

	"movie_gozero/rpc/user/internal/svc"
	"movie_gozero/rpc/user/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ResetAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetAccountLogic {
	return &ResetAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ResetAccount 密码重置
func (l *ResetAccountLogic) ResetAccount(in *pb.ResetAccountReq) (*pb.ResetAccountRsp, error) {
	// todo: add your logic here and delete this line

	return &pb.ResetAccountRsp{}, nil
}
