// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"movie_gozero/api/film/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/film/hotPlayMovies",
				Handler: hotPlayMoviesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/movieDetail",
				Handler: movieDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/movieCreditsWithTypes",
				Handler: movieCreditsWithTypesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/imageAll",
				Handler: imageAllHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/locationMovies",
				Handler: locationMoviesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/movieComingNew",
				Handler: movieComingNewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/search",
				Handler: searchHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/film/getFilmsByCidADay",
				Handler: getFilmsByCidADayHandler(serverCtx),
			},
		},
	)
}