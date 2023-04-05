package dao

import (
	"context"
	"go-zero-demo-master/go-zero-demo-master/book/service/user/model"
)

type UserDao struct {
}

func NewUseDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	return nil
}
