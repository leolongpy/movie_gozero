// Code generated by goctl. DO NOT EDIT!
// Source: place.proto

package server

import (
	"context"

	"movie_gozero/rpc/place/internal/logic"
	"movie_gozero/rpc/place/internal/svc"
	"movie_gozero/rpc/place/pb"
)

type PlaceServiceExtServer struct {
	svcCtx *svc.ServiceContext
}

func NewPlaceServiceExtServer(svcCtx *svc.ServiceContext) *PlaceServiceExtServer {
	return &PlaceServiceExtServer{
		svcCtx: svcCtx,
	}
}

//  获取所有电影院的地点
func (s *PlaceServiceExtServer) HotCitiesByCinema(ctx context.Context, in *pb.HotCitiesByCinemaReq) (*pb.HotCitiesByCinemaRep, error) {
	l := logic.NewHotCitiesByCinemaLogic(ctx, s.svcCtx)
	return l.HotCitiesByCinema(in)
}
