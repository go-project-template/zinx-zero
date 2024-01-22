package main

import (
	"autoGenCache/internal/config"
	"autoGenCache/internal/svc"
	"autoGenCache/templ"
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/autoGenCache-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()

	ctx := svc.NewServiceContext(c)
	// gen cache template code
	findAll, err := ctx.RedisCacheKeyModel.FindAll(context.Background(), ctx.RedisCacheKeyModel.SelectBuilder(), "")
	logx.Must(err)
	logx.Must(os.RemoveAll("cachex"))
	for _, val := range findAll {
		templ.GenAll(val)
	}
	fmt.Println("")
	fmt.Println("Gen cache code success")
}
