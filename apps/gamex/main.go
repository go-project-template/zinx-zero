package main

import (
	"flag"
	"zinx-zero/apps/acommon"
	"zinx-zero/apps/gamex/internal/config"
	"zinx-zero/apps/gamex/internal/gamex"
	"zinx-zero/apps/gamex/internal/svc"

	"github.com/aceld/zinx/zlog"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "conf/app.yaml", "the config file")

func main() {
	flag.Parse()

	// ## Prepare service config
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	// Set zinx use go-zero logger
	zlog.SetLogger(new(acommon.MyLogger))

	// Special config for development environment
	if ctx.Config.Mode != service.ProMode {
		logx.SetLevel(logx.DebugLevel)
	}

	// ## Start service logic
	s := gamex.NewGameX(ctx)
	s.Start()
}
