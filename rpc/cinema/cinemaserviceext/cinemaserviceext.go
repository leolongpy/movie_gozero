// Code generated by goctl. DO NOT EDIT!
// Source: cinema.proto

package cinemaserviceext

import (
	"context"

	"movie_gozero/rpc/cinema/pb"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	LocationCinemaReq        = pb.LocationCinemaReq
	HotMovie                 = pb.HotMovie
	GetCinemaMessageRsp      = pb.GetCinemaMessageRsp
	GetCinemaMessageByCidReq = pb.GetCinemaMessageByCidReq
	FilmMessage              = pb.FilmMessage
	GetCinemaMessageByCidRsp = pb.GetCinemaMessageByCidRsp
	GetMovieHallByMHIdReq    = pb.GetMovieHallByMHIdReq
	LocationCinemaRsp        = pb.LocationCinemaRsp
	Cinema                   = pb.Cinema
	GetCinemaMessageReq      = pb.GetCinemaMessageReq
	GetMovieHallByMHIdRsp    = pb.GetMovieHallByMHIdRsp
	Movie                    = pb.Movie

	CinemaServiceExt interface {
		//  地点影城
		LocationCinema(ctx context.Context, in *LocationCinemaReq) (*LocationCinemaRsp, error)
		//  根据位置查看有销售对应电影的影院信息
		GetCinemaMessageByCid(ctx context.Context, in *GetCinemaMessageByCidReq) (*GetCinemaMessageByCidRsp, error)
		//  根据mh_id获取影厅信息
		GetMovieHallByMHId(ctx context.Context, in *GetMovieHallByMHIdReq) (*GetMovieHallByMHIdRsp, error)
	}

	defaultCinemaServiceExt struct {
		cli zrpc.Client
	}
)

func NewCinemaServiceExt(cli zrpc.Client) CinemaServiceExt {
	return &defaultCinemaServiceExt{
		cli: cli,
	}
}

//  地点影城
func (m *defaultCinemaServiceExt) LocationCinema(ctx context.Context, in *LocationCinemaReq) (*LocationCinemaRsp, error) {
	client := pb.NewCinemaServiceExtClient(m.cli.Conn())
	return client.LocationCinema(ctx, in)
}

//  根据位置查看有销售对应电影的影院信息
func (m *defaultCinemaServiceExt) GetCinemaMessageByCid(ctx context.Context, in *GetCinemaMessageByCidReq) (*GetCinemaMessageByCidRsp, error) {
	client := pb.NewCinemaServiceExtClient(m.cli.Conn())
	return client.GetCinemaMessageByCid(ctx, in)
}

//  根据mh_id获取影厅信息
func (m *defaultCinemaServiceExt) GetMovieHallByMHId(ctx context.Context, in *GetMovieHallByMHIdReq) (*GetMovieHallByMHIdRsp, error) {
	client := pb.NewCinemaServiceExtClient(m.cli.Conn())
	return client.GetMovieHallByMHId(ctx, in)
}
