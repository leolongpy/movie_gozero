package logic

import (
	"context"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageLogic {
	return &UpdateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMessageLogic) UpdateMessage(in *pb.UpdateMessageReq) (*pb.UpdateMessageRsp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateMessageRsp{}, nil
}
