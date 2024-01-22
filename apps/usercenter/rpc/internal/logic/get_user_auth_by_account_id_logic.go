package logic

import (
	"context"

	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAccountIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAccountIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAccountIdLogic {
	return &GetUserAuthByAccountIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAccountIdLogic) GetUserAuthByAccountId(in *pb.GetUserAuthByAccountIdReq) (*pb.GetUserAuthyAccountIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAuthyAccountIdResp{}, nil
}
