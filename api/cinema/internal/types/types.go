// Code generated by goctl. DO NOT EDIT.
package types

type LocationCinemaReq struct {
	LocationId int64 `form:"locationId"` // 地点id
}

type LocationCinemaRsp struct {
	Cinemas []*Cinema `json:"cinemas"`
}

type Cinema struct {
	CinemaName     string `json:"cinemaName"`     // 影院名字
	CinemaAddress  string `json:"cinemaAddress"`  // 影院位置
	CinemaSupport  string `json:"cinemaSupport"`  // 影院支持的功能比如改签，用|隔开
	CinemaCard     int64  `json:"cinemaCard"`     // 是否支持影城卡
	CinemaMinPrice int64  `json:"cinemaMinPrice"` // 几元起
	CinemaDiscount int64  `json:"cinemaDiscount"` // 最低减价多少
	CinemaId       int64  `json:"cinemaId"`       // 影院id
}

type GetCinemaMessageReq struct {
	MovieId    int64  `json:"movieId"`    // 影片id
	LocationId int64  `json:"locationId"` // 地点id
	Day        string `json:"day"`        // 查询时间,今天明天后天
}

type GetCinemaMessageRsp struct {
	CinemaName     string `json:"cinemaName"`     // 影院名字
	CinemaAddress  string `json:"cinemaAddress"`  // 影院位置
	CinemaSupport  string `json:"cinemaSupport"`  // 影院支持的功能比如改签，用|隔开
	CinemaCard     int64  `json:"cinemaCard"`     // 是否支持影城卡
	CinemaMinPrice int64  `json:"cinemaMinPrice"` // 几元起
	CinemaDiscount int64  `json:"cinemaDiscount"` // 最低减价多少
}

type GetCinemaMessageByCidReq struct {
	CinemaId int64 `form:"cinemaId"`
}

type FilmMessage struct {
	FilmName    string   `json:"filmName"`
	RatingFinal float64  `json:"ratingFinal"`
	Length      int64    `json:"length"`
	Type        string   `json:"type"`      // 类型，如剧情
	ActorName   []string `json:"actorName"` // 主演名字
	MovieId     int64    `json:"movieId"`   // 影片id
	Img         string   `json:"img"`       // 影片logo
}

type GetCinemaMessageByCidRsp struct {
	Cinema      Cinema         `json:"cinema"`
	FilmMessage []*FilmMessage `json:"filmMessage"` // 影片信息
}

type GetMovieHallByMHIdReq struct {
	MhId int64 `form:"mhId"`
}

type GetMovieHallByMHIdRsp struct {
	MhAddress string `json:"mhAddress"`
}

type Movie struct {
	Img               string  `json:"img"`
	MovieId           int64   `json:"movieId"`
	TitleCn           string  `json:"titleCn"`
	MoviesSupportType string  `json:"moviesSupportType"`
	FilmDirector      string  `json:"filmDirector"`
	Actors            string  `json:"actors"`
	FilmDrama         string  `json:"filmDrama"`
	RatingFinal       float64 `json:"ratingFinal"`
}

type HotMovie struct {
	ActorName     []string `json:"actorName"`     // 演员名字
	CommonSpecial string   `json:"commonSpecial"` // 影片简介
	DirectorName  string   `json:"directorName"`  // 导演名字
	Img           string   `json:"img"`           // 影片logo
	Is3D          int64    `json:"is3D"`          // 是否3D
	IsDMAX        int64    `json:"isDmax"`        // 是否DMAX
	IsFilter      int64    `json:"isFilter"`      // 是否过滤
	IsHot         int64    `json:"isHot"`         // 是否热映
	IsIMAX        int64    `json:"isImax"`        // 是否IMAX
	IsIMAX3D      int64    `json:"isImax3D"`      // 是否IMAX3D
	IsNew         int64    `json:"isNew"`         // 是否新上映
	Length        int64    `json:"length"`        // 时长
	MovieId       int64    `json:"movieId"`       // 电影时长
	RDay          int64    `json:"rDay"`          // 上映日期
	RMonth        int64    `json:"rMonth"`        // 上映月份
	RYear         int64    `json:"rYear"`         // 上映年份
	RatingFinal   float64  `json:"ratingFinal"`   // 最终评分
	T             string   `json:"t"`             // 名字
	TitleCn       string   `json:"titleCn"`       // 中文名字
	TitleEn       string   `json:"titleEn"`       // 英文名字
	Type          string   `json:"type"`          // 影片类型
	WantedCount   int64    `json:"wantedCount"`   // 想看的人数
}
