package logic

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/usercenter/cmd/rpc/internal/svc"
	"zinx-zero/apps/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = aerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = aerr.NewErrMsg("账号或密码不正确")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResp{}, nil
}
