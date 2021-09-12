package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/rpc/cms/internal/entity"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterCinemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCinemaLogic {
	return &RegisterCinemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterCinemaLogic) RegisterCinema(req *pb.RegisterCinemaReq) (*pb.RegisterCinemaRsp, error) {
	adminID := req.AdminID
	if adminID == 0 {
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
	if admin.CinemaID != 0 {
		return nil, errors.ErrorCMSAlreadyRegister
	}
	cinema := entity.Cinema{
		CinemaAdd:      req.CinemaAddress,
		CinemaCard:     req.CinemaCard,
		CinemaDiscount: req.CinemaDiscount,
		CinemaMinPrice: req.CinemaMinPrice,
		CinemaName:     req.CinemaName,
		CinemaPhone:    req.CinemaPhone,
		CinemaSupport:  req.CinemaSupport,
		CinemaTypes:    req.CinemaTypes,
		LocationId:     req.LocationID,
	}
	err = db.InsertCinema(&cinema)
	if err != nil {
		l.Logger.Error("error", "InsertCinema", err)
		return nil, errors.ErrorCMSFailed
	}
	cinemaTmp, err := db.SelectCinema(&cinema)
	if err != nil {
		l.Logger.Error("error", "SelectCinema", err)
		return nil, errors.ErrorCMSFailed
	}
	err = db.UpdateAdminUser(adminID, req.CinemaName, cinemaTmp.CinemaId)
	if err != nil {
		l.Logger.Error("error", "UpdateAdminUser", err)
		return nil, errors.ErrorCMSFailed
	}
	rsp := pb.RegisterCinemaRsp{}
	rsp.CinemaID = admin.CinemaID
	return &rsp, nil
}
