package logic

import (
	"context"
	"go.uber.org/zap"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllAddressLogic {
	return &AllAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllAddressLogic) AllAddress(req *pb.AllAddressReq) (*pb.AllAddressRsp, error) {
	rsp := pb.AllAddressRsp{}
	adminID := req.AdminID
	page := req.Page
	if adminID == 0 || page == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", zap.Any("SelectAdminByAUID", err))
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return nil, errors.ErrorCMSForbiddenParam
	}
	total, err := db.SelectPlaceTotal()
	if err != nil {
		l.Logger.Error("error", "SelectPlaceTotal", err)
		return nil, errors.ErrorCMSFailed
	}
	rsp.Total = total
	places, err := db.SelectAllPlace(page, 20)
	if err != nil {
		l.Logger.Error("error", "SelectAllPlace", err)
		return nil, errors.ErrorCMSFailed
	}
	placeAllsPB := []*pb.PlaceAll{}
	for _, place := range places {
		placePB := place.ToProtoPlaceAll()
		placeAllsPB = append(placeAllsPB, placePB)
	}
	rsp.Places = placeAllsPB
	return &rsp, nil
}
