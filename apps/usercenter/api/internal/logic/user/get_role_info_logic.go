package user

import (
	"context"

	"zinx-zero/apps/acommon/ctxdata"
	"zinx-zero/apps/usercenter/api/internal/svc"
	"zinx-zero/apps/usercenter/api/internal/types"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(req *types.GetRoleInfoReq) (*types.GetRoleInfoResp, error) {
	accountId, err := ctxdata.GetAccountIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	getRoleInfoResp, err := l.svcCtx.UsercenterRpc.GetRoleInfo(l.ctx, &usercenter.GetRoleInfoReq{
		AccountId: accountId,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetRoleInfoResp
	_ = copier.Copy(&resp, getRoleInfoResp)

	return &resp, nil

}
