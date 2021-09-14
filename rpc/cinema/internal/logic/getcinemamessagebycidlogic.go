package logic

import (
	"context"
	"movie_gozero/rpc/cinema/internal/db"
	"movie_gozero/utils/common"
	"movie_gozero/utils/errors"
	"time"

	"movie_gozero/rpc/cinema/internal/svc"
	"movie_gozero/rpc/cinema/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetCinemaMessageByCidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCinemaMessageByCidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaMessageByCidLogic {
	return &GetCinemaMessageByCidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCinemaMessageByCid 根据位置查看有销售对应电影的影院信息
func (l *GetCinemaMessageByCidLogic) GetCinemaMessageByCid(req *pb.GetCinemaMessageByCidReq) (*pb.GetCinemaMessageByCidRsp, error) {
	rsp := &pb.GetCinemaMessageByCidRsp{}
	cinemaId := req.CinemaId
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	cinema, err := db.SelectCinemaByCid(cinemaId)
	if err != nil {
		l.Logger.Error("error", err)
		return nil, errors.ErrorCinemaFailed
	}

	movieIds, err := db.SelectMidByCid(cinemaId, int64(year), common.SwitchMonth(month.String()), int64(day))
	if err != nil {
		l.Logger.Error("error", err)
		return nil, errors.ErrorCinemaFailed
	}

	filmsPB := []*pb.FilmMessage{}
	for _, movieId := range movieIds {

		film, err := db.SelectFilmMesage(movieId)
		if err != nil {
			l.Logger.Error("error", err)
			return nil, errors.ErrorCinemaFailed
		}
		if film != nil {
			actors, err := db.SelectActorNameByMid(film.MovieId)
			if err != nil {
				l.Logger.Error("error", err)
				return nil, errors.ErrorCinemaFailed
			}
			if actors != nil {
				for _, actor := range actors {

					film.ActorName = append(film.ActorName, actor.ActorName)
				}
			}
			filmPB := pb.FilmMessage{

				FilmName:    film.TitleCn,
				RatingFinal: film.RatingFinal,
				Length:      film.Length,
				Type:        film.Type,
				MovieId:     film.MovieId,
				ActorName:   film.ActorName,
				Img:         film.Img,
			}
			filmsPB = append(filmsPB, &filmPB)
			rsp.FilmMessage = filmsPB
		}
	}
	if cinema != nil {
		cinemaPB := pb.Cinema{
			CinemaName:     cinema.CinemaName,
			CinemaAddress:  cinema.CinemaAdd,
			CinemaSupport:  cinema.CinemaSupport,
			CinemaCard:     cinema.CinemaCard,
			CinemaMinPrice: cinema.CinemaMinPrice,
			CinemaDiscount: cinema.CinemaDiscount,
			CinemaId:       cinema.CinemaId,
		}
		rsp.Cinema = &cinemaPB
	}
	return rsp, nil
}
