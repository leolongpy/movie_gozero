package logic

import (
	"context"

	"movie_gozero/rpc/film/internal/svc"
	"movie_gozero/rpc/film/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  图片搜索
func (l *SearchLogic) Search(in *pb.SearchReq) (*pb.SearchRsp, error) {
	rsp := &pb.SearchRsp{}
	rsp.Total = 2
	genres1 := pb.Genres{
		Type: "喜剧",
	}
	genres2 := pb.Genres{
		Type: "悲剧",
	}
	rating := pb.Rating{
		Average: 3.7,
	}
	genres := []*pb.Genres{}
	genres = append(genres, &genres1)
	genres = append(genres, &genres2)
	image1 := pb.Images{
		Small: "https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p54250.jpg",
	}
	movie1 := pb.SearchMovie{
		Image:  &image1,
		Genres: genres,
		Title:  "金钱掌控",
		Id:     "1",
		Year:   "2017",
		Rating: &rating,
	}
	movie2 := pb.SearchMovie{
		Image:  &image1,
		Genres: genres,
		Title:  "金钱掌控",
		Id:     "1",
		Year:   "2017",
		Rating: &rating,
	}
	searchMovies := []*pb.SearchMovie{}
	searchMovies = append(searchMovies, &movie1)
	searchMovies = append(searchMovies, &movie2)
	rsp.Subjects = searchMovies
	return rsp, nil
}
