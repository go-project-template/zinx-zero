package user

import (
	"context"

	"zinx-zero/apps/acommon/ctxdata"
	"zinx-zero/apps/usercenter/api/internal/svc"
	"zinx-zero/apps/usercenter/api/internal/types"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.CreateRoleResp, err error) {
	accountId, err := ctxdata.GetAccountIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	createRoleResp, err := l.svcCtx.UsercenterRpc.CreateRole(l.ctx, &usercenter.CreateRoleReq{
		AccountId: accountId,
		Nickname:  req.Nickname,
		Sex:       req.Sex,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateRoleResp{
		RoleId: createRoleResp.RoleId,
	}, nil
}
