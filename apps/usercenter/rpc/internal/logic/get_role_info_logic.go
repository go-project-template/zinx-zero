package logic

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleInfoLogic {
	return &GetRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleInfoLogic) GetRoleInfo(in *pb.GetRoleInfoReq) (*pb.GetRoleInfoResp, error) {
	selectBuilder := l.svcCtx.UserRoleModel.SelectBuilder()
	selectBuilder.Where("account_id", in.GetAccountId())
	selectBuilder.Limit(1)
	roles, err := l.svcCtx.UserRoleModel.FindAll(l.ctx, selectBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "GetUserInfo find user_account db err , id:%d , err:%v", in.AccountId, err)
	}
	if len(roles) == 0 {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.AccountId)
	}
	role := roles[0]
	return &pb.GetRoleInfoResp{RoleId: role.RoleId, Nickname: role.Nickname}, nil
}
