// Code generated by goctl. DO NOT EDIT.
package types

type HotCommentReq struct {
	MovieId int64 `form:"movieId"`
}

type HotCommentRsp struct {
	Data CommentData `json:"data"`
}

type CommentData struct {
	Mini CommentMini `json:"mini"`
	Plus CommentPlus `json:"plus"`
}

type CommentMini struct {
	List  []*CommentRecord `json:"list"`
	Total int64            `json:"total"`
}

type CommentPlus struct {
	List  []*CommentRecord `json:"list"`
	Total int64            `json:"total"`
}

type CommentRecord struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	HeadImg   string `json:"headImg"`
	Nickname  string `json:"nickname"`
	CreateAt  string `json:"createAt"`
	UpNum     int64  `json:"upNum"`
	CommentID int64  `json:"commentId"`
}

type MakeCommentReq struct {
	MovieId  int64  `form:"movieId"` // 影片id
	Title    string `form:"title"`   // 标题
	HeadImg  string `form:"headImg"`
	Nickname string `form:"nickname"` // 用户昵称
	UserId   int64  `form:"userId"`   // 用户id
	Content  string `form:"content"`
}

type MakeCommentRsp struct {
}

type UpNumCommentReq struct {
	CommentID int64 `form:"commentId"`
}

type UpNumCommentRsp struct {
	UpNum int64 `json:"upNum"`
}

type MyCommentsReq struct {
	UserId int64 `form:"userId"`
}

type MyCommentsRsp struct {
	MyComments []*MyComment `json:"myComments"`
}

type MyComment struct {
	FilmImage string `json:"filmImage"`
	FilmName  string `json:"filmName"`
	Score     string `json:"score"`
	CommentID int64  `json:"commentId"`
	Content   string `json:"content"`
	UpNum     int64  `json:"upNum"`
}

type DeleteCommentReq struct {
	CommentID int64 `form:"commentId"`
}

type DeleteCommentRsp struct {
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
