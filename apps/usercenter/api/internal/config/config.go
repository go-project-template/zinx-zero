package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 微信小程序配置
type WxMiniConf struct {
	AppId  string `json:"AppId"`  //微信appId
	Secret string `json:"Secret"` //微信secret
}

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	WxMiniConf        WxMiniConf
	UsercenterRpcConf zrpc.RpcClientConf
}
