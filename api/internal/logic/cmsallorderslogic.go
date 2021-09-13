package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsallOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsallOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsallOrdersLogic {
	return CmsallOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsallOrdersLogic) CmsallOrders(req types.AllOrdersReq) (*types.AllOrdersRsp, error) {
	resp, err := l.svcCtx.Cms.AllOrders(l.ctx, &cmsservice.AllOrdersReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllOrdersRsp{}, err
	}
	items := []*types.OrderAll{}
	for _, v := range resp.Orders {
		item := &types.OrderAll{
			OrderID:     v.OrderID,
			OrderNum:    v.OrderNum,
			OrderStatus: v.OrderStatus,
			OrderPrice:  v.OrderPrice,
			CreateAt:    v.CreateAt,
			PayAt:       v.PayAt,
			MhID:        v.MhID,
			OrderX:      v.OrderX,
			OrderY:      v.OrderY,
			UserID:      v.UserID,
			MovieID:     v.MovieID,
			OrderScore:  v.OrderScore,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			OrderStat:   v.OrderStat,
		}
		items = append(items, item)
	}
	return &types.AllOrdersRsp{
		Orders: items,
		Total:  resp.Total,
	}, nil
}
