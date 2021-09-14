package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/cms/internal/svc"
	"movie_gozero/api/cms/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateAddressLogic {
	return UpdateAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAddressLogic) UpdateAddress(req types.UpdateAddressReq) (*types.UpdateAddressRsp, error) {
	_, err := l.svcCtx.Cms.UpdateAddress(l.ctx, &cmsservice.UpdateAddressReq{
		Id:          req.Id,
		Count:       req.Count,
		Name:        req.Name,
		PinyinFull:  req.PinyinFull,
		PinyinShort: req.PinyinShort,
		AdminID:     req.AdminID,
	})
	if err != nil {
		return &types.UpdateAddressRsp{}, err
	}

	return &types.UpdateAddressRsp{}, nil
}
