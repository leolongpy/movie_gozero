package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllOrdersLogic {
	return &AllOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllOrdersLogic) AllOrders(in *pb.AllOrdersReq) (*pb.AllOrdersRsp, error) {
	rsp := pb.AllOrdersRsp{}
	adminID := in.AdminID
	page := in.Page
	if adminID == 0 || page == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", "SelectAdminByAUID", err)
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	// 超级管理员可以查看所有的信息
	if admin.AdminNum == 1 {
		total, err := db.SelectOrderTotal()
		if err != nil {
			l.Logger.Error("error", "SelectOrderTotal", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		orders, err := db.SelectAllOrder(page, 20)
		if err != nil {
			l.Logger.Error("error", "SelectAllOrder", err)
			return nil, errors.ErrorCMSFailed
		}
		ordersPB := []*pb.OrderAll{}
		for _, order := range orders {
			orderPB := order.ToProtoOrder()
			if order.OrderStatus == 0 {
				orderPB.OrderStat = "未支付"
			} else {
				orderPB.OrderStat = "已支付"
			}
			ordersPB = append(ordersPB, orderPB)
		}
		rsp.Orders = ordersPB
	}
	// 影院管理员可以查看所属影院信息
	if admin.AdminNum == 0 {
		total, err := db.SelectOrderTotalByFilmId(admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectOrderTotalByFilmId", err)
			return nil, errors.ErrorCMSFailed

		}
		rsp.Total = total
		ordersPB := []*pb.OrderAll{}
		movieHalls, err := db.SelectAllMovieHallsBycinemaID(admin.CinemaID)
		if err != nil {
			l.Logger.Error("error", "SelectAllMovieHallsBycinemaID", err)
			return nil, errors.ErrorCMSFailed
		}
		for _, movieHall := range movieHalls {
			orders, err := db.SelectOrderByFilmId(page, 20, movieHall.MhID)
			if err != nil {
				l.Logger.Error("error", "SelectOrderByFilmId", err)
				return nil, errors.ErrorCMSFailed
			}
			for _, order := range orders {
				orderPB := order.ToProtoOrder()
				ordersPB = append(ordersPB, orderPB)
			}
		}
		rsp.Orders = ordersPB
	}

	return &rsp, nil
}
