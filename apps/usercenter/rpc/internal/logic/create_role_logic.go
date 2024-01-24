package logic

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/acommon/arand"
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	userAccount, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.AccountId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "CreateRole find user_account db err , id:%d , err:%v", in.AccountId, err)
	}
	if userAccount == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.AccountId)
	}
	if len(in.Nickname) == 0 {
		in.Nickname = arand.RandLetterN(1) + arand.RandDigitN(7)
	}
	roleIdStr, err := l.svcCtx.RedisClient.SpopCtx(l.ctx, globalkey.Cache_GenRoleId_UserIdPool)
	if err != nil {
		return nil, errors.Wrapf(aerr.NewErrCode(aerr.REDIS_ERROR), "CreateRole user gen roleId err:%v", err)
	}
	var roleId = cast.ToInt64(roleIdStr)
	if err := l.svcCtx.UserRoleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		userRoleIdPool, err := l.svcCtx.UserRoleidPoolModel.TxFindOneUnuse(ctx, session, roleId)
		if err != nil && err != model.ErrNotFound {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "CreateRole find userRoleidPoolModel db err , id:%d , err:%v", roleId, err)
		}
		if userRoleIdPool == nil {
			return errors.Wrapf(ErrUserNoExistsError, "id:%d", roleId)
		}
		userRole := new(model.UserRole)
		userRole.RoleId = userRoleIdPool.RoleId
		userRole.Sex = int64(in.Sex)
		userRole.Nickname = in.Nickname
		userRole.AccountId = in.AccountId
		_, err = l.svcCtx.UserRoleModel.Insert(ctx, session, userRole)
		if err != nil {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "CreateRole db user_role Insert err:%v,user:%+v", err, userRole)
		}
		userRoleIdPool.IsUse = 1
		_, err = l.svcCtx.UserRoleidPoolModel.Update(ctx, session, userRoleIdPool)
		if err != nil {
			return errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "GetUserInfo update UserRoleidPoolModel db err , id:%d , err:%v", roleId, err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &usercenter.CreateRoleResp{
		RoleId: roleId,
	}, nil
}
