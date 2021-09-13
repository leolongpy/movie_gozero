package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsupdateMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsupdateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsupdateMessageLogic {
	return CmsupdateMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsupdateMessageLogic) CmsupdateMessage(req types.UpdateMessageReq) (*types.UpdateMessageRsp, error) {
	_, err := l.svcCtx.Cms.UpdateMessage(l.ctx, &cmsservice.UpdateMessageReq{
		Table: req.Table,
		Json:  req.Json,
		Num:   req.Num,
	})
	if err != nil {
		return &types.UpdateMessageRsp{}, err
	}

	return &types.UpdateMessageRsp{}, nil
}
