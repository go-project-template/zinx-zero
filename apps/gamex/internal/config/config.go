package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ZinxConf *ZinxConf
	JwtAuth  struct {
		AccessSecret string
		AccessExpire int64
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}

type ZinxConf struct {
	Name             string // The name of the current server.(当前服务器名称)
	Host             string // The IP address of the current server. (当前服务器主机IP)
	WsPort           int    // The port number on which the server listens for WebSocket connections.(当前服务器主机websocket监听端口)
	MaxPacketSize    uint32 // The maximum size of the packets that can be sent or received.(读写数据包的最大值)
	MaxConn          int    // The maximum number of connections that the server can handle.(当前服务器主机允许的最大链接个数)
	MaxWorkerTaskLen uint32 // The maximum number of tasks that a worker pool can handle.(业务工作Worker对应负责的任务队列最大任务存储数量)
	WorkerMode       string // The way to assign workers to connections.(为链接分配worker的方式)
	MaxMsgChanLen    uint32 // The maximum length of the send buffer message queue.(SendBuffMsg发送消息的缓冲最大长度)
	IOReadBuffSize   uint32 // The maximum size of the read buffer for each IO operation.(每次IO最大的读取长度)
	//The server mode, which can be "tcp" or "websocket". If it is empty, both modes are enabled.
	//"tcp":tcp监听, "websocket":websocket 监听 为空时同时开启
	Mode string
	// A boolean value that indicates whether the new or old version of the router is used. The default value is false.
	// 路由模式 false为旧版本路由，true为启用新版本的路由 默认使用旧版本
	RouterSlicesMode  bool
	LogIsolationLevel int
	HeartbeatMax      int
}
