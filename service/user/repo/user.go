package repo

import (
	"context"
	"go-zero-demo-master/go-zero-demo-master/book/service/user/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
}
