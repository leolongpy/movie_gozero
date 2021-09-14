package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateMessageLogic {
	return UpdateMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMessageLogic) UpdateMessage(req types.UpdateMessageReq) (*types.UpdateMessageRsp, error) {
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
