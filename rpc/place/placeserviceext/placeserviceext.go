// Code generated by goctl. DO NOT EDIT!
// Source: place.proto

package placeserviceext

import (
	"context"

	"movie_gozero/rpc/place/pb"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Place                = pb.Place
	HotCitiesByCinemaReq = pb.HotCitiesByCinemaReq
	HotCitiesByCinemaRep = pb.HotCitiesByCinemaRep

	PlaceServiceExt interface {
		//  获取所有电影院的地点
		HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq) (*HotCitiesByCinemaRep, error)
	}

	defaultPlaceServiceExt struct {
		cli zrpc.Client
	}
)

func NewPlaceServiceExt(cli zrpc.Client) PlaceServiceExt {
	return &defaultPlaceServiceExt{
		cli: cli,
	}
}

//  获取所有电影院的地点
func (m *defaultPlaceServiceExt) HotCitiesByCinema(ctx context.Context, in *HotCitiesByCinemaReq) (*HotCitiesByCinemaRep, error) {
	client := pb.NewPlaceServiceExtClient(m.cli.Conn())
	return client.HotCitiesByCinema(ctx, in)
}
