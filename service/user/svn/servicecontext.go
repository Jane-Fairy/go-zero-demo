package svn

import (
	"book/service/user/config"
	"book/service/user/datebase"
	"book/service/user/repo"
)

type ServiceContext struct {
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := datebase.Connect(c)
}
