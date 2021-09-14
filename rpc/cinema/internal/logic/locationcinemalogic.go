package logic

import (
	"context"
	"movie_gozero/rpc/cinema/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cinema/internal/svc"
	"movie_gozero/rpc/cinema/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LocationCinemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLocationCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LocationCinemaLogic {
	return &LocationCinemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LocationCinema 地点影城
func (l *LocationCinemaLogic) LocationCinema(req *pb.LocationCinemaReq) (*pb.LocationCinemaRsp, error) {
	locationId := req.LocationId
	cinemas, err := db.SelectCinemasByLid(locationId)
	if err != nil {
		l.Logger.Error("error", err)
		return nil, errors.ErrorCinemaFailed
	}
	pbCinemas := []*pb.Cinema{}
	for _, cinema := range cinemas {

		pbCinema := pb.Cinema{
			CinemaId:       cinema.CinemaId,
			CinemaAddress:  cinema.CinemaAdd,
			CinemaName:     cinema.CinemaName,
			CinemaSupport:  cinema.CinemaSupport,
			CinemaCard:     cinema.CinemaCard,
			CinemaMinPrice: cinema.CinemaMinPrice,
			CinemaDiscount: cinema.CinemaDiscount,
		}
		pbCinemas = append(pbCinemas, &pbCinema)

	}
	rsp := &pb.LocationCinemaRsp{}
	rsp.Cinemas = pbCinemas
	return rsp, nil
}
