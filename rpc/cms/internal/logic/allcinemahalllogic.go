package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllCinemaHallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllCinemaHallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllCinemaHallLogic {
	return &AllCinemaHallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllCinemaHallLogic) AllCinemaHall(req *pb.AllCinemaHallReq) (*pb.AllCinemaHallRsp, error) {
	rsp := pb.AllCinemaHallRsp{}
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
	if admin.AdminNum == 0 {
		if admin.CinemaID != req.CinemaID {
			return nil, errors.ErrorCMSForbiddenParam
		}
	}
	movieHalls, err := db.SelectAllMovieHallsBycinemaID(admin.CinemaID)
	if err != nil {
		l.Logger.Error("error", "SelectAllMovieHallsBycinemaID", err)
		return nil, errors.ErrorCMSFailed
	}
	halls := []*pb.HallAddressList{}
	for _, movieHall := range movieHalls {
		hall := pb.HallAddressList{
			MhName: movieHall.MhName,
			MhID:   movieHall.MhID,
		}
		halls = append(halls, &hall)
	}
	rsp.HallAddresses = halls
	return &rsp, nil
}
