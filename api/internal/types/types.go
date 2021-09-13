// Code generated by goctl. DO NOT EDIT.
package types

type RegistAccountReq struct {
	Email    string `form:"email"`
	UserName string `form:"userName"`
	Password string `form:"password"`
}

type RegistAccountRsp struct {
}

type LoginAccountReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginAccountRsp struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type ResetAccountReq struct {
}

type ResetAccountRsp struct {
}

type WantScoreReq struct {
	UserID  int64 `form:"userId"`
	MovieId int64 `form:"movieId"`
	Score   int64 `form:"score"`
}

type WantScoreRsp struct {
}

type UpdateUserProfileReq struct {
	UserImage string `form:"userImage"`
	UserName  string `form:"userName"`
	UserEmail string `form:"userEmail"`
	UserPhone string `form:"userPhone"`
	UserID    int64  `form:"userId"`
}

type UpdateUserProfileRsp struct {
}

type UserLoginReq struct {
	User     string `form:"user"`
	Password string `form:"password"`
}

type UserLoginRsp struct {
	CinemaID   int64  `json:"cinemaId"`
	AdminID    int64  `json:"adminId"`
	CinemaName string `json:"cinemaName"`
	AdminNum   int64  `json:"adminNum"`
}

type UpdateMessageReq struct {
	Table string `form:"table"` // 表名
	Json  string `form:"json"`  // {k:v}
	Num   string `form:"num"`   //  权限值
}

type UpdateMessageRsp struct {
}

type AllFilmsReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllFilmsRsp struct {
	Total int64   `json:"total"`
	Films []*Film `json:"films"`
}

type Film struct {
	MovieID              int64    `json:"movieId"`
	Img                  string   `json:"img"`
	Length               int64    `json:"length"`
	IsSelectSeat         int64    `json:"isSelectSeat"`
	FilmPrice            float32  `json:"filmPrice"`
	FilmScreenwriter     string   `json:"filmScreenwriter"`
	FilmDirector         string   `json:"filmDirector"`
	CommentNum           int64    `json:"commentNum"`
	TitleCn              string   `json:"titleCn"`
	TitleEn              string   `json:"titleEn"`
	IsSupportInlineWatch int64    `json:"isSupportInlineWatch"`
	CreateAt             string   `json:"createAt"`
	Type                 string   `json:"type"`
	FilmDrama            string   `json:"filmDrama"`
	CommonSpecial        string   `json:"commonSpecial"`
	UserAccessTimes      int64    `json:"userAccessTimes"`
	FilmBoxoffice        float64  `json:"filmBoxoffice"`
	WantedCount          int64    `json:"wantedCount"`
	UserCommentTimes     int64    `json:"userCommentTimes"`
	CompanyIssued        string   `json:"companyIssued"`
	Country              string   `json:"country"`
	RatingFinal          float32  `json:"ratingFinal"`
	Is3D                 int64    `json:"is3D"`
	IsDMAX               int64    `json:"isDmax"`
	IsFilter             int64    `json:"isFilter"`
	IsHot                int64    `json:"isHot"`
	IsIMAX               int64    `json:"isImax"`
	IsIMAX3D             int64    `json:"isImax3D"`
	IsNew                int64    `json:"isNew"`
	IsTicking            int64    `json:"isTicking"`
	RDay                 int64    `json:"rDay"`
	RMonth               int64    `json:"rMonth"`
	RYear                int64    `json:"rYear"`
	ActorNames           []string `json:"actorNames"`
	RYMD                 string   `json:"rYmd"`         // 时间格式化
	TicketStatus         string   `json:"ticketStatus"` // 上映状态
}

type AllUsersReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllUsersRsp struct {
	Users []*User `json:"users"`
	Total int64   `json:"total"`
}

type User struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type AllAdminUsersReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllAdminUsersRsp struct {
	AdminUsers []*AdminUser `json:"adminUsers"`
	Total      int64        `json:"total"`
}

type AdminUser struct {
	AuID               int64  `json:"auId"`
	AdminName          string `json:"adminName"`
	AdminPassword      string `json:"adminPassword"`
	AdminCinemaID      int64  `json:"adminCinemaId"`
	AdminLastLoginTime string `json:"adminLastLoginTime"`
	AdminNum           int64  `json:"adminNum"`
}

type AllCommentsReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllCommentsRsp struct {
	Comments []*Comment `json:"comments"`
	Total    int64      `json:"total"`
}

type Comment struct {
	CommentID int64  `json:"commentId"`
	FilmID    int64  `json:"filmId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	HeadImg   string `json:"headImg"`
	NickName  string `json:"nickName"`
	CreateAt  string `json:"createAt"`
	UpNum     int64  `json:"upNum"`
}

type AllOrdersReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllOrdersRsp struct {
	Orders []*OrderAll `json:"orders"`
	Total  int64       `json:"total"`
}

type OrderAll struct {
	OrderID     int64   `json:"orderId"`
	OrderNum    string  `json:"orderNum"`
	OrderStatus int64   `json:"orderStatus"`
	OrderPrice  float32 `json:"orderPrice"`
	CreateAt    string  `json:"createAt"`
	PayAt       string  `json:"payAt"`
	MhID        int64   `json:"mhId"`
	OrderX      int64   `json:"orderX"`
	OrderY      int64   `json:"orderY"`
	UserID      int64   `json:"userId"`
	MovieID     int64   `json:"movieId"`
	OrderScore  int64   `json:"orderScore"`
	StartTime   string  `json:"startTime"`
	EndTime     string  `json:"endTime"`
	OrderStat   string  `json:"orderStat"`
}

type AllAddressReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type AllAddressRsp struct {
	Places []*PlaceAll `json:"places"`
	Total  int64       `json:"total"`
}

type PlaceAll struct {
	Id          int64  `json:"id"`
	Count       int64  `json:"count"`
	Name        string `json:"name"`
	PinyinFull  string `json:"pinyinFull"`
	PinyinShort string `json:"pinyinShort"`
}

type AddFilmReq struct {
	AdminID              int64   `form:"adminId"`
	Img                  string  `form:"img"`
	Length               int64   `form:"length"`
	IsSelectSeat         int64   `form:"isSelectSeat"`
	FilmPrice            float32 `form:"filmPrice"`
	FilmScreenwriter     string  `form:"filmScreenwriter"`
	FilmDirector         string  `form:"filmDirector"`
	CommentNum           int64   `form:"commentNum"`
	TitleCn              string  `form:"titleCn"`
	TitleEn              string  `form:"titleEn"`
	IsSupportInlineWatch int64   `form:"isSupportInlineWatch"`
	CreateAt             string  `form:"createAt"`
	Type                 string  `form:"type"`
	FilmDrama            string  `form:"filmDrama"`
	CommonSpecial        string  `form:"commonSpecial"`
	UserAccessTimes      int64   `form:"userAccessTimes"`
	FilmBoxoffice        float32 `form:"filmBoxoffice"`
	WantedCount          int64   `form:"wantedCount"`
	UserCommentTimes     int64   `form:"userCommentTimes"`
	CompanyIssued        string  `form:"companyIssued"`
	Country              string  `form:"country"`
	RatingFinal          float32 `form:"ratingFinal"`
	Is3D                 int64   `form:"is3D"`
	IsDMAX               int64   `form:"isDmax"`
	IsFilter             int64   `form:"isFilter"`
	IsHot                int64   `form:"isHot"`
	IsIMAX               int64   `form:"isImax"`
	IsIMAX3D             int64   `form:"isImax3D"`
	IsNew                int64   `form:"isNew"`
	IsTicking            int64   `form:"isTicking"`
	RDay                 int64   `form:"rDay"`
	RMonth               int64   `form:"rMonth"`
	RYear                int64   `form:"rYear"`
	FilmDirectorImg      string  `form:"filmDirectorImg"`
	FilmActor1           string  `form:"filmActor1"`
	FilmActor1Img        string  `form:"filmActor1Img"`
	FilmActor2           string  `form:"filmActor2"`
	FilmActor2Img        string  `form:"filmActor2Img"`
}

type AddFilmRsp struct {
}

type UpdateFilmReq struct {
	MovieID       int64   `form:"movieId"`
	Img           string  `form:"img"`
	Length        int64   `form:"length"`
	FilmPrice     float32 `form:"filmPrice"`
	FilmDirector  string  `form:"filmDirector"`
	TitleCn       string  `form:"titleCn"`
	TitleEn       string  `form:"titleEn"`
	Type          string  `form:"type"`
	FilmDrama     string  `form:"filmDrama"`
	CommonSpecial string  `form:"commonSpecial"`
	CompanyIssued string  `form:"companyIssued"`
	Country       string  `form:"country"`
	RDay          int64   `form:"rDay"`
	RMonth        int64   `form:"rMonth"`
	RYear         int64   `form:"rYear"`
	RYMD          string  `form:"rYmd"`
	AdminID       int64   `form:"adminId"`
	IsTicking     int64   `form:"isTicking"`
}

type UpdateFilmRsp struct {
}

type DeleteFilmReq struct {
	MovieID int64 `form:"movieId"`
	AdminID int64 `form:"adminId"`
}

type DeleteFilmRsp struct {
}

type AddAdminUserReq struct {
	AdminID       int64  `form:"adminId"`
	AdminName     string `form:"adminName"`
	AdminPassword string `form:"adminPassword"`
	AdminCinemaID int64  `form:"adminCinemaId"`
	AdminNum      int64  `form:"adminNum"`
}

type AddAdminUserRsp struct {
}

type AddAddressReq struct {
	AdminID     int64  `form:"adminId"`
	Count       int64  `form:"count"`
	Name        string `form:"name"`
	PinyinFull  string `form:"pinyinFull"`
	PinyinShort string `form:"pinyinShort"`
}

type AddAddressRsp struct {
}

type UpdateAddressReq struct {
	Id          int64  `form:"id"`
	Count       int64  `form:"count"`
	Name        string `form:"name"`
	PinyinFull  string `form:"pinyinFull"`
	PinyinShort string `form:"pinyinShort"`
	AdminID     int64  `form:"adminId"`
}

type UpdateAddressRsp struct {
}

type DeleteAddressReq struct {
	Id      int64 `form:"id"`
	AdminID int64 `form:"adminId"`
}

type DeleteAddressRsp struct {
}

type DeleteAdminUserReq struct {
	AuID    int64 `form:"auId"`
	AdminID int64 `form:"adminId"`
}

type DeleteAdminUserRsp struct {
}

type AllMovieHallReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type MovieHall struct {
	MhId      int64  `json:"mhId"`
	MhName    string `json:"mhName"`
	MhAddress string `json:"mhAddress"`
	CinemaId  int64  `json:"cinemaId"`
}

type AllMovieHallRsp struct {
	MovieHalls []*MovieHall `json:"movieHalls"`
	Total      int64        `json:"total"`
}

type AddMovieHallReq struct {
	AdminID   int64  `form:"adminId"`
	MhName    string `form:"mhName"`
	MhAddress string `form:"mhAddress"`
	CinemaId  int64  `form:"cinemaId"`
}

type AddMovieHallRsp struct {
}

type UpdateMovieHallReq struct {
	AdminID   int64  `form:"adminId"`
	MhName    string `form:"mhName"`
	MhAddress string `form:"mhAddress"`
	CinemaId  int64  `form:"cinemaId"`
	MhId      int64  `form:"mhId"`
}

type UpdateMovieHallRsp struct {
}

type DeleteMovieHallReq struct {
	AdminID int64 `form:"adminId"`
	MhId    int64 `form:"mhId"`
}

type DeleteMovieHallRsp struct {
}

type AllCinemaFilmsReq struct {
	Page    int64 `form:"page"`
	AdminID int64 `form:"adminId"`
}

type CinemaFilm struct {
	CinemaID         int64   `json:"cinemaId"`
	FilmID           int64   `json:"filmId"`
	HallID           int64   `json:"hallId"`
	FilmName         string  `json:"filmName"`
	CinemaName       string  `json:"cinemaName"`
	ReleaseTimeYear  int64   `json:"releaseTimeYear"`
	ReleaseTimeMonth int64   `json:"releaseTimeMonth"`
	ReleaseTimeDay   int64   `json:"releaseTimeDay"`
	ReleaseTime      string  `json:"releaseTime"`
	ReleaseType      string  `json:"releaseType"`
	ReleaseAdd       string  `json:"releaseAdd"`
	CfID             int64   `json:"cfId"`
	Length           int64   `json:"length"`
	ReleaseDiscount  float64 `json:"releaseDiscount"`
	HallName         string  `json:"hallName"`
}

type AllCinemaFilmsRsp struct {
	CinemaFilms []*CinemaFilm `json:"cinemaFilms"`
	Total       int64         `json:"total"`
}

type AddCinemaFilmReq struct {
	CinemaID         int64   `form:"cinemaId"`
	MovieID          int64   `form:"movieId"`
	HallID           int64   `form:"hallId"`
	TitleCn          string  `form:"titleCn"`
	CinemaName       string  `form:"cinemaName"`
	ReleaseTimeYear  int64   `form:"releaseTimeYear"`
	ReleaseTimeMonth int64   `form:"releaseTimeMonth"`
	ReleaseTimeDay   int64   `form:"releaseTimeDay"`
	ReleaseTime      string  `form:"releaseTime"`
	Type             string  `form:"type"`
	ReleaseAdd       string  `form:"releaseAdd"`
	AdminID          int64   `form:"adminId"`
	Length           int64   `form:"length"`
	ReleaseDiscount  float32 `form:"releaseDiscount"`
}

type AddCinemaFilmRsp struct {
}

type UpdateCinemaFilmReq struct {
	CinemaID         int64   `form:"cinemaId"`
	FilmID           int64   `form:"filmId"`
	HallID           int64   `form:"hallId"`
	FilmName         string  `form:"filmName"`
	CinemaName       string  `form:"cinemaName"`
	ReleaseTimeYear  int64   `form:"releaseTimeYear"`
	ReleaseTimeMonth int64   `form:"releaseTimeMonth"`
	ReleaseTimeDay   int64   `form:"releaseTimeDay"`
	ReleaseTime      string  `form:"releaseTime"`
	ReleaseType      string  `form:"releaseType"`
	ReleaseAdd       string  `form:"releaseAdd"`
	AdminID          int64   `form:"adminId"`
	Length           int64   `form:"length"`
	ReleaseDiscount  float32 `form:"releaseDiscount"`
	CfID             int64   `form:"cfId"`
}

type UpdateCinemaFilmRsp struct {
}

type DeleteCinemaFilmReq struct {
	AdminID int64 `form:"adminId"`
	CfId    int64 `form:"cfId"`
}

type DeleteCinemaFilmRsp struct {
}

type RegisterCinemaReq struct {
	AdminID        int64  `form:"adminId"`
	CinemaName     string `form:"cinemaName"`
	CinemaAddress  string `form:"cinemaAddress"`
	LocationID     int64  `form:"locationId"` // 影院城市对应的位置
	CinemaTypes    string `form:"cinemaTypes"`
	CinemaCard     int64  `form:"cinemaCard"`     // 影城卡
	CinemaMinPrice int64  `form:"cinemaMinPrice"` // 几元起
	CinemaSupport  string `form:"cinemaSupport"`  // 影院提供的支持，包括退签等,用|隔开
	CinemaDiscount int64  `form:"cinemaDiscount"` // 影城卡最低减价多少元
	CinemaPhone    int64  `form:"cinemaPhone"`    // 影院电话
}

type RegisterCinemaRsp struct {
	CinemaID int64 `json:"cinemaId"`
}

type AllCinemaHallReq struct {
	CinemaID int64 `form:"cinemaId"`
	AdminID  int64 `form:"adminId"`
}

type AllCinemaHallRsp struct {
	HallAddresses []*HallAddressList `json:"hallAddresses"`
}

type HallAddressList struct {
	MhID   int64  `json:"mhId"`
	MhName string `json:"mhName"`
}
