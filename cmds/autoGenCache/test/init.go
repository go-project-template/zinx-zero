package test

import (
	"autoGenCache/internal/config"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/autoGenCache-api.yaml", "the config file")

func GetConf() config.Config {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()
	return c
}
