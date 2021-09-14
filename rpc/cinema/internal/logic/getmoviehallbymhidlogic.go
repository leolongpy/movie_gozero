package logic

import (
	"context"
	"go.uber.org/zap"
	"movie_gozero/rpc/cinema/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cinema/internal/svc"
	"movie_gozero/rpc/cinema/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetMovieHallByMHIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMovieHallByMHIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMovieHallByMHIdLogic {
	return &GetMovieHallByMHIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMovieHallByMHId 根据mh_id获取影厅信息
func (l *GetMovieHallByMHIdLogic) GetMovieHallByMHId(req *pb.GetMovieHallByMHIdReq) (*pb.GetMovieHallByMHIdRsp, error) {
	rsp := &pb.GetMovieHallByMHIdRsp{}
	mhAddress, err := db.SelectMHAddress(req.MhId)
	if err != nil {
		l.Logger.Error("error", zap.Error(err))
		return rsp, errors.ErrorCinemaFailed
	}
	rsp.MhAddress = mhAddress
	return rsp, nil
}
