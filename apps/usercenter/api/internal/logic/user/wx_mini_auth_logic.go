package user

import (
	"context"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/usercenter/api/internal/svc"
	"zinx-zero/apps/usercenter/api/internal/types"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"github.com/pkg/errors"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

// ErrWxMiniAuthFailError error
var ErrWxMiniAuthFailError = aerr.NewErrMsg("wechat mini auth fail")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {

	//1、Wechat-Mini
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起授权请求失败 err : %v , code : %s  , authResult : %+v", err, req.Code, authResult)
	}
	//2、Parsing WeChat-Mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "解析数据失败 req : %+v , err: %v , authResult:%+v ", req, err, authResult)
	}

	//3、bind user or login.
	var accountId int64
	rpcRsp, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
		AuthType: globalkey.Model_UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
	}
	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.AccountId == 0 {
		//bind user.

		//Wechat-Mini Decrypted data
		mobile := userData.PhoneNumber
		registerRsp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
			AuthKey:  authResult.OpenID,
			AuthType: globalkey.Model_UserAuthTypeSmallWX,
			Mobile:   mobile,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
		}

		return &types.WXMiniAuthResp{
			AccessToken:  registerRsp.AccessToken,
			AccessExpire: registerRsp.AccessExpire,
			RefreshAfter: registerRsp.RefreshAfter,
		}, nil

	} else {
		accountId = rpcRsp.UserAuth.AccountId
		tokenResp, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
			AccountId: accountId,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterRpc.GenerateToken err :%v, accountId : %d", err, accountId)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}, nil
	}
}