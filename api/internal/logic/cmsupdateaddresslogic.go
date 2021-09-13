package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsupdateAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsupdateAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsupdateAddressLogic {
	return CmsupdateAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsupdateAddressLogic) CmsupdateAddress(req types.UpdateAddressReq) (*types.UpdateAddressRsp, error) {
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
