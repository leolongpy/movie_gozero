package logic

import (
	"context"
	"fmt"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AllFilmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllFilmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllFilmsLogic {
	return &AllFilmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllFilmsLogic) AllFilms(in *pb.AllFilmsReq) (*pb.AllFilmsRsp, error) {
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
	total, err := db.SelectFilmsTotal()
	if err != nil {
		l.Logger.Error("error", "SelectFilmsTotal", err)
		return nil, errors.ErrorCMSFailed

	}
	rsp := pb.AllFilmsRsp{}
	rsp.Total = total
	films, err := db.SelectAllFilms(page, 20)
	if err != nil {
		l.Logger.Error("error", "SelectAllFilms", err)
		return nil, errors.ErrorCMSFailed
	}
	filmsPB := []*pb.Film{}
	for _, film := range films {
		filmPB := film.ToProtoMovies()
		switch film.IsTicking {
		case 0:
			filmPB.TicketStatus = "已经上映"
			break
		case 1:
			filmPB.TicketStatus = "正在上映"
			break
		case 2:
			filmPB.TicketStatus = "即将上映"
			break
		default:
			filmPB.TicketStatus = "未知"
		}
		filmPB.RYMD = fmt.Sprintf("%d-%d-%d", filmPB.RYear, film.RMonth, film.RDay)
		filmsPB = append(filmsPB, filmPB)
	}
	rsp.Films = filmsPB

	return &rsp, nil
}
