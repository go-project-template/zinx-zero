package logic

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/acommon/astring"
	"zinx-zero/apps/acommon/autils"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/cmd/rpc/internal/svc"
	"zinx-zero/apps/usercenter/cmd/rpc/pb"
	"zinx-zero/apps/usercenter/cmd/rpc/usercenter"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrUserAlreadyRegisterError = aerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	_user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "mobile:%s,err:%v", in.Mobile, err)
	}
	if _user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", in.Mobile, err)
	}

	var userId int64
	if len(in.Nickname) == 0 {
		in.Nickname = astring.RandLetterN(1) + astring.RandDigitN(7)
	}
	if len(in.Password) > 0 {
		in.Password = autils.Md5HexByString(in.Password)
	}
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := new(model.User)
		user.Mobile = in.Mobile
		user.Nickname = in.Nickname
		user.Password = in.Password
		_, err := l.svcCtx.UserModel.Insert(ctx, session, user)
		if err != nil {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
		}
		userId, err = l.svcCtx.IDWorker.NextID()
		if err != nil {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "Register db user IDWorker.NextID err:%v,user:%+v", err, user)
		}

		userAuth := new(model.UserAuth)
		userAuth.UserId = userId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &usercenter.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
