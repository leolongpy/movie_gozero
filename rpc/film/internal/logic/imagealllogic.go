package logic

import (
	"context"
	"movie_gozero/rpc/film/internal/db"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ImageAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImageAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageAllLogic {
	return &ImageAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取海报信息
func (l *ImageAllLogic) ImageAll(req *pb.ImageAllReq) (*pb.ImageAllRsp, error) {
	rsp := &pb.ImageAllRsp{}
	movieId := req.MovieId
	images, err := db.SelectFilimImages(movieId)
	imagesPB := []*pb.Image{}
	if err != nil {
		l.Logger.Error("error", err)
		return rsp, errors.ErrorFilmFailed
	}
	for _, image := range images {
		imagePB := pb.Image{
			Image: image.ImageUrl,
		}
		imagesPB = append(imagesPB, &imagePB)
	}
	rsp.Images = imagesPB
	return rsp, nil
}
