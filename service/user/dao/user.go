package dao

import (
	"book/service/user/database"
	"book/service/user/model"
	"context"
	"fmt"
)

type UserDao struct {
	*database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	//todo dao
	sql := fmt.Sprintf("insert into %s (id ,name) values(? , ?)", user.TableName())
	execCtx, err := d.Conn.ExecCtx(ctx, sql, user.Id, user.Name)
	if err != nil {
		return err
	}
	id, err := execCtx.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}
