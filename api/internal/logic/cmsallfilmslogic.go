package logic

import (
	"context"
	"movie_gozero/rpc/cms/cmsservice"

	"movie_gozero/api/internal/svc"
	"movie_gozero/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CmsallFilmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCmsallFilmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CmsallFilmsLogic {
	return CmsallFilmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CmsallFilmsLogic) CmsallFilms(req types.AllFilmsReq) (*types.AllFilmsRsp, error) {
	resp, err := l.svcCtx.Cms.AllFilms(l.ctx, &cmsservice.AllFilmsReq{
		Page:    req.Page,
		AdminID: req.AdminID,
	})
	if err != nil {
		return &types.AllFilmsRsp{}, err
	}
	items := []*types.Film{}
	for _, v := range resp.Films {
		item := &types.Film{
			MovieID:              v.MovieID,
			Img:                  v.Img,
			Length:               v.Length,
			IsSelectSeat:         v.IsSelectSeat,
			FilmPrice:            v.FilmPrice,
			FilmScreenwriter:     v.FilmScreenwriter,
			FilmDirector:         v.FilmDirector,
			CommentNum:           v.CommentNum,
			TitleCn:              v.TitleCn,
			TitleEn:              v.TitleEn,
			IsSupportInlineWatch: v.IsSupportInlineWatch,
			CreateAt:             v.CreateAt,
			Type:                 v.Type,
			FilmDrama:            v.FilmDrama,
			CommonSpecial:        v.CommonSpecial,
			UserAccessTimes:      v.UserAccessTimes,
			FilmBoxoffice:        float64(v.FilmBoxoffice),
			WantedCount:          v.WantedCount,
			UserCommentTimes:     v.UserCommentTimes,
			CompanyIssued:        v.CompanyIssued,
			Country:              v.Country,
			RatingFinal:          v.RatingFinal,
			Is3D:                 v.Is3D,
			IsDMAX:               v.IsDMAX,
			IsFilter:             v.IsFilter,
			IsHot:                v.IsHot,
			IsIMAX:               v.IsIMAX,
			IsIMAX3D:             v.IsIMAX3D,
			IsNew:                v.IsNew,
			IsTicking:            v.IsTicking,
			RDay:                 v.RDay,
			RMonth:               v.RMonth,
			RYear:                v.RYear,
			ActorNames:           v.ActorNames,
			RYMD:                 v.RYMD,
			TicketStatus:         v.TicketStatus,
		}
		items = append(items, item)
	}
	return &types.AllFilmsRsp{
		Total: resp.Total,
		Films: items,
	}, nil
}
