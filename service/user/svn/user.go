package svn

import (
	"book/service/user/config"
	"book/service/user/dao"
	"book/service/user/database"
	"book/service/user/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	source := c.MySQLConf.DataSource
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(database.Connect(source)),
	}
}
