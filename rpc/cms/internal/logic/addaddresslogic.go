package logic

import (
	"context"
	"movie_gozero/rpc/cms/internal/db"
	"movie_gozero/rpc/cms/internal/entity"
	"movie_gozero/utils/errors"

	"movie_gozero/rpc/cms/internal/svc"
	"movie_gozero/rpc/cms/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddAddressLogic) AddAddress(req *pb.AddAddressReq) (*pb.AddAddressRsp, error) {
	adminID := req.AdminID
	if adminID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	admin, err := db.SelectAdminByAUID(adminID)
	if err != nil {
		l.Logger.Error("error", "SelectAdminByAUID", err)
		return nil, errors.ErrorCMSFailed
	}
	if admin == nil || admin.AuID == 0 {
		return nil, errors.ErrorCMSFailedParam
	}
	if admin.AdminNum == 0 {
		return nil, errors.ErrorCMSForbiddenParam
	}
	place := entity.Place{
		Name:        req.Name,
		PinyinFull:  req.PinyinFull,
		PinyinShort: req.PinyinShort,
	}
	err = db.AddPlace(&place)
	if err != nil {
		l.Logger.Error("error", "AddPlace", err)
		return nil, errors.ErrorCMSFailed
	}
	return nil, nil
}
