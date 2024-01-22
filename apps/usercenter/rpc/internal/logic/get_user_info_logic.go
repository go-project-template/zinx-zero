package logic

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.AccountId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.AccountId, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.AccountId)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user)

	return &usercenter.GetUserInfoResp{
		User: &respUser,
	}, nil

}
