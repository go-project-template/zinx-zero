package user

import (
	"context"

	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/usercenter/cmd/api/internal/svc"
	"zinx-zero/apps/usercenter/cmd/api/internal/types"
	"zinx-zero/apps/usercenter/cmd/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: globalkey.Model_UserAuthTypeSystem,
		AuthKey:  req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	_ = copier.Copy(resp, loginResp)

	return resp, nil
}
