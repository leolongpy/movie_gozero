// Code generated by goctl. DO NOT EDIT.
package types

type Place struct {
	Count       int64  `json:"count"`
	Id          int64  `json:"id"`
	N           string `json:"n"`
	PinyinFull  string `json:"pinyinFull"`
	PinyinShort string `json:"pinyinShort"`
}

type HotCitiesByCinemaReq struct {
}

type HotCitiesByCinemaRep struct {
	P []*Place `json:"p"`
}