package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySQLConf MySQLConf
}

type MySQLConf struct {
	DataSource string
}
