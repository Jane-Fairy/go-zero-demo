package repo

import (
	"book/service/user/model"
	"context"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
}
