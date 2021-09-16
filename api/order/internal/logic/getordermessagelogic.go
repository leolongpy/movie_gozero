package logic

import (
	"context"
	"movie_gozero/rpc/order/orderserviceext"

	"movie_gozero/api/order/internal/svc"
	"movie_gozero/api/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetOrderMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderMessageLogic {
	return GetOrderMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderMessageLogic) GetOrderMessage(req types.GetOrderMessageReq) (*types.GetOrderMessageRsp, error) {
	resp, err := l.svcCtx.Order.GetOrderMessage(l.ctx, &orderserviceext.GetOrderMessageReq{
		OrderNum: req.OrderNum,
		UserId:   req.UserId,
	})

	if err != nil {
		return &types.GetOrderMessageRsp{}, err
	}

	//items := []types.Movie{}
	//for _, v := range resp.Movies {
	//	item := types.Movie{
	//		Img:               v.Img,
	//		MovieId:           v.MovieId,
	//		TitleCn:           v.TitleCn,
	//		MoviesSupportType: v.MoviesSupportType,
	//		FilmDirector:      v.FilmDirector,
	//		Actors:            v.Actors,
	//		FilmDrama:         v.FilmDrama,
	//		RatingFinal:       float64(v.RatingFinal),
	//	}
	//	items = append(items, item)
	//}

	return &types.GetOrderMessageRsp{
		Ticketdetail: types.TicketDetail{
			FilmName:      resp.Ticketdetail.FilmName,
			FilmImg:       resp.Ticketdetail.FilmImg,
			StartTime:     resp.Ticketdetail.StartTime,
			EndTime:       resp.Ticketdetail.EndTime,
			CinemaName:    resp.Ticketdetail.CinemaName,
			MhName:        resp.Ticketdetail.MhName,
			Seat:          resp.Ticketdetail.Seat,
			OrderNum:      resp.Ticketdetail.OrderNum,
			CinemaAddress: resp.Ticketdetail.CinemaAddress,
			Price:         float64(resp.Ticketdetail.Price),
			CreateAt:      resp.Ticketdetail.CreateAt,
			Phone:         resp.Ticketdetail.Phone,
			CinemaPhone:   resp.Ticketdetail.CinemaPhone,
		},
	}, nil
}
